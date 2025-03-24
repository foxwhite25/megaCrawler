package main

import (
	"fmt"
	"github.com/gocolly/colly/v2/proxy"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/sourcegraph/conc/pool"

	"megaCrawler/crawlers"
	"megaCrawler/crawlers/tester"

	"github.com/olekukonko/tablewriter"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func TestWebMap(_ *testing.T) {
	crawlers.Threads = 64
	for _, v := range crawlers.WebMap {
		println(v.BaseURL.Hostname())
	}
}

func waitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()
	select {
	case <-c:
		return false // completed normally
	case <-time.After(timeout):
		return true // timed out
	}
}

func TestTester(t *testing.T) {
	buf, err := os.Create("table.txt")
	bufMutex := sync.Mutex{}
	if err != nil {
		t.Error(err)
		return
	}
	crawlers.Threads = 16

	targetEnv := os.Getenv("TARGET")
	targets := strings.Split(targetEnv, ",")
	if len(targets) == 1 && strings.TrimSpace(targets[0]) == "" {
		targets = make([]string, len(crawlers.WebMap))
		i := 0
		for k := range crawlers.WebMap {
			targets[i] = k
			i++
		}
	}
	_, err = fmt.Fprintf(buf, "Testing targets %s.\n\n", targets)
	if err != nil {
		t.Error(err)
		return
	}

	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./log/debug.jsonl",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
	})
	var fileCore zapcore.Core
	ProductionEncoder := zap.NewProductionEncoderConfig()

	fileCore = zapcore.NewCore(
		zapcore.NewJSONEncoder(ProductionEncoder),
		w,
		zap.DebugLevel,
	)

	logger := zap.New(fileCore)

	crawlers.Sugar = logger.Sugar()

	if p := os.Getenv("HTTP_PROXY"); p != "" {
		fmt.Printf("Using Proxy %s\n", p)
		rp, err := proxy.RoundRobinProxySwitcher(p)

		if err == nil {
			crawlers.Proxy = rp
		} else {
			crawlers.Sugar.Panicf("Cannot parse proxy in HTTP_PROXY: %s", p)
		}
	}

	completed := 0
	max := len(targets)

	handle := func(target string) {
		c, ok := crawlers.WebMap[target]
		if !ok {
			bufMutex.Lock()
			_, _ = fmt.Fprintf(buf, "No such target %s.\n\n", target)
			bufMutex.Unlock()
			return
		}

		c.Test = &tester.Tester{
			WG: &sync.WaitGroup{},
			News: tester.Status{
				Name: "News",
			},
			Index: tester.Status{
				Name: "Index",
			},
			Expert: tester.Status{
				Name: "Expert",
			},
			Report: tester.Status{
				Name: "Report",
			},
			Sugar: crawlers.Sugar,
		}
		c.Test.WG.Add(1)
		go crawlers.StartEngine(c, true)
		if waitTimeout(c.Test.WG, 2*time.Minute) {
			if !c.Test.Done {
				c.Test.Complete("timeout", c.ID)
			}
		}

		table := tablewriter.NewWriter(buf)
		table.SetHeader([]string{"Field", "Total", "Passed", "Coverage"})

		c.Test.News.FillTable(table)
		c.Test.Index.FillTable(table)
		c.Test.Expert.FillTable(table)
		c.Test.Report.FillTable(table)

		bufMutex.Lock()
		_, _ = buf.WriteString(target + "; " + c.Test.Reason + ":\n")
		table.Render()
		_, err := buf.WriteString("\n")
		if err != nil {
			_, _ = fmt.Fprintf(buf, "Writing Table Errored: %s", err)
		}
		bufMutex.Unlock()
		completed += 1
		t.Log("Completed "+target, "(", completed, "/", max, ")")
	}

	runner := pool.New().WithMaxGoroutines(16)
	n := 0
	for _, target := range targets {
		n += 1
		tar := target
		t.Log(tar, "(", n, "/", max, ")")
		runner.Go(func() {
			handle(tar)
		})
	}
	runner.Wait()
}

func TestName(t *testing.T) {
	parseAny := crawlers.TimeCleanup("From 10 Sep 2019 9:00\n")
	t.Log(parseAny)
}

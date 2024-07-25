package main

import (
	"fmt"
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

func TestTester(t *testing.T) {
	buf, err := os.Create("table.txt")
	bufMutex := sync.Mutex{}
	if err != nil {
		t.Error(err)
		return
	}
	crawlers.Threads = 64

	targetEnv := os.Getenv("TARGET")
	targets := strings.Split(targetEnv, ",")
	if len(targets) == 1 && targets[0] == "" {
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

		go func() {
			time.Sleep(2 * time.Minute)
			if !c.Test.Done {
				c.Test.Complete("timeout", c.ID)
			}
		}()
		go crawlers.StartEngine(c, true)
		c.Test.WG.Wait()

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

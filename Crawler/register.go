package Crawler

import (
	"github.com/jpillora/go-tld"
	"megaCrawler/Crawler/config"
	"sync"
	"time"
)

var (
	WebMap    = make(map[string]*WebsiteEngine)
	nextTime  = time.Now().Add(3 * time.Second)
	timeMutex = sync.RWMutex{}
)

// Register 注册插件控制器
func Register(service string, name string, baseUrl string) *WebsiteEngine {
	k, err := tld.Parse(baseUrl)
	if err != nil {
		panic(err)
	}
	engine := NewEngine(service, *k)
	if c, ok := config.Configs[service]; !ok {
		engine.Config = &config.Config{
			Id:       service,
			LastIter: time.Time{},
			Disabled: false,
			Name:     name,
		}
		config.Configs[service] = *engine.Config
	} else {
		engine.Config = &c
	}
	go func() {
		timeMutex.Lock()
		engine.Scheduler.Every(168).Hour().StartAt(nextTime).Do(startEngine, engine)
		nextTime = nextTime.Add(1 * time.Minute)
		engine.Scheduler.StartAsync()
		timeMutex.Unlock()
	}()

	WebMap[service] = engine
	return engine
}

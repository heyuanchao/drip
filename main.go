package main

import (
	"drip/log"
	"time"
	"fmt"
)

func main() {
	var logger *log.Logger
	logger, err := log.New("debug", "D:", 16)
	if err != nil {
		panic(err)
	}
	log.Export(logger)
	// 测试模拟写入
	go func() {
		for {
			log.Debug("come on !")
			c:=time.NewTimer(500*time.Millisecond)
			<-c.C
		}
	}()
	// 切换新日志
	for {
		log.Debug("old")
		oldLogger := *logger
		newLogger := log.Change("debug", "D:", 16)
		if newLogger == nil {
			log.Debug("no new logger")
			newLogger = &oldLogger
		}
		logger = newLogger
		if newLogger == nil {
			oldLogger.Close()
		}
		oldLogger.Close()
		log.Export(logger)
		fmt.Println(logger)
		log.Debug("new")
		c := time.NewTimer(10 * time.Second)
		<-c.C
	}
	defer logger.Close()
}

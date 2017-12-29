package main

import (
	"drip/log"
	"os"
	"os/signal"
)

func main() {
	logger, err := log.New("debug", "D:", 16)
	if err != nil {
		panic(err)
	}
	log.Export(logger)

	defer logger.Close()

	log.Release("Drip starting up")
	//log.Export(log.Change("debug", "D:", 16))
	log.Release("Drip changed")
	// close
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	log.Release("Drip closing down (signal: %v)", sig)
}

package app

import (
	"github.com/kynmh69/gorutine-test/pkg/logging"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var counter int
var counter2 int

func Updater() {
	// update the app
	var wg sync.WaitGroup
	quit := make(chan struct{})
	logger := logging.GetLogger()
	logger.Infoln("start the updater")
	go updateToken(&wg, quit)
	go updateToken2(&wg, quit)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	close(quit)
	wg.Wait()
	logger.Infoln("stop the updater")
	logger.Infof("exit program: %d", counter)
}

func updateToken(wg *sync.WaitGroup, quit chan struct{}) {
	wg.Add(1)
	defer wg.Done()
	logger := logging.GetLogger()
	logger.Infoln("update the token")
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			counter++
			logger.Infof("update the token %d times", counter)
		case <-quit:
			logger.Infoln("quit the updater")
			return
		}
	}
}

func updateToken2(wg *sync.WaitGroup, quit chan struct{}) {
	wg.Add(1)
	defer wg.Done()
	logger := logging.GetLogger()
	logger.Infoln("update the token")
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			counter2++
			logger.Infof("update the token %d times", counter2)
		case <-quit:
			logger.Infoln("quit the updater")
			return
		}
	}
}

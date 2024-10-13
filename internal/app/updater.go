package app

import (
	"github.com/kynmh69/gorutine-test/pkg/logging"
	"os"
	"os/signal"
	"sync"
	"time"
)

var counter int

func Updater() {
	// update the app
	var wg sync.WaitGroup
	quit := make(chan struct{})
	logger := logging.GetLogger()
	logger.Infoln("start the updater")
	wg.Add(1)
	go updateToken(&wg, quit)
	c := make(chan os.Signal, 1)
	signal.Notify(c)
	<-c
	close(quit)
	wg.Wait()
	logger.Infoln("stop the updater")
	logger.Infof("exit program: %d\n", counter)
}

func updateToken(wg *sync.WaitGroup, quit chan struct{}) {
	defer wg.Done()
	logger := logging.GetLogger()
	logger.Infoln("update the token")
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			counter++
			logger.Infof("update the token %d times\n", counter)
		case <-quit:
			logger.Infoln("quit the updater")
			return
		}
	}
}

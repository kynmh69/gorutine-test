package app

import (
	"github.com/kynmh69/gorutine-test/pkg/logging"
	"sync"
	"time"
)

var counter int

func Updater(wg *sync.WaitGroup) {
	// update the app
	logger := logging.GetLogger()
	logger.Infoln("start the updater")
	wg.Add(1)
	go updateToken(wg)
}

func updateToken(wg *sync.WaitGroup) {
	defer wg.Done()
	logger := logging.GetLogger()
	logger.Infoln("update the token")
	time.Sleep(5 * time.Second)
	counter++
	logger.Infoln("update the token done")
}

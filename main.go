package main

import (
	"github.com/kynmh69/gorutine-test/internal/app"
	"github.com/kynmh69/gorutine-test/pkg/logging"
)

func init() {
	logging.InitLogger()
}

func main() {
	app.Updater()
}

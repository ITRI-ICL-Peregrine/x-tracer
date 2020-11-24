package main

import (
	"flag"
	"github.com/ITRI-ICL-Peregrine/x-tracer/database"
	"github.com/ITRI-ICL-Peregrine/x-tracer/events"
	"github.com/ITRI-ICL-Peregrine/x-tracer/pkg"
	"github.com/ITRI-ICL-Peregrine/x-tracer/ui"
)

func main() {

	database.Init()

	ui.SubscribeListeners()
	pkg.SubscribeListeners()

	go events.Run()

	port := flag.String("port", "6666", "")
	pkg.SetPort(*port)
	go pkg.StartServer()

	ui.InitGui()

}
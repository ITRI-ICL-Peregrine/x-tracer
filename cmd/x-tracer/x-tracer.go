package main

import (
	"flag"
	"github.com/ITRI-ICL-Peregrine/x-tracer/database"
	"github.com/ITRI-ICL-Peregrine/x-tracer/pkg"
	"github.com/ITRI-ICL-Peregrine/x-tracer/ui"

)

type LOG_MODE func(string, int64)

func main() {
	database.Init()
	port := flag.String("port", "6666", "")
	pkg.SetPort(*port)
	go pkg.StartServer()
	pkg.LOG_MODE = ui.RefreshLogs
	ui.InitGui()

}

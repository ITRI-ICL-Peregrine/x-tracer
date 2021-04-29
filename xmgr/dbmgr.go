package xmgr

import (
	"github.com/ITRI-ICL-Peregrine/x-tracer/database"
	m "github.com/ITRI-ICL-Peregrine/x-tracer/module"
	"os"

)


func DbInit(){

	database.Init()

}


func Storetcplogs(logs m.ReceiveLogEvent) {

	err := database.UpdateLogs(database.TcpLog(logs))
	if err != nil {
		os.Exit(1)
	}

}


func Getlogs(probe string) string {

	logs := database.GetActiveLogs(probe)
	return logs

}

package main

import (
	//"flag"
	//"github.com/ITRI-ICL-Peregrine/x-tracer/database"
	//"github.com/ITRI-ICL-Peregrine/x-tracer/pkg"
//"github.com/ITRI-ICL-Peregrine/x-tracer/ui"
"github.com/ITRI-ICL-Peregrine/x-tracer/xmgr"
"fmt"


)

//type LOG_MODE func(string, int64)

func main() {

	fmt.Println(xmgr.GetNamespaces())
//	fmt.Println(xmgr.GetPods("kube-public"))
	fmt.Println(xmgr.GetProbes())
	xmgr.TracePod("iperf3","tcpconnect")
	fmt.Println(xmgr.Getlogs("tcpconnect"))
	
	//database.Init()
	//port := flag.String("port", "6666", "")
	//pkg.SetPort(*port)
	//go pkg.StartServer()
	//pkg.LOG_MODE = ui.RefreshLogs
	//ui.InitGui()

}

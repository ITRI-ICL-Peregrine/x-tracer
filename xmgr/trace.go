package xmgr

import (

	"github.com/ITRI-ICL-Peregrine/x-tracer/kube"
	"github.com/ITRI-ICL-Peregrine/x-tracer/pkg"
	m "github.com/ITRI-ICL-Peregrine/x-tracer/module"
	"flag"
	"fmt"
	"time"
)

//var StreamLogs chan StoreLogs

func TracePod(pod string, probe string)(string){

	//Initializing Data Base
	DbInit()
	//Start server on port 6666
	port := flag.String("port", "6666", "")
	pkg.SetPort(*port)
	go pkg.StartServer()
	//Starting x-agent services and pod
	kube.StartAgent(pod, probe)

	//Waiting for the logs
	m.TcplogChan = make(chan m.ReceiveLogEvent, 1)
	go func() {

		for val := range m.TcplogChan {

		//	fmt.Println(val)
			Storetcplogs(val)
			fmt.Println(Getlogs("tcpconnect"))

		}
	}()

	for {

		time.Sleep(time.Duration(1) * time.Second)
	}

}

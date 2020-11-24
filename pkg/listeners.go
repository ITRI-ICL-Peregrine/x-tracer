package pkg

import (
	"github.com/ITRI-ICL-Peregrine/x-tracer/database"
	"github.com/ITRI-ICL-Peregrine/x-tracer/events"
	"os"
)

func receiveLog(e events.Event) {
	if e, ok := e.(events.ReceiveLogEvent); ok {

		tcp := events.ReceiveLogEvent{ProbeName: e.ProbeName,
			Sys_Time: e.Sys_Time,
			T:        e.T,
			Pid:      e.Pid,
			Pname:    e.Pname,
			Ip:       e.Ip,
			Saddr:    e.Saddr,
			Daddr:    e.Daddr,
			Dport:    e.Dport,
			Sport:    e.Sport,
		}
		tcplogs := database.TcpLog(tcp)

		err := database.UpdateLogs(tcplogs)
		if err != nil {

			os.Exit(1)
		}

		if Integ == true {

			events.PublishEvent("logs:refreshsingle", events.EmptyMessage{Pn: e.ProbeName})

		} else {

			events.PublishEvent("logs:refreshinteg", events.EmptyMessage{Pn: e.ProbeName})

		}

	}
}

func tcplifeLog(e events.Event) {
	if e, ok := e.(events.TcpLifeLogEvent); ok {

		tcp := events.TcpLifeLogEvent{TimeStamp: e.TimeStamp,
			ProbeName: e.ProbeName,
			Sys_Time:  e.Sys_Time,
			Pid:       e.Pid,
			Pname:     e.Pname,
			Laddr:     e.Laddr,
			Lport:     e.Lport,
			Raddr:     e.Raddr,
			Rport:     e.Rport,
			Tx_kb:     e.Tx_kb,
			Rx_kb:     e.Rx_kb,
			Ms:        e.Ms,
		}
		tcplogs := database.TcpLifeLog(tcp)
		err := database.UpdateTcpLifeLogs(tcplogs)
		if err != nil {

			os.Exit(1)
		}

		if Integ == true {

			events.PublishEvent("logs:refreshsingle", events.EmptyMessage{Pn: e.ProbeName})

		} else {

			events.PublishEvent("logs:refreshinteg", events.EmptyMessage{Pn: e.ProbeName})

		}

	}
}

func execsnoopLog(e events.Event) {
	if e, ok := e.(events.ExecSnoopLogEvent); ok {

		tcp := events.ExecSnoopLogEvent{TimeStamp: e.TimeStamp,
			ProbeName: e.ProbeName,
			Sys_Time:  e.Sys_Time,
			T:         e.T,
			Pname:     e.Pname,
			Pid:       e.Pid,
			Ppid:      e.Ppid,
			Ret:       e.Ret,
			Args:      e.Args,
		}
		eslogs := database.ExecSnoopLog(tcp)
		err := database.UpdateEsLogs(eslogs)
		if err != nil {

			os.Exit(1)
		}

		if Integ == true {

			events.PublishEvent("logs:refreshsingle", events.EmptyMessage{Pn: e.ProbeName})

		} else {

			events.PublishEvent("logs:refreshinteg", events.EmptyMessage{Pn: e.ProbeName})

		}

	}
}

func biosnoopLog(e events.Event) {
	if e, ok := e.(events.BioSnoopLogEvent); ok {

		tcp := events.BioSnoopLogEvent{TimeStamp: e.TimeStamp,
			ProbeName: e.ProbeName,
			Sys_Time:  e.Sys_Time,
			T:         e.T,
			Pname:     e.Pname,
			Pid:       e.Pid,
			Disk:      e.Disk,
			Rw:        e.Rw,
			Sector:    e.Sector,
			Bytes:     e.Bytes,
			Lat:       e.Lat,
		}
		bslogs := database.BioSnoopLog(tcp)
		err := database.UpdateBsLogs(bslogs)
		if err != nil {

			os.Exit(1)
		}
		if Integ == true {

			events.PublishEvent("logs:refreshsingle", events.EmptyMessage{Pn: e.ProbeName})

		} else {

			events.PublishEvent("logs:refreshinteg", events.EmptyMessage{Pn: e.ProbeName})

		}

	}
}

func cachestatLog(e events.Event) {
	if e, ok := e.(events.CacheStatLogEvent); ok {

		tcp := events.CacheStatLogEvent{TimeStamp: e.TimeStamp,
			ProbeName: e.ProbeName,
			Sys_Time:  e.Sys_Time,
			Pid:       e.Pid,
			Uid:       e.Uid,
			Cmd:       e.Cmd,
			Hits:      e.Hits,
			Miss:      e.Miss,
			Dirties:   e.Dirties,
			Read_hit:  e.Read_hit,
			Write_hit: e.Write_hit,
		}
		cslogs := database.CacheStatLog(tcp)
		err := database.UpdateCsLogs(cslogs)
		if err != nil {

			os.Exit(1)
		}
		if Integ == true {

			events.PublishEvent("logs:refreshsingle", events.EmptyMessage{Pn: e.ProbeName})

		} else {

			events.PublishEvent("logs:refreshinteg", events.EmptyMessage{Pn: e.ProbeName})

		}

	}
}

func SubscribeListeners() {
	events.Subscribe(tcplifeLog, "log:tcplife")
	events.Subscribe(receiveLog, "log:receive")
	events.Subscribe(execsnoopLog, "log:execsnoop")
	events.Subscribe(biosnoopLog, "log:biosnoop")
	events.Subscribe(cachestatLog, "log:cachestat")
}

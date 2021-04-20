package database

import (
	"fmt"
	"github.com/gogo/protobuf/sortkeys"
	"strings"
	"log"
)

var bufLogs []string
var wbLogs []string
var csbufLogs []string
var cswbLogs []string

//var bsbufLogs []string
//var bswbLogs []string
var esbufLogs []string
var eswbLogs []string
var tlbufLogs []string
var tlwbLogs []string



func GetActiveLogs(pn string) string {
	var err error

	var keys []int64

	if pn == "tcplife" {
		var tlLogs []string
		logs := GetTcpLifeLogs()

		if err != nil {
			log.Panic(err)
		}

		for k := range logs {
			keys = append(keys, k)

		}

		sortkeys.Int64s(keys)

		for _, log := range keys {
			val := logs[log]
			tlLogs = append(tlLogs, fmt.Sprintf("{Probe:%s |Sys_Time:%s | PID:%s | PNAME:%s | LADDR:%s | LPORT:%s | RADDR:%s | RPORT:%s | Tx_kb:%s | Rx_kb:%s | Ms: %s \n", val.ProbeName, val.Sys_Time, val.Pid, val.Pname, val.Laddr, val.Lport, val.Raddr, val.Rport, val.Tx_kb, val.Rx_kb, val.Ms))

		}

		for i := range tlLogs {
			tlbufLogs = append(tlbufLogs, tlLogs[i])
		}
		if len(tlbufLogs) >= 9 {

			tlwbLogs = tlbufLogs
			tlbufLogs = nil
			del := DeleteTlLogs()
			return strings.Join(tlwbLogs, "\n")
			fmt.Println(del)
		} else {

			return strings.Join(tlwbLogs, "\n")

		}

	} else if pn == "execsnoop" {
		var esLogs []string
		logs := GetExecSnoopLogs()

		if err != nil {
			log.Panic(err)
		}

		for k := range logs {

			keys = append(keys, k)

		}

		sortkeys.Int64s(keys)

		for _, log := range keys {
			val := logs[log]
			esLogs = append(esLogs, fmt.Sprintf("{Probe:%s |Sys_Time:%s | T:%s | PNAME:%s | PID:%s | PPID:%s | RET:%s | ARGS:%s \n", val.ProbeName, val.Sys_Time, val.T, val.Pname, val.Pid, val.Ppid, val.Ret, val.Args))

		}

		for i := range esLogs {
			esbufLogs = append(esbufLogs, esLogs[i])
		}
		if len(esbufLogs) >= 9 {

			eswbLogs = esbufLogs
			esbufLogs = nil
			del := DeleteESLogs()
			return strings.Join(eswbLogs, "\n")
			fmt.Println(del)
		} else {

			return strings.Join(eswbLogs, "\n")

		}

	} else if pn == "biosnoop" {
		var bsLogs []string
		logs := GetBioSnoopLogs()

		if err != nil {
			log.Panic(err)
		}

		for k := range logs {

			keys = append(keys, k)

		}

		sortkeys.Int64s(keys)

		for _, log := range keys {
			val := logs[log]
			bsLogs = append(bsLogs, fmt.Sprintf("{Probe:%s |Sys_Time:%s | T:%s | PNAME:%s | PID:%s | DISK:%s | R/W:%s | SECTOR:%s | BYTES:%s | LAT:%s \n", val.ProbeName, val.Sys_Time, val.T, val.Pname, val.Pid, val.Disk, val.Rw, val.Sector, val.Bytes, val.Lat))

		}
		return strings.Join(bsLogs, "\n")

	} else if pn == "cachestat" {
		var csLogs []string
		logs := GetCacheStatLogs()

		if err != nil {
			log.Panic(err)
		}

		for k := range logs {

			keys = append(keys, k)

		}

		sortkeys.Int64s(keys)

		for _, log := range keys {
			val := logs[log]
			csLogs = append(csLogs, fmt.Sprintf("{Probe:%s |Sys_Time:%s | PID:%s | UID:%s | CMD:%s | HITS:%s | MISS:%s | DIRTIES:%s | READ_HIT%:%s | WRITE_HIT%:%s \n", val.ProbeName, val.Sys_Time, val.Pid, val.Uid, val.Cmd, val.Hits, val.Miss, val.Dirties, val.Read_hit, val.Write_hit))

		}

		for i := range csLogs {
			csbufLogs = append(csbufLogs, csLogs[i])
		}
		if len(csbufLogs) >= 9 {

			cswbLogs = csbufLogs
			csbufLogs = nil
			del := DeleteCSLogs()
			return strings.Join(cswbLogs, "\n")
			fmt.Println(del)
		} else {

			return strings.Join(cswbLogs, "\n")

		}

	} else {
		var tcpLogs []string

		logs := GetLogs()

		if err != nil {
			log.Panic(err)
		}

		for k := range logs {

			keys = append(keys, k)

		}

		sortkeys.Int64s(keys)

		for _, log := range keys {
			val := logs[log]
			if val.ProbeName == "tcpconnect" {
				tcpLogs = append(tcpLogs, fmt.Sprintf("{Probe:%s |Sys_Time:%s |T:%s | PID:%s | PNAME:%s | IP:%s | SADDR:%s | DADDR:%s | DPORT:%s \n", val.ProbeName, val.Sys_Time, val.T, val.Pid, val.Pname, val.Ip, val.Saddr, val.Daddr, val.Dport))

			} else if val.ProbeName == "tcptracer" {
				tcpLogs = append(tcpLogs, fmt.Sprintf("{Probe:%s |Sys_Time:%s |T:%s | PID:%s | PNAME:%s | IP:%s | SADDR:%s | DADDR:%s | DPORT:%s | SPORT:%s \n", val.ProbeName, val.Sys_Time, val.T, val.Pid, val.Pname, val.Ip, val.Saddr, val.Daddr, val.Dport, val.Sport))

			} else if val.ProbeName == "tcpaccept" {
				tcpLogs = append(tcpLogs, fmt.Sprintf("{Probe:%s |Sys_Time:%s |T:%s | PID:%s | PNAME:%s | IP:%s | LADDR:%s | RADDR:%s | LPORT:%s |RPORT: %s \n", val.ProbeName, val.Sys_Time, val.T, val.Pid, val.Pname, val.Ip, val.Saddr, val.Daddr, val.Sport, val.Dport))
			}
		}

		for i := range tcpLogs {
			bufLogs = append(bufLogs, tcpLogs[i])
		}
		if len(bufLogs) >= 9 {

			wbLogs = bufLogs
			bufLogs = nil
			del := DeleteTcpLogs()
			return strings.Join(wbLogs, "\n")
			fmt.Println(del)
		} else {

			return strings.Join(wbLogs, "\n")

		}

	}

	return "Nothing yet"

}

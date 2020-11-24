package pkg

import (
	"fmt"
	pb "github.com/ITRI-ICL-Peregrine/x-tracer/api"
	"github.com/ITRI-ICL-Peregrine/x-tracer/database"
	"github.com/ITRI-ICL-Peregrine/x-tracer/events"
	"github.com/gogo/protobuf/sortkeys"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"strings"
)

type StreamServer struct {
	//port string
}

var (
	port string
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

func (s *StreamServer) RouteLog(stream pb.SentLog_RouteLogServer) error {
	for {
		r, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.Response{
				Res: "Stream closed",
			})
		}
		if err != nil {
			return err
		}

		parse := strings.Fields(string(r.Log))

		if r.ProbeName == "tcpconnect" {
			events.PublishEvent("log:receive", events.ReceiveLogEvent{ProbeName: r.ProbeName,
				Sys_Time: parse[0],
				T:        parse[1],
				Pid:      parse[3],
				Pname:    parse[4],
				Ip:       parse[5],
				Saddr:    parse[6],
				Daddr:    parse[7],
				Dport:    parse[8],
				Sport:    "0",
			})
		} else if r.ProbeName == "tcptracer" {
			events.PublishEvent("log:receive", events.ReceiveLogEvent{ProbeName: r.ProbeName,
				Sys_Time: parse[0],
				T:        parse[1],
				Pid:      parse[3],
				Pname:    parse[4],
				Ip:       parse[5],
				Saddr:    parse[6],
				Daddr:    parse[7],
				Dport:    parse[9],
				Sport:    parse[8],
			})
		} else if r.ProbeName == "tcpaccept" {
			events.PublishEvent("log:receive", events.ReceiveLogEvent{ProbeName: r.ProbeName,
				Sys_Time: parse[0],
				T:        parse[1],
				Pid:      parse[3],
				Pname:    parse[4],
				Ip:       parse[5],
				Saddr:    parse[8],
				Daddr:    parse[6],
				Dport:    parse[7],
				Sport:    parse[9],
			})
		} else if r.ProbeName == "tcplife" {

			events.PublishEvent("log:tcplife", events.TcpLifeLogEvent{TimeStamp: 0,
				ProbeName: r.ProbeName,
				Sys_Time:  parse[0],
				Pid:       parse[2],
				Pname:     parse[3],
				Laddr:     parse[4],
				Lport:     parse[5],
				Raddr:     parse[6],
				Rport:     parse[7],
				Tx_kb:     parse[8],
				Rx_kb:     parse[9],
				Ms:        parse[10],
			})
		} else if r.ProbeName == "execsnoop" {
			if len(parse) < 8 {
				events.PublishEvent("log:execsnoop", events.ExecSnoopLogEvent{TimeStamp: 0,
					ProbeName: r.ProbeName,
					Sys_Time:  parse[0],
					T:         parse[1],
					Pname:     parse[3],
					Pid:       parse[4],
					Ppid:      parse[5],
					Ret:       parse[6],
					Args:      parse[3],
				})

			} else {
				events.PublishEvent("log:execsnoop", events.ExecSnoopLogEvent{TimeStamp: 0,
					ProbeName: r.ProbeName,
					Sys_Time:  parse[0],
					T:         parse[1],
					Pname:     parse[3],
					Pid:       parse[4],
					Ppid:      parse[5],
					Ret:       parse[6],
					Args:      parse[7],
				})
			}
		} else if r.ProbeName == "biosnoop" {

			events.PublishEvent("log:biosnoop", events.BioSnoopLogEvent{TimeStamp: 0,
				ProbeName: r.ProbeName,
				Sys_Time:  parse[0],
				T:         parse[1],
				Pname:     parse[2],
				Pid:       parse[3],
				Disk:      parse[4],
				Rw:        parse[5],
				Sector:    parse[6],
				Bytes:     parse[7],
				Lat:       parse[9],
			})
		} else if r.ProbeName == "cachestat" {

			events.PublishEvent("log:cachestat", events.CacheStatLogEvent{TimeStamp: 0,
				ProbeName: r.ProbeName,
				Sys_Time:  parse[0],
				Pid:       parse[1],
				Uid:       parse[2],
				Cmd:       parse[3],
				Hits:      parse[5],
				Miss:      parse[6],
				Dirties:   parse[7],
				Read_hit:  parse[8],
				Write_hit: parse[9],
			})
		}

	}
}


func SetPort(sport string) {
	port = sport
}

func StartServer() {
	server := grpc.NewServer()
	pb.RegisterSentLogServer(server, &StreamServer{})

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalln("net.Listen error:", err)
	}

	_ = server.Serve(lis)
}

func GetActiveLogs(pn string) string {
	var err error

	var keys []int64

	if pn == "tcplife" {
		var tlLogs []string
		logs := database.GetTcpLifeLogs()

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
			del := database.DeleteTlLogs()
			return strings.Join(tlwbLogs, "\n")
			fmt.Println(del)
		} else {

			return strings.Join(tlwbLogs, "\n")

		}



	} else if pn == "execsnoop" {
		var esLogs []string
		logs := database.GetExecSnoopLogs()

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
			del := database.DeleteESLogs()
			return strings.Join(eswbLogs, "\n")
			fmt.Println(del)
		} else {

			return strings.Join(eswbLogs, "\n")

		}



	} else if pn == "biosnoop" {
		var bsLogs []string
		logs := database.GetBioSnoopLogs()

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
		logs := database.GetCacheStatLogs()

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
			del := database.DeleteCSLogs()
			return strings.Join(cswbLogs, "\n")
			fmt.Println(del)
		} else {

			return strings.Join(cswbLogs, "\n")

		}



	} else {
		var tcpLogs []string

		logs := database.GetLogs()

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
			del := database.DeleteTcpLogs()
			return strings.Join(wbLogs, "\n")
			fmt.Println(del)
		} else {

			return strings.Join(wbLogs, "\n")

		}

	}

	return "Nothing yet"

}

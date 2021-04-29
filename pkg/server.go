package pkg

import (

	pb "github.com/ITRI-ICL-Peregrine/x-tracer/api"
	"github.com/ITRI-ICL-Peregrine/x-tracer/database"
	"github.com/ITRI-ICL-Peregrine/x-tracer/datastruct"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
	"net"
	"strings"
)

type StreamServer struct {
	//port string
}

var (
	port      string
	LOG_MODE  func(string,int64)
)
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
			tcp := datastruct.ReceiveLogEvent{ProbeName: r.ProbeName,

				Sys_Time: parse[0],
				T:        parse[1],
				Pid:      parse[3],
				Pname:    parse[4],
				Ip:       parse[5],
				Saddr:    parse[6],
				Daddr:    parse[7],
				Dport:    parse[8],
				Sport:    "0",
			}
			tcplog := database.TcpLog(tcp)
			err := database.UpdateLogs(tcplog)
			if err != nil {
				os.Exit(1)
			}

			getLogs(LOG_MODE,r.ProbeName, r.Pid)




		} else if r.ProbeName == "tcptracer" {
			tcp := datastruct.ReceiveLogEvent{ProbeName: r.ProbeName,
				Sys_Time: parse[0],
				T:        parse[1],
				Pid:      parse[3],
				Pname:    parse[4],
				Ip:       parse[5],
				Saddr:    parse[6],
				Daddr:    parse[7],
				Dport:    parse[9],
				Sport:    parse[8],
			}
			tcplog := database.TcpLog(tcp)
			err := database.UpdateLogs(tcplog)
			if err != nil {
				os.Exit(1)
			}
			getLogs(LOG_MODE,r.ProbeName, r.Pid)

		} else if r.ProbeName == "tcpaccept" {
			tcp := datastruct.ReceiveLogEvent{ProbeName: r.ProbeName,
				Sys_Time: parse[0],
				T:        parse[1],
				Pid:      parse[3],
				Pname:    parse[4],
				Ip:       parse[5],
				Saddr:    parse[8],
				Daddr:    parse[6],
				Dport:    parse[7],
				Sport:    parse[9],
			}
			tcplog := database.TcpLog(tcp)
			err := database.UpdateLogs(tcplog)
			if err != nil {
				os.Exit(1)
			}
			getLogs(LOG_MODE,r.ProbeName, r.Pid)
		


		} else if r.ProbeName == "tcplife" {

			tllogs := datastruct.TcpLifeLogEvent{TimeStamp: 0,
				
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
			}

			tcplife := database.TcpLifeLog(tllogs)
			err := database.UpdateTcpLifeLogs(tcplife)
			if err != nil {
				os.Exit(1)
			}
			getLogs(LOG_MODE,r.ProbeName, r.Pid)
		} else if r.ProbeName == "execsnoop" {
			if len(parse) < 8 {
				eslogs := datastruct.ExecSnoopLogEvent{TimeStamp: 0,
					
					ProbeName: r.ProbeName,
					Sys_Time:  parse[0],
					T:         parse[1],
					Pname:     parse[3],
					Pid:       parse[4],
					Ppid:      parse[5],
					Ret:       parse[6],
					Args:      parse[3],
				}

				eslog := database.ExecSnoopLog(eslogs)
				err := database.UpdateEsLogs(eslog)
				if err != nil {
					os.Exit(1)
				}
				getLogs(LOG_MODE,r.ProbeName, r.Pid)
			} else {
				eslogs := datastruct.ExecSnoopLogEvent{TimeStamp: 0,
					
					ProbeName: r.ProbeName,
					Sys_Time:  parse[0],
					T:         parse[1],
					Pname:     parse[3],
					Pid:       parse[4],
					Ppid:      parse[5],
					Ret:       parse[6],
					Args:      parse[7],
				}

				eslog := database.ExecSnoopLog(eslogs)
				err := database.UpdateEsLogs(eslog)
				if err != nil {
					os.Exit(1)
				}
				getLogs(LOG_MODE,r.ProbeName, r.Pid)
		}


		} else if r.ProbeName == "biosnoop" {

			bslogs := datastruct.BioSnoopLogEvent{TimeStamp: 0,
				
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
			}

			bslog := database.BioSnoopLog(bslogs)
			err := database.UpdateBsLogs(bslog)
			if err != nil {
				os.Exit(1)
			}
			getLogs(LOG_MODE,r.ProbeName, r.Pid)

		} else if r.ProbeName == "cachestat" {

			cslogs := datastruct.CacheStatLogEvent{TimeStamp: 0,
				
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
			}
			cslog := database.CacheStatLog(cslogs)
			err := database.UpdateCsLogs(cslog)
			if err != nil {
				os.Exit(1)
			}
			getLogs(LOG_MODE,r.ProbeName, r.Pid)
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


func getLogs(logs func(string, int64), pn string, ps int64){

	logs(pn,ps)

}

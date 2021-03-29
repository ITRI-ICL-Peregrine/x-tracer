package pkg

import (
	"context"
	"fmt"
	pb "github.com/ITRI-ICL-Peregrine/x-tracer/api"
	pp "github.com/ITRI-ICL-Peregrine/x-tracer/parse/probeparser"
	"google.golang.org/grpc"
	"log"
	"time"
)

type StreamClient struct {
	port string
	ip   string
}

var (
	client    pb.SentLogClient
	probe_num int64
)

func New(servicePort string, masterIp string) *StreamClient {
	return &StreamClient{
		servicePort,
		masterIp}
}

func (c *StreamClient) StartClient(probename []string, pidList [][]string, ucmd string) {

	connect, err := grpc.Dial(c.ip+":"+c.port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}

	defer connect.Close()

	client = pb.NewSentLogClient(connect)

	if len(probename) > 4 {

		probe_num = 4
		logtcpconnect := make(chan pp.Log, 1)

		go pp.RunTcpconnect(probename[1], logtcpconnect, pidList[0][0])

		go func() {

			for val := range logtcpconnect {

				err = c.startLogStream(client, &pb.Log{
					Pid:       probe_num,
					ProbeName: val.Probe,
					Log:       val.Fulllog,
					TimeStamp: "TimeStamp",
				})
				if err != nil {
					log.Fatalf("startLogStream fail.err: %v", err)
				}
			}

		}()

		logtcptracer := make(chan pp.Log, 1)
		go pp.RunTcptracer(probename[0], logtcptracer, pidList[0][0])
		go func() {

			for val := range logtcptracer {
				log.Printf("logtcptracer")
				err = c.startLogStream(client, &pb.Log{
					Pid:       probe_num,
					ProbeName: val.Probe,
					Log:       val.Fulllog,
					TimeStamp: "TimeStamp",
				})
				if err != nil {
					log.Fatalf("startLogStream fail.err: %v", err)
				}

			}

		}()

		logtcpaccept := make(chan pp.Log, 1)
		go pp.RunTcpaccept(probename[2], logtcpaccept, pidList[0][0])
		go func() {

			for val := range logtcpaccept {
				err = c.startLogStream(client, &pb.Log{
					Pid:       probe_num,
					ProbeName: val.Probe,
					Log:       val.Fulllog,
					TimeStamp: "TimeStamp",
				})
				if err != nil {
					log.Fatalf("startLogStream fail.err: %v", err)
				}

			}

		}()

		logtcplife := make(chan pp.Log, 1)
		go pp.RunTcplife(probename[3], logtcplife, pidList[0][0])
		go func() {
			for val := range logtcplife {
				err = c.startLogStream(client, &pb.Log{
					Pid:       probe_num,
					ProbeName: val.Probe,
					Log:       val.Fulllog,
					TimeStamp: "TimeStamp",
				})
				if err != nil {
					log.Fatalf("startLogStream fail.err: %v", err)
				}

			}

		}()

		logexecsnoop := make(chan pp.Log, 1)
		go pp.RunExecsnoop(probename[4], logexecsnoop, pidList[0][0])
		go func() {

			for val := range logexecsnoop {
				err = c.startLogStream(client, &pb.Log{
					Pid:       probe_num,
					ProbeName: val.Probe,
					Log:       val.Fulllog,
					TimeStamp: "TimeStamp",
				})

				if err != nil {
					log.Fatalf("startLogStream fail.err: %v", err)
				}
			}

		}()

		logbiosnoop := make(chan pp.Log, 1)
		go pp.RunBiosnoop(probename[5], logbiosnoop, pidList[0][0])
		go func() {

			for val := range logbiosnoop {

				err = c.startLogStream(client, &pb.Log{
					Pid:       probe_num,
					ProbeName: val.Probe,
					Log:       val.Fulllog,
					TimeStamp: "TimeStamp",
				})

				if err != nil {
					log.Fatalf("startLogStream fail.err: %v", err)
				}
			}

		}()

		logcachetop := make(chan pp.Log, 1)
		go pp.RunCachetop(probename[6], logcachetop, pidList[0][0])
		go func() {

			for val := range logcachetop {
				err = c.startLogStream(client, &pb.Log{
					Pid:       probe_num,
					ProbeName: val.Probe,
					Log:       val.Fulllog,
					TimeStamp: "TimeStamp",
				})

				if err != nil {
					log.Fatalf("startLogStream fail.err: %v", err)
				}
			}

		}()

	} else if len(probename) == 3 {
		probe_num = 3
		logtcpconnect := make(chan pp.Log, 1)

		go pp.RunTcpconnect(probename[1], logtcpconnect, pidList[0][0])

		go func() {

			for val := range logtcpconnect {

				err = c.startLogStream(client, &pb.Log{
					Pid:       probe_num,
					ProbeName: val.Probe,
					Log:       val.Fulllog,
					TimeStamp: "TimeStamp",
				})
				if err != nil {
					log.Fatalf("startLogStream fail.err: %v", err)
				}
			}

		}()

		logtcptracer := make(chan pp.Log, 1)
		go pp.RunTcptracer(probename[0], logtcptracer, pidList[0][0])
		go func() {

			for val := range logtcptracer {
				log.Printf("logtcptracer")
				err = c.startLogStream(client, &pb.Log{
					Pid:       probe_num,
					ProbeName: val.Probe,
					Log:       val.Fulllog,
					TimeStamp: "TimeStamp",
				})
				if err != nil {
					log.Fatalf("startLogStream fail.err: %v", err)
				}

			}

		}()

		logtcpaccept := make(chan pp.Log, 1)
		go pp.RunTcpaccept(probename[2], logtcpaccept, pidList[0][0])
		go func() {

			for val := range logtcpaccept {
				err = c.startLogStream(client, &pb.Log{
					Pid:       probe_num,
					ProbeName: val.Probe,
					Log:       val.Fulllog,
					TimeStamp: "TimeStamp",
				})
				if err != nil {
					log.Fatalf("startLogStream fail.err: %v", err)
				}

			}

		}()

	}else if len(probename) == 4 {
			loguretprobe := make(chan pp.Log, 1)
			go pp.RunUretprobe(probename[0], loguretprobe, ucmd)
			go func() {

				for val := range loguretprobe {

					err = c.startLogStream(client, &pb.Log{
						Pid:       probe_num,
						ProbeName: val.Probe,
						Log:       val.Fulllog,
						TimeStamp: "TimeStamp",
					})
					if err != nil {
						log.Fatalf("startLogStream fail.err: %v", err)
					}

				}

			}()


			loguretcountprobe := make(chan pp.Log, 1)
			go pp.RunUretprobeCount(probename[1], loguretcountprobe, ucmd)
			go func() {

				for val := range loguretcountprobe {

					err = c.startLogStream(client, &pb.Log{
						Pid:       probe_num,
						ProbeName: val.Probe,
						Log:       val.Fulllog,
						TimeStamp: "TimeStamp",
					})
					if err != nil {
						log.Fatalf("startLogStream fail.err: %v", err)
					}

				}

			}()


			loguretfreqprobe := make(chan pp.Log, 1)
			go pp.RunUretprobeFreq(probename[2], loguretfreqprobe, ucmd)
			go func() {

				for val := range loguretfreqprobe {

					err = c.startLogStream(client, &pb.Log{
						Pid:       probe_num,
						ProbeName: val.Probe,
						Log:       val.Fulllog,
						TimeStamp: "TimeStamp",
					})
					if err != nil {
						log.Fatalf("startLogStream fail.err: %v", err)
					}

				}

			}()

	}else{
		probe_num = 1
		switch probename[0] {
		case "Retval":

			loguretprobe := make(chan pp.Log, 1)
			go pp.RunUretprobe(probename[0], loguretprobe, ucmd)
			go func() {

				for val := range loguretprobe {

					err = c.startLogStream(client, &pb.Log{
						Pid:       probe_num,
						ProbeName: val.Probe,
						Log:       val.Fulllog,
						TimeStamp: "TimeStamp",
					})
					if err != nil {
						log.Fatalf("startLogStream fail.err: %v", err)
					}

				}

			}()

		case "Count":
			fmt.Println("At Count")
			loguretcountprobe := make(chan pp.Log, 1)
			go pp.RunUretprobeCount(probename[0], loguretcountprobe, ucmd)
			go func() {

				for val := range loguretcountprobe {
						err = c.startLogStream(client, &pb.Log{
						Pid:       probe_num,
						ProbeName: val.Probe,
						Log:       val.Fulllog,
						TimeStamp: "TimeStamp",
					})
					if err != nil {
						log.Fatalf("startLogStream fail.err: %v", err)
					}

				}

			}()

		case "Time":

			loguretfreqprobe := make(chan pp.Log, 1)
			go pp.RunUretprobeFreq(probename[0], loguretfreqprobe, ucmd)
			go func() {

				for val := range loguretfreqprobe {

					err = c.startLogStream(client, &pb.Log{
						Pid:       probe_num,
						ProbeName: val.Probe,
						Log:       val.Fulllog,
						TimeStamp: "TimeStamp",
					})
					if err != nil {
						log.Fatalf("startLogStream fail.err: %v", err)
					}

				}

			}()


		case "tcptracer":

			logtcptracer := make(chan pp.Log, 1)
			go pp.RunTcptracer(probename[0], logtcptracer, pidList[0][0])
			go func() {

				for val := range logtcptracer {
					err = c.startLogStream(client, &pb.Log{
						Pid:       probe_num,
						ProbeName: val.Probe,
						Log:       val.Fulllog,
						TimeStamp: "TimeStamp",
					})
					if err != nil {
						log.Fatalf("startLogStream fail.err: %v", err)
					}

				}

			}()

		case "tcpconnect":
			logtcpconnect := make(chan pp.Log, 1)

			go pp.RunTcpconnect(probename[0], logtcpconnect, pidList[0][0])

			go func() {

				for val := range logtcpconnect {

					err = c.startLogStream(client, &pb.Log{
						Pid:       probe_num,
						ProbeName: val.Probe,
						Log:       val.Fulllog,
						TimeStamp: "TimeStamp",
					})
					if err != nil {
						log.Fatalf("startLogStream fail.err: %v", err)
					}

				}

			}()

		case "tcpaccept":

			logtcpaccept := make(chan pp.Log, 1)
			go pp.RunTcpaccept(probename[0], logtcpaccept, pidList[0][0])
			go func() {

				for val := range logtcpaccept {
					err = c.startLogStream(client, &pb.Log{
						Pid:       probe_num,
						ProbeName: val.Probe,
						Log:       val.Fulllog,
						TimeStamp: "TimeStamp",
					})
					if err != nil {
						log.Fatalf("startLogStream fail.err: %v", err)
					}

				}

			}()

		case "tcplife":

			logtcplife := make(chan pp.Log, 1)
			go pp.RunTcplife(probename[0], logtcplife, pidList[0][0])
			go func() {

				for val := range logtcplife {

					err = c.startLogStream(client, &pb.Log{
						Pid:       probe_num,
						ProbeName: val.Probe,
						Log:       val.Fulllog,
						TimeStamp: "TimeStamp",
					})
					if err != nil {
						log.Fatalf("startLogStream fail.err: %v", err)
					}

				}

			}()

		case "execsnoop":

			logexecsnoop := make(chan pp.Log, 1)
			go pp.RunExecsnoop(probename[0], logexecsnoop, pidList[0][0])
			go func() {

				for val := range logexecsnoop {
					fmt.Println("execsnoop", val.Fulllog)
					err = c.startLogStream(client, &pb.Log{
						Pid:       probe_num,
						ProbeName: val.Probe,
						Log:       val.Fulllog,
						TimeStamp: "TimeStamp",
					})
					if err != nil {
						log.Fatalf("startLogStream fail.err: %v", err)
					}

				}

			}()
		case "biosnoop":

			logbiosnoop := make(chan pp.Log, 1)
			go pp.RunBiosnoop(probename[0], logbiosnoop, pidList[0][0])
			go func() {

				for val := range logbiosnoop {

					err = c.startLogStream(client, &pb.Log{
						Pid:       probe_num,
						ProbeName: val.Probe,
						Log:       val.Fulllog,
						TimeStamp: "TimeStamp",
					})
					if err != nil {
						log.Fatalf("startLogStream fail.err: %v", err)
					}

				}

			}()
		case "cachestat":

			logcachetop := make(chan pp.Log, 1)
			go pp.RunCachetop(probename[0], logcachetop, pidList[0][0])
			go func() {

				for val := range logcachetop {

					err = c.startLogStream(client, &pb.Log{
						Pid:       probe_num,
						ProbeName: val.Probe,
						Log:       val.Fulllog,
						TimeStamp: "TimeStamp",
					})
					if err != nil {
						log.Fatalf("startLogStream fail.err: %v", err)
					}

				}

			}()

		}

	}
	for {

		time.Sleep(time.Duration(1) * time.Second)
	}

}

func (c *StreamClient) startLogStream(client pb.SentLogClient, r *pb.Log) error {

	stream, err := client.RouteLog(context.Background())
	if err != nil {
		return err
	}

	err = stream.Send(r)
	if err != nil {
		return err
	}
	resp, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}

	log.Printf("Response from the Server ;) : %v", resp.Res)
	return nil
}

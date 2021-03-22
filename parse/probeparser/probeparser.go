package probeparser

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

type Log struct {
	Fulllog string
	Pid     int64
	Time    float64
	Probe   string
}

const (
	timestamp int = 0
)

func GetNS(pid string) string {
	cmdName := "ls"
	out, err := exec.Command(cmdName, fmt.Sprintf("/proc/%s/ns/net", pid), "-al").Output()
	if err != nil {
		println(err)
	}
	ns := string(out)
	parse := strings.Fields(string(ns))
	s := strings.SplitN(parse[10], "[", 10)
	sep := strings.Split(s[1], "]")
	return sep[0]

}

func RunUretprobeFreq(tool string, loguretprobe chan Log, pid string, filepath string, funcname string) {


	command := `uprobe:` + filepath + `:` + funcname + `{ @start = nsecs; }` + `uretprobe:` + filepath + `:` + funcname + `/@start/ { @time = ((nsecs - @start)/1000);  print(@time); delete(@start); }`
	cmd := exec.Command("bpftrace", "-p", pid , "-e ", command)
	//cmd.Dir = "/usr/share/bcc/tools/ebpf"
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Start()
	buf := bufio.NewReader(stdout)

	for {

		line, _, _ := buf.ReadLine()
		/*parsedLine := strings.Fields(string(line))

		if parsedLine[0] != "Attaching 1 Probe..." {
			ppid, err := strconv.ParseInt(parsedLine[0], 10, 64)
				if err != nil {
					println("Uretprobe PID Error")
				}*/

				timest := 0.00
				n := Log{Fulllog: string(line), Pid: 1234, Time: timest, Probe: tool}
				loguretprobe <- n

		//}
	}
}


func RunUretprobeCount(tool string, loguretprobe chan Log, pid string, filepath string, funcname string) {


	command  := `uretprobe:` + filepath + `:` + funcname + `{ @[comm,pid,retval] = count(); }` +`interval:s:1` + `{ time("%H:%M:%S"); print(@); clear(@); }`
	cmd := exec.Command("bpftrace", "-p", pid , "-e ", command)
	//cmd.Dir = "/usr/share/bcc/tools/ebpf"
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Start()
	buf := bufio.NewReader(stdout)

	for {

		line, _, _ := buf.ReadLine()
		/*parsedLine := strings.Fields(string(line))

		if parsedLine[0] != "Attaching 1 Probe..." {
			ppid, err := strconv.ParseInt(parsedLine[0], 10, 64)
				if err != nil {
					println("Uretprobe PID Error")
				}*/

				timest := 0.00
				n := Log{Fulllog: string(line), Pid: 1234, Time: timest, Probe: tool}
				loguretprobe <- n

		//}
	}
}


func RunUretprobe(tool string, loguretprobe chan Log, pid string, filepath string, funcname string) {


	command := `uretprobe:` + filepath + `:` + funcname + `{ printf("Pid:%d    RetValue:%d\n", pid, retval); }`
	cmd := exec.Command("bpftrace", "-p", pid , "-e ", command)
	//cmd.Dir = "/usr/share/bcc/tools/ebpf"
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Start()
	buf := bufio.NewReader(stdout)

	for {

		line, _, _ := buf.ReadLine()
		/*parsedLine := strings.Fields(string(line))

		if parsedLine[0] != "Attaching 1 Probe..." {
			ppid, err := strconv.ParseInt(parsedLine[0], 10, 64)
				if err != nil {
					println("Uretprobe PID Error")
				}*/

				timest := 0.00
				n := Log{Fulllog: string(line), Pid: 1234, Time: timest, Probe: tool}
				loguretprobe <- n

		//}
	}
}

func RunTcptracer(tool string, logtcptracer chan Log, pid string) {

	sep := GetNS(pid)

	cmd := exec.Command("./tcptracer.py", "-T", "-t", "-N"+sep)
	cmd.Dir = "/usr/share/bcc/tools/ebpf"
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Start()
	buf := bufio.NewReader(stdout)

	for {

		line, _, _ := buf.ReadLine()
		parsedLine := strings.Fields(string(line))

		if parsedLine[0] != "Tracing" {
			if parsedLine[0] != "TIME(s)" {
				ppid, err := strconv.ParseInt(parsedLine[3], 10, 64)
				if err != nil {
					println("Tcptracer PID Error")
				}

				timest := 0.00
				n := Log{Fulllog: string(line), Pid: ppid, Time: timest, Probe: tool}
				logtcptracer <- n

			}
		}
	}
}

func RunTcpconnect(tool string, logtcpconnect chan Log, pid string) {

	sep := GetNS(pid)
	cmd := exec.Command("./tcpconnect.py", "-T", "-t", "-N"+sep)
	cmd.Dir = "/usr/share/bcc/tools/ebpf"
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Start()
	buf := bufio.NewReader(stdout)

	for {
		line, _, _ := buf.ReadLine()
		parsedLine := strings.Fields(string(line))
		if parsedLine[0] != "TIME(s)" {
			ppid, err := strconv.ParseInt(parsedLine[3], 10, 64)
			if err != nil {
				println("TCPConnect PID Error")
			}

			timest := 0.00
			n := Log{Fulllog: string(line), Pid: ppid, Time: timest, Probe: tool}
			logtcpconnect <- n

		}
	}
}

func RunTcpaccept(tool string, logtcpaccept chan Log, pid string) {

	sep := GetNS(pid)
	cmd := exec.Command("./tcpaccept.py", "-T", "-t", "-N"+sep)
	cmd.Dir = "/usr/share/bcc/tools/ebpf"
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Start()
	buf := bufio.NewReader(stdout)

	for {
		line, _, _ := buf.ReadLine()
		parsedLine := strings.Fields(string(line))

		if parsedLine[0] != "TIME(s)" {
			ppid, err := strconv.ParseInt(parsedLine[3], 10, 64)
			if err != nil {
				println("TCPaccept PID Error")
			}

			timest := 0.00

			n := Log{Fulllog: string(line), Pid: ppid, Time: timest, Probe: tool}
			logtcpaccept <- n

		}
	}
}

func RunTcplife(tool string, logtcplife chan Log, pid string) {

	sep := GetNS(pid)
	cmd := exec.Command("./tcplife.py", "-T", "-N"+sep)
	cmd.Dir = "/usr/share/bcc/tools/ebpf"
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Start()
	buf := bufio.NewReader(stdout)

	for {
		line, _, _ := buf.ReadLine()
		parsedLine := strings.Fields(string(line))
		println(parsedLine[0])
		if parsedLine[0] != "TIME(s)" {
			ppid, err := strconv.ParseInt(parsedLine[2], 10, 64)
			if err != nil {
				println("TCPlife PID Error")
			}

			timest := 0.00

			n := Log{Fulllog: string(line), Pid: ppid, Time: timest, Probe: tool}
			logtcplife <- n

		}
	}
}

func RunExecsnoop(tool string, logexecsnoop chan Log, pid string) {

	sep := GetNS(pid)
	cmd := exec.Command("./execsnoop.py", "-T", "-t", "-N"+sep)
	cmd.Dir = "/usr/share/bcc/tools/ebpf"
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Start()
	buf := bufio.NewReader(stdout)

	for {
		line, _, _ := buf.ReadLine()
		parsedLine := strings.Fields(string(line))

		ppid, err := strconv.ParseInt(parsedLine[4], 10, 64)
		if err != nil {
			println("Execsnoop PID Error")
		}

		timest := 0.00

		n := Log{Fulllog: string(line), Pid: ppid, Time: timest, Probe: tool}
		logexecsnoop <- n

	}
}

func RunBiosnoop(tool string, logbiosnoop chan Log, pid string) {

	sep := GetNS(pid)
	cmd := exec.Command("./biosnoop.py", "-T", "-N"+sep)
	cmd.Dir = "/usr/share/bcc/tools/ebpf"
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Start()
	buf := bufio.NewReader(stdout)

	for {
		line, _, _ := buf.ReadLine()
		parsedLine := strings.Fields(string(line))

		ppid, err := strconv.ParseInt(parsedLine[3], 10, 64)
		if err != nil {
			println("Biosnoop PID Error")
		}
		timest := 0.00

		n := Log{Fulllog: string(line), Pid: ppid, Time: timest, Probe: tool}
		logbiosnoop <- n

	}
}

func RunCachetop(tool string, logcachetop chan Log, pid string) {

	sep := GetNS(pid)
	cmd := exec.Command("./Cachetop.py", "-T", "-N"+sep)
	cmd.Dir = "/usr/share/bcc/tools/ebpf"
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Start()
	buf := bufio.NewReader(stdout)

	for {
		line, _, _ := buf.ReadLine()
		parsedLine := strings.Fields(string(line))

		ppid, err := strconv.ParseInt(parsedLine[1], 10, 64)
		if err != nil {
			println("Cachetop PID Error")
		}
		timest := 0.00

		n := Log{Fulllog: string(line), Pid: ppid, Time: timest, Probe: tool}
		logcachetop <- n

	}
}

package probeparser

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type Log struct {
	Fulllog string
	Pid     string
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
				timest := 0.00
				n := Log{Fulllog: string(line), Pid: parsedLine[3], Time: timest, Probe: tool}
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
		if parsedLine[0] != "SYS_TIME" {

			ppid := parsedLine[3]
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
			timest := 0.00

			n := Log{Fulllog: string(line), Pid: parsedLine[3], Time: timest, Probe: tool}
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

		if parsedLine[0] != "TIME(s)" {
			timest := 0.00

			n := Log{Fulllog: string(line), Pid: parsedLine[2], Time: timest, Probe: tool}
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

		timest := 0.00

		n := Log{Fulllog: string(line), Pid: parsedLine[4], Time: timest, Probe: tool}
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

		timest := 0.00

		n := Log{Fulllog: string(line), Pid: parsedLine[3], Time: timest, Probe: tool}
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

		timest := 0.00

		n := Log{Fulllog: string(line), Pid: parsedLine[1], Time: timest, Probe: tool}
		logcachetop <- n

	}
}

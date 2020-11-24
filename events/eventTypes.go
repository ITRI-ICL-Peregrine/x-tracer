package events

type EmptyMessage struct {
	Pn string
}

type ReceiveLogEvent struct {
	ProbeName string
	Sys_Time  string
	T         string
	Pid       string
	Pname     string
	Ip        string
	Saddr     string
	Daddr     string
	Dport     string
	Sport     string
}

type TcpLifeLogEvent struct {
	TimeStamp int64
	ProbeName string
	Sys_Time  string
	Pid       string
	Pname     string
	Laddr     string
	Lport     string
	Raddr     string
	Rport     string
	Tx_kb     string
	Rx_kb     string
	Ms        string
}

type ExecSnoopLogEvent struct {
	TimeStamp int64
	ProbeName string
	Sys_Time  string
	T         string
	Pname     string
	Pid       string
	Ppid      string
	Ret       string
	Args      string
}

type BioSnoopLogEvent struct {
	TimeStamp int64
	ProbeName string
	Sys_Time  string
	T         string
	Pname     string
	Pid       string
	Disk      string
	Rw        string
	Sector    string
	Bytes     string
	Lat       string
}

type CacheStatLogEvent struct {
	TimeStamp int64
	ProbeName string
	Sys_Time  string
	Pid       string
	Uid       string
	Cmd       string
	Hits      string
	Miss      string
	Dirties   string
	Read_hit  string
	Write_hit string
}


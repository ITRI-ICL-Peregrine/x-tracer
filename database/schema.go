package database

type Log struct {
	Timestamp int64
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

type TcpLifeLog struct {
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

type TcpLog struct {
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

type ExecSnoopLog struct {
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

type BioSnoopLog struct {
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

type CacheStatLog struct {
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

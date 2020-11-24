package database

import (
	"fmt"
	memdb "github.com/hashicorp/go-memdb"
	"time"
)

var (
	db   *memdb.MemDB
	tldb *memdb.MemDB
	es   *memdb.MemDB
	bs   *memdb.MemDB
	cs   *memdb.MemDB
)

func Init() {
	var err error
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"tcpconnect": {
				Name: "tcpconnect",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.IntFieldIndex{Field: "Timestamp"},
					},
					"pn": {
						Name:    "pn",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "ProbeName"},
					},
					"sys_time": {
						Name:    "sys_time",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Sys_Time"},
					},

					"t": {
						Name:    "t",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "T"},
					},
					"pid": {
						Name:    "pid",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Pid"},
					},

					"pname": {
						Name:    "pname",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Pname"},
					},
					"ip": {
						Name:    "ip",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Ip"},
					},

					"saddr": {
						Name:    "saddr",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Saddr"},
					},
					"daddr": {
						Name:    "daddr",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Daddr"},
					},
					"dport": {
						Name:    "dport",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Dport"},
					},
					"sport": {
						Name:    "sport",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Sport"},
					},
				},
			},
		},
	}

	schematcplife := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"tcplife": {
				Name: "tcplife",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.IntFieldIndex{Field: "TimeStamp"},
					},
					"pn": {
						Name:    "pn",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "ProbeName"},
					},
					"sys_time": {
						Name:    "sys_time",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Sys_Time"},
					},

					"pid": {
						Name:    "pid",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Pid"},
					},

					"pname": {
						Name:    "pname",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Pname"},
					},

					"laddr": {
						Name:    "laddr",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Laddr"},
					},
					"lport": {
						Name:    "lport",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Lport"},
					},
					"raddr": {
						Name:    "raddr",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Raddr"},
					},
					"rport": {
						Name:    "rport",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Rport"},
					},

					"tx_kb": {
						Name:    "tx_kb",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Tx_kb"},
					},

					"rx_kb": {
						Name:    "rx_kb",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Rx_kb"},
					},

					"ms": {
						Name:    "ms",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Ms"},
					},
				},
			},
		},
	}

	//Schema for Execsnoop
	schemaes := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"execsnoop": {
				Name: "execsnoop",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.IntFieldIndex{Field: "TimeStamp"},
					},
					"pn": {
						Name:    "pn",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "ProbeName"},
					},
					"sys_time": {
						Name:    "sys_time",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Sys_Time"},
					},

					"pname": {
						Name:    "pname",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Pname"},
					},

					"pid": {
						Name:    "pid",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Pid"},
					},

					"ppid": {
						Name:    "ppid",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Ppid"},
					},
					"ret": {
						Name:    "ret",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Ret"},
					},
					"args": {
						Name:    "args",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Args"},
					},
				},
			},
		},
	}

	//Schema for Biosnoop

	schemabs := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"biosnoop": {
				Name: "biosnoop",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.IntFieldIndex{Field: "TimeStamp"},
					},
					"pn": {
						Name:    "pn",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "ProbeName"},
					},
					"sys_time": {
						Name:    "sys_time",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Sys_Time"},
					},

					"t": {
						Name:    "t",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "T"},
					},
					"pname": {
						Name:    "pname",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Pname"},
					},

					"pid": {
						Name:    "pid",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Pid"},
					},

					"disk": {
						Name:    "disk",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Disk"},
					},
					"rw": {
						Name:    "rw",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Rw"},
					},

					"sector": {
						Name:    "sector",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Sector"},
					},
					"bytes": {
						Name:    "bytes",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Bytes"},
					},
					"lat": {
						Name:    "lat",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Lat"},
					},
				},
			},
		},
	}

	//Schema for Cachestat

	schemacs := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"cachestat": {
				Name: "cachestat",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.IntFieldIndex{Field: "TimeStamp"},
					},
					"pn": {
						Name:    "pn",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "ProbeName"},
					},
					"sys_time": {
						Name:    "sys_time",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Sys_Time"},
					},

					"pid": {
						Name:    "pid",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Pid"},
					},

					"uid": {
						Name:    "uid",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Uid"},
					},
					"cmd": {
						Name:    "cmd",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Cmd"},
					},

					"hits": {
						Name:    "hits",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Hits"},
					},
					"miss": {
						Name:    "miss",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Miss"},
					},
					"dirties": {
						Name:    "dirties",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Dirties"},
					},
					"rh": {
						Name:    "rh",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Read_hit"},
					},
					"wh": {
						Name:    "wh",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Write_hit"},
					},
				},
			},
		},
	}

	//Create a new data base for tcplogs
	db, err = memdb.NewMemDB(schema)
	if err != nil {
		panic(err)
	}

	//Create a new data base for tcplife
	tldb, err = memdb.NewMemDB(schematcplife)
	if err != nil {
		panic(err)
	}

	//Create a new data base for execsnoop
	es, err = memdb.NewMemDB(schemaes)
	if err != nil {
		panic(err)
	}

	//Create a new data base for biosnoop
	bs, err = memdb.NewMemDB(schemabs)
	if err != nil {
		panic(err)
	}

	//Create a new data base for cachestat
	cs, err = memdb.NewMemDB(schemacs)
	if err != nil {
		panic(err)
	}

}


func UpdateLogs(log TcpLog) error {

	txn := db.Txn(true)
	timestamp := time.Now().UnixNano()
	logs := []*Log{

		{timestamp, log.ProbeName, log.Sys_Time, log.T, log.Pid, log.Pname, log.Ip, log.Saddr, log.Daddr, log.Dport, log.Sport},
	}

	for _, p := range logs {
		if err := txn.Insert("tcpconnect", p); err != nil {
			return err
		}
	}

	txn.Commit()

	return nil

}

func UpdateTcpLifeLogs(log TcpLifeLog) error {

	txn := tldb.Txn(true)
	timestamp := time.Now().UnixNano()
	logs := []*TcpLifeLog{

		{timestamp, log.ProbeName, log.Sys_Time, log.Pid, log.Pname, log.Laddr, log.Lport, log.Raddr, log.Rport, log.Tx_kb, log.Rx_kb, log.Ms},
	}

	for _, p := range logs {
		if err := txn.Insert("tcplife", p); err != nil {

			return err
		}
	}

	txn.Commit()

	return nil

}

//update execsnoop table
func UpdateEsLogs(log ExecSnoopLog) error {

	txn := es.Txn(true)
	timestamp := time.Now().UnixNano()
	logs := []*ExecSnoopLog{

		{timestamp, log.ProbeName, log.Sys_Time, log.T, log.Pname, log.Pid, log.Ppid, log.Ret, log.Args},
	}

	for _, p := range logs {
		if err := txn.Insert("execsnoop", p); err != nil {
			return err
		}
	}

	txn.Commit()

	return nil

}

//update biosnoop  table
func UpdateBsLogs(log BioSnoopLog) error {

	txn := bs.Txn(true)
	timestamp := time.Now().UnixNano()
	logs := []*BioSnoopLog{

		{timestamp, log.ProbeName, log.Sys_Time, log.T, log.Pname, log.Pid, log.Disk, log.Rw, log.Sector, log.Bytes, log.Lat},
	}

	for _, p := range logs {
		if err := txn.Insert("biosnoop", p); err != nil {
			return err
		}
	}

	txn.Commit()

	return nil

}

//update cachestat table
func UpdateCsLogs(log CacheStatLog) error {

	txn := cs.Txn(true)
	timestamp := time.Now().UnixNano()
	logs := []*CacheStatLog{

		{timestamp, log.ProbeName, log.Sys_Time, log.Pid, log.Uid, log.Cmd, log.Hits, log.Miss, log.Dirties, log.Read_hit, log.Write_hit},
	}

	for _, p := range logs {
		if err := txn.Insert("cachestat", p); err != nil {
			return err
		}
	}

	txn.Commit()

	return nil

}

func GetLogs() map[int64]*Log {

	txn := db.Txn(false)
	defer txn.Abort()

	logs := make(map[int64]*Log)

	it, err := txn.Get("tcpconnect", "id")
	if err != nil {
		panic(err)
	}



	for obj := it.Next(); obj != nil; obj = it.Next() {
		p := obj.(*Log)
		timestamp := p.Timestamp
		logs[timestamp] = p

	}

	return logs
}

func GetTcpLifeLogs() map[int64]*TcpLifeLog {

	txn := tldb.Txn(false)
	defer txn.Abort()

	logs := make(map[int64]*TcpLifeLog)

	it, err := txn.Get("tcplife", "id")
	if err != nil {
		panic(err)
	}



	for obj := it.Next(); obj != nil; obj = it.Next() {
		p := obj.(*TcpLifeLog)
		timestamp := p.TimeStamp
		logs[timestamp] = p

	}

	return logs
}

//Get execsnoop logs

func GetExecSnoopLogs() map[int64]*ExecSnoopLog {

	txn := es.Txn(false)
	defer txn.Abort()

	logs := make(map[int64]*ExecSnoopLog)

	it, err := txn.Get("execsnoop", "id")
	if err != nil {
		panic(err)
		//os.Exit(1)
	}



	for obj := it.Next(); obj != nil; obj = it.Next() {
		p := obj.(*ExecSnoopLog)
		timestamp := p.TimeStamp
		logs[timestamp] = p

	}

	return logs

}

//Get Biosnoop logs

func GetBioSnoopLogs() map[int64]*BioSnoopLog {

	txn := bs.Txn(false)
	defer txn.Abort()

	logs := make(map[int64]*BioSnoopLog)

	it, err := txn.Get("biosnoop", "id")
	if err != nil {
		panic(err)
	}



	for obj := it.Next(); obj != nil; obj = it.Next() {
		p := obj.(*BioSnoopLog)
		timestamp := p.TimeStamp
		logs[timestamp] = p

	}

	return logs
}

//Get Cachestat logs

func GetCacheStatLogs() map[int64]*CacheStatLog {

	txn := cs.Txn(false)
	defer txn.Abort()

	logs := make(map[int64]*CacheStatLog)

	it, err := txn.Get("cachestat", "id")
	if err != nil {
		panic(err)
	}



	for obj := it.Next(); obj != nil; obj = it.Next() {
		p := obj.(*CacheStatLog)
		timestamp := p.TimeStamp
		logs[timestamp] = p

	}

	return logs
}

func DeleteTcpLogs() int {

	txn := db.Txn(true)

	del, err := txn.DeleteAll("tcpconnect", "id")
	if err != nil {
		panic(err)
		return 0
	}

	txn.Commit()

	return del

}

func DeleteTlLogs() int {

	txn := tldb.Txn(true)

	del, err := txn.DeleteAll("tcplife", "id")
	if err != nil {
		fmt.Println("TCPLOGS DELETION ERROR")

		return 0
	}

	txn.Commit()

	return del

}

func DeleteCSLogs() int {

	txn := cs.Txn(true)

	del, err := txn.DeleteAll("cachestat", "id")
	if err != nil {

		return 0
		panic(err)
	}

	txn.Commit()

	return del

}

func DeleteESLogs() int {

	txn := es.Txn(true)

	del, err := txn.DeleteAll("execsnoop", "id")
	if err != nil {
		panic(err)
	}

	txn.Commit()

	return del

}


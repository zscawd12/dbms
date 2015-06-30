package main

import (
	"flag"
	"fmt"
	"github.com/zscawd12/dbms/internal/query"
	"log"
	"os"
	"runtime"
)

const (
	path    = "/usr/local/etc/dbms"
	logPath = "/var/log/dbms.log"
)

func initLogger(logFile string) *log.Logger {
	f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open log file %v: %v\n", logFile, err)
		os.Exit(1)
	}

	return log.New(f, "", log.Ldate|log.Ltime)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()

	logger := initLogger(logPath)
	dbms := query.New(logger, path)

	var fir string
	fmt.Scanf("%s", &fir)

	switch fir {
	case "create":
		err := dbms.CreateTable()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create table", err)
			return
		}
		break

	case "insert":
		err := dbms.Insert()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to query insert", err)
			return
		}
		break

	default:
		fmt.Fprintf(os.Stderr, "no menu\n")

	}

}

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

	var input string
	for {
		fmt.Print("\n\nmenu : create, insert, select, quit\n")

		fmt.Scanf("%s", &input)

		switch input {
		case "create":
			err := dbms.CreateTable()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to create table", err)
				return
			}

			fmt.Println("테이블이 성공적으로 생성되었습니다.")
			break

		case "insert":
			err := dbms.Insert()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to insert query", err)
				return
			}

			fmt.Println("튜플이 성공적으로 저장되었습니다.")
			break

		case "select":
			err := dbms.Select()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to select query", err)
				return
			}

			fmt.Println("데이터 조회 완료!!")
			break

		case "quit":
			fmt.Print("Bye\n")
			os.Exit(0)

		default:
			fmt.Fprintf(os.Stderr, "no menu\n")

		}
	}
}

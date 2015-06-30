package query

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type DBMS struct {
	Logger   *log.Logger
	FilePath string
}

func New(logger *log.Logger, filePath string) *DBMS {
	return &DBMS{logger, filePath}
}

func (d *DBMS) CreateTable() error {
	var table, column string

	fmt.Print("\ntable name : ")
	fmt.Scanf("%s", &table)

	fmt.Print("column name(ex. (a,b)) : ")
	fmt.Scanf("%s", &column)

	query := []byte("table : " + table + " | column : " + column + "\n")

	file, err := os.OpenFile(d.FilePath+"/table.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		d.Logger.Printf("failed to file open : %v", err)
		return err
	}
	defer file.Close()

	_, err = file.Write(query)
	if err != nil {
		d.Logger.Printf("failed to file wirte : %v", err)
		return err
	}

	return nil
}

func (d *DBMS) Insert() error {
	var table, column, value, value2 string

	inFile, err := os.OpenFile(d.FilePath+"/table.txt", os.O_RDONLY, 0644)
	if err != nil {
		d.Logger.Printf("failed to file open : %v", err)
		return err
	}

	var str string

	buff := make([]byte, 1024*1024)
	n, err := inFile.Read(buff)
	for err == nil {
		str = string(buff[0:n])
		n, err = inFile.Read(buff)
	}

	tempArr := strings.Split(str, "\n")
	strArr := strings.Split(tempArr[0], "|")
	tableStr := strings.Split(strArr[0], ":")
	columnStr := strings.Split(strArr[1], ":")

	fmt.Print("\ntable name : ")
	fmt.Scanf("%s", &table)

	if table != strings.TrimSpace(tableStr[1]) {
		fmt.Println("없는 테이블입니다.")
		return nil
	}

	fmt.Print("column name(ex. (a,b)) : ")
	fmt.Scanf("%s", &column)

	if column != strings.TrimSpace(columnStr[1]) {
		fmt.Println("없는 컬럼입니다.")
		return nil
	}

	fmt.Print("value(ex. a b) : ")
	fmt.Scanf("%s %s", &value, &value2)

	buf := []byte(value + " " + value2 + "\n")

	file, err := os.OpenFile(d.FilePath+"/table.txt", os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		d.Logger.Printf("failed to file open : %v", err)
		return err
	}
	defer file.Close()

	_, err = file.Write(buf)
	if err != nil {
		d.Logger.Printf("failed to file wirte : %v", err)
		return err
	}

	return nil
}

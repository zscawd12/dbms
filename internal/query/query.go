package query

import (
	"fmt"
	"io/ioutil"
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

	fmt.Print("column name(ex. a,b) : ")
	fmt.Scanf("%s", &column)

	query := []byte("table : " + table + " | column : " + column)

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

	byteArr, err := ioutil.ReadFile(d.FilePath + "/table.txt")
	if err != nil {
		d.Logger.Printf("failed to file read : %v", err)
		return err
	}

	str := string(byteArr)

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

	fmt.Print("column name(ex. a,b) : ")
	fmt.Scanf("%s", &column)

	if column != strings.TrimSpace(columnStr[1]) {
		fmt.Println("없는 컬럼입니다.")
		return nil
	}

	fmt.Print("value(ex. a b) : ")
	fmt.Scanf("%s %s", &value, &value2)

	buf := []byte("\n" + value + " " + value2)

	err = ioutil.WriteFile(d.FilePath+"/table.txt", buf, 0644)
	if err != nil {
		d.Logger.Printf("failed to file write : %v", err)
		return err
	}

	indexFile()

	return nil
}

func (d *DBMS) Select() error {
	var column, keyword string

	byteArr, err := ioutil.ReadFile(d.FilePath + "/table.txt")
	if err != nil {
		d.Logger.Printf("failed to file read : %v", err)
		return err
	}

	str := string(byteArr)

	tempArr := strings.Split(str, "\n")
	strArr := strings.Split(tempArr[0], "|")
	columnStr := strings.Split(strArr[1], ":")
	columnArr := strings.Split(columnStr[1], ",")

	fmt.Print("column name : ")
	fmt.Scanf("%s", &column)
	if column != strings.TrimSpace(columnArr[0]) && column != strings.TrimSpace(columnArr[1]) {
		fmt.Println("없는 컬럼입니다.")
		return nil
	}

	fmt.Print("keyword : ")
	fmt.Scanf("%s", &keyword)

	if keyword == "" {
		fmt.Println("키워드가 입력되지 않았습니다.")
		return nil
	}

	var ci int
	if column == strings.TrimSpace(columnArr[0]) {
		ci = 0
	} else {
		ci = 1
	}

	fmt.Printf("%v %v\n", columnArr[0], columnArr[1])
	for i := 1; i < len(tempArr); i++ {
		dataArr := strings.Split(tempArr[i], " ")
		if dataArr[ci] == keyword {
			fmt.Printf("%v\n", tempArr[i])
		}
	}

	return nil
}

func indexFile() {
	//	var m map[string]string
	//	m["test"] = "test"

	//	fmt.Printf("%v", m)
}

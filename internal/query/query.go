package query

import (
	//	"io/ioutil"
	"fmt"
	"log"
	"os"
)

type DBMS struct {
	Logger   *log.Logger
	FilePath string
}

func New(logger *log.Logger, filePath string) *DBMS {
	return &DBMS{logger, filePath}
}

func (d *DBMS) CreateTable() error {
	table := []byte("int,string\n")

	file, err := os.OpenFile(d.FilePath+"/table.txt", os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		d.Logger.Printf("failed tp file open : %v", err)
		return err
	}
	defer file.Close()

	_, err = file.Write(table)
	if err != nil {
		d.Logger.Printf("failed to file wirte : %v", err)
		return err
	}

	return nil
}

func (d *DBMS) Insert() error {
	var table, column, value string

	fmt.Print("\ntable name : ")
	fmt.Scanf("%s", &table)

	fmt.Print("column name(ex. (a,b)) : ")
	fmt.Scanf("%s", &column)

	fmt.Print("value(ex. (a,b)) : ")
	fmt.Scanf("%s", &value)

	fmt.Printf("%s, %s, %s\n", table, column, value)

	return nil
}

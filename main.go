// File: main.go
package main

import (
	"autocreate/common/config"
	"autocreate/common/logging"
	"fmt"
	"os"
	"time"
)

var f *os.File

func main() {

	var strDate string
	var filestrDate string
	// AppConfig := config.LoadConfig()
	logging.InfoLn(">> Start...")

	mydate := time.Now().AddDate(0, 0, -1)
	strDate = mydate.Format("2006-01-02")
	filestrDate = mydate.Format("20060102")
	fmt.Println("size:", len(os.Args))

	if len(os.Args) == 2 {
		strDate = os.Args[1] + "T11:45:26.371Z"
		fmt.Println(strDate)
		mydate, err := time.Parse(time.RFC3339, strDate)
		strDate = mydate.Format("2006-01-02")
		filestrDate = mydate.Format("20060102")
		if err != nil {
			fmt.Println("err: ", err)
			os.Exit(0)
		}

	} else if len(os.Args) > 2 {
		fmt.Println("invalid parameter.")
		os.Exit(0)
	}

	// GetTransactions(0, 50)

	//Create a folder/directory at a full qualified path
	err := os.Mkdir(filestrDate, 0755)
	if err != nil {
		// log.Fatal(err)
		fmt.Println(err)
	}
	Appcfg := config.LoadConfig()
	if Appcfg.Clients != nil {
		// DoGetClients(filestrDate)
		// DoGetClientsProd(filestrDate)
		Dodelete(filestrDate)
	} else {
		fmt.Println("Clients Null")
	}

	fmt.Println(">> Finish.")
	fmt.Println()
}

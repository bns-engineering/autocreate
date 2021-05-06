// File: main.go
package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bns-engineering/autocreate/common/logging"
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

	// DoGetClients(filestrDate)
	// DoGetClientsProd(filestrDate)
	DoCreate(filestrDate)
	// DoUpdate(filestrDate)
	// DoGetAccounts(filestrDate)

	fmt.Println(">> Wait Until Process Done.\nPress Ctrl+C for Stop.")

	// fmt.Println()
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-done
	log.Println("All server stopped!")
}

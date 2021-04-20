package main

import (
	"autocreate/common/config"
	"autocreate/common/helper"
	"autocreate/util"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func Dodelete(filestrDate string) {
	strFileName := fmt.Sprintf("%s.log", "./"+filestrDate+"/log")
	f, err := os.Create(strFileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Println(helper.GetStrDateTime(), ">> DeleteClient")
	strOut := fmt.Sprintf("Start")
	f.WriteString(strOut + fmt.Sprintln(""))
	Appcfg := config.LoadConfig()
	if Appcfg.Clients != nil {
		m := Appcfg.Clients
		for i, j := range m {
			clientID := j.ClientID
			fmt.Println(i, clientID)
			DodeleteItem(f, clientID)
		}
	}
	fmt.Print(helper.GetStrDateTime(), " GetClients done.\n")
	fmt.Print(helper.GetStrDateTime(), " File Output:", strFileName, "\n\n")
}

func DodeleteItem(pfile *os.File, clientID string) {

	strEndpoint := config.LoadConfig().Mambu.Endpoint
	requestBody, _ := json.Marshal("")
	timeout := time.Duration(10 * time.Second)
	baseUrl := strEndpoint + `/api/clients/` + clientID
	header := map[string]string{
		"Content-Type":  "application/json",
		"Connection":    "keep-alive",
		"Authorization": config.LoadConfig().Mambu.Authorization,
		"Accept":        "application/vnd.mambu.v2+json",
	}
	var hasil map[string]interface{}
	statusCode, responseBody, err := util.SendHTTP(http.MethodDelete, baseUrl, int(timeout), &hasil, header, bytes.NewBuffer(requestBody))

	fmt.Println("Statuscode:", statusCode)
	if statusCode >= 400 {
		strerror := hasil["errors"]
		strerror = strerror.([]interface{})[0]
		strerrordata := strerror.(map[string]interface{})
		fmt.Println("Statuscode:", strerrordata["errorCode"])
		// fmt.Println("errorSource:", strerrordata["errorSource"])
		// fmt.Println("errorReason:", strerrordata["errorReason"])
		fmt.Println(clientID + " delete FAILED.")
	}

	if err != nil {
		// fmt.Println(strlog)
	}

	if statusCode >= 200 && statusCode < 300 {
		if responseBody == "" {
			fmt.Println(responseBody)
		}
		fmt.Println(clientID + " delete SUCCESS.")

	}
	return
}

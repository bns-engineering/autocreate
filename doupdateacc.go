package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/bns-engineering/autocreate/common/config"
	"github.com/bns-engineering/autocreate/common/helper"
	"github.com/bns-engineering/autocreate/util"
)

func DoGetAccounts(filestrDate string) {
	strFileName := fmt.Sprintf("%s.sql", "./"+filestrDate+"/depositAccount")
	f, err := os.Create(strFileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Println(helper.GetStrDateTime(), ">> GetAccountes")
	// DoUpdateDepositAccount(f, "11260000101", "encodedKey")
	GetAccounts(f, 0, 1000)
	fmt.Print(helper.GetStrDateTime(), " GetAccountes done.\n")
	fmt.Print(helper.GetStrDateTime(), " File Output:", strFileName, "\n\n")
}

func GetAccounts(pfile *os.File, noffset int, nlimit int) {

	requestBody, _ := json.Marshal("")
	timeout := time.Duration(10 * time.Second)
	baseUrl := config.LoadConfig().Mambu.Endpoint + `/api/deposits?detailsLevel=FULL&offset=` + fmt.Sprint(noffset) + `&limit=` + fmt.Sprint(nlimit)
	header := map[string]string{
		"Content-Type":  "application/json",
		"Connection":    "keep-alive",
		"Authorization": config.LoadConfig().Mambu.Authorization,
		"Accept":        "application/vnd.mambu.v2+json",
	}
	var hasil map[string]interface{}
	statusCode, responseBody, err := util.SendHTTP(http.MethodGet, baseUrl, int(timeout), &hasil, header, bytes.NewBuffer(requestBody))

	fmt.Println("Statuscode:", statusCode)
	if statusCode >= 400 {
		strerror := hasil["errors"]
		strerror = strerror.([]interface{})[0]
		strerrordata := strerror.(map[string]interface{})
		fmt.Println("Statuscode:", strerrordata["errorCode"])
		fmt.Println("errorSource:", strerrordata["errorSource"])
		fmt.Println("errorReason:", strerrordata["errorReason"])
	}

	if err != nil {
		// fmt.Println(strlog)
	}

	if statusCode == 200 {
		b := []byte(responseBody)
		var f interface{}
		err := json.Unmarshal(b, &f)
		if err == nil {
			m := f.([]interface{})
			for i, j := range m {
				datarow := j.(map[string]interface{})
				accountType := fmt.Sprintf("%v", datarow["accountType"])
				name := fmt.Sprintf("%v", datarow["name"])
				creationDate := fmt.Sprintf("%v", datarow["creationDate"])
				if accountType == "FIXED_DEPOSIT" {
					if datarow["_otherInformation"] != nil {
						_otherInformation := datarow["_otherInformation"]
						data_otherInformation := _otherInformation.(map[string]interface{})
						if data_otherInformation["nisbahAkhir"] != nil {
							// nisbah := fmt.Sprintf("%v", data_otherInformation["nisbahAkhir"])

							if name == "FS dummy FIXED" {
								id := fmt.Sprintf("%v", datarow["id"])
								go func() {
									// DoUpdateAccount(pfile, id)
									SetMaturityDate(pfile, id, creationDate[0:5]+"05"+creationDate[7:10])
									fmt.Println(i)
								}()
							}
						}
					}

				}

			}
			if len(m) == nlimit {
				// go func() {
				fmt.Println(noffset + nlimit)
				GetAccounts(pfile, noffset+nlimit, nlimit)
				// }()
			}

		}

	}
	return
}

func DoUpdateAccount(pfile *os.File, accountID string) {

	strEndpoint := config.LoadConfig().Mambu.Endpoint
	// requestBody, _ := json.Marshal("")
	var strrequestBody string
	strrequestBody = `[
		{
			"op": "REPLACE",
			"path": "_otherInformation",
			"value": {
				"nisbahAkhir": "10.6",
				"nisbahCounter": "10.6",
				"purpose": "Tabungan",
				"sourceOfFund": "Gaji",
				"tncVersion": "1.0",
				"tenor":"1",
				"aroNonAro":"."
			}
		}
	]`
	requestBody := []byte(strrequestBody)
	timeout := time.Duration(10 * time.Second)
	baseUrl := strEndpoint + `/api/deposits/` + accountID
	header := map[string]string{
		"Content-Type":  "application/json",
		"Connection":    "keep-alive",
		"Authorization": config.LoadConfig().Mambu.Authorization,
		"Accept":        "application/vnd.mambu.v2+json",
	}
	var hasil map[string]interface{}
	statusCode, responseBody, err := util.SendHTTP(http.MethodPatch, baseUrl, int(timeout), &hasil, header, bytes.NewBuffer(requestBody))

	fmt.Println("Statuscode:", statusCode)
	if statusCode >= 400 {
		if hasil["errors"] != nil {
			strerror := hasil["errors"]
			strerror = strerror.([]interface{})[0]
			strerrordata := strerror.(map[string]interface{})
			strLog := fmt.Sprint("errorCode:", strerrordata["errorCode"]) + ";"
			fmt.Println("errorCode:", strerrordata["errorCode"])
			if strerrordata["errorSource"] != nil {
				strLog = strLog + fmt.Sprint("errorSource:", strerrordata["errorSource"]) + ";"
			}
			strLog = strLog + fmt.Sprint("errorReason:", strerrordata["errorReason"]) + ";"

			strOut := fmt.Sprintf(accountID + " " + strLog)
			fmt.Println(strOut)
			f.WriteString(strOut + fmt.Sprintln(""))
		}

	}

	if err != nil {
		fmt.Println(responseBody)
	}

	if statusCode >= 200 && statusCode < 300 {
		if responseBody == "" {
			fmt.Println(responseBody)
		}
		strOut := fmt.Sprintf(accountID + " Update DepositAccount SUCCESS.")
		f.WriteString(strOut + fmt.Sprintln(""))
		fmt.Println(strOut)

	}
	return
}

func SetMaturityDate(pfile *os.File, accountID string, MaturityDate string) {

	strEndpoint := config.LoadConfig().Mambu.Endpoint
	// requestBody, _ := json.Marshal("")
	var strrequestBody string
	strrequestBody = `{
			"maturityDate": "{{MaturityDate}}",
			"notes": "Update"
		}`
	strrequestBody = strings.ReplaceAll(strrequestBody, "{{MaturityDate}}", MaturityDate)
	requestBody := []byte(strrequestBody)
	timeout := time.Duration(10 * time.Second)
	baseUrl := strEndpoint + `/api/deposits/` + accountID + `:startMaturity`
	header := map[string]string{
		"Content-Type":  "application/json",
		"Connection":    "keep-alive",
		"Authorization": config.LoadConfig().Mambu.Authorization,
		"Accept":        "application/vnd.mambu.v2+json",
	}
	var hasil map[string]interface{}
	statusCode, responseBody, err := util.SendHTTP(http.MethodPost, baseUrl, int(timeout), &hasil, header, bytes.NewBuffer(requestBody))

	fmt.Println("Statuscode:", statusCode)
	if statusCode >= 400 {
		if hasil["errors"] != nil {
			strerror := hasil["errors"]
			strerror = strerror.([]interface{})[0]
			strerrordata := strerror.(map[string]interface{})
			strLog := fmt.Sprint("errorCode:", strerrordata["errorCode"]) + ";"
			fmt.Println("errorCode:", strerrordata["errorCode"])
			if strerrordata["errorSource"] != nil {
				strLog = strLog + fmt.Sprint("errorSource:", strerrordata["errorSource"]) + ";"
			}
			strLog = strLog + fmt.Sprint("errorReason:", strerrordata["errorReason"]) + ";"

			strOut := fmt.Sprintf(accountID + " " + strLog)
			fmt.Println(strOut)
			pfile.WriteString(helper.GetStrtimestamp() + " " + strOut + fmt.Sprintln(""))

		}

	}

	if err != nil {
		fmt.Println(responseBody)
	}

	if statusCode >= 200 && statusCode < 300 {
		if responseBody == "" {
			fmt.Println(responseBody)
		}
		strOut := fmt.Sprintf(accountID + " Update Maturity Date SUCCESS.")
		pfile.WriteString(helper.GetStrtimestamp() + " " + strOut + fmt.Sprintln(""))
		fmt.Println(strOut)

	}
	return
}

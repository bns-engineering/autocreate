package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/bns-engineering/autocreate/common/config"
	"github.com/bns-engineering/autocreate/common/helper"
	"github.com/bns-engineering/autocreate/util"
	//2021-04-23 08:38:16 >> CreateClient 100
	//2021-04-23 08:41:50 CreateClient done. 3 menit
)

func DoCreate(filestrDate string) {
	var wg sync.WaitGroup
	Appcfg := config.LoadConfig()
	wg.Add(Appcfg.Until - Appcfg.Start + 1)

	strFileName := fmt.Sprintf("%s.log", "./"+filestrDate+"/log")
	f, err := os.Create(strFileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Println(helper.GetStrDateTime(), ">> CreateClient")
	strOut := fmt.Sprintf(helper.GetStrtimestamp() + " Start")
	f.WriteString(strOut + fmt.Sprintln(""))

	if Appcfg.Start > -1 {
		for loop := Appcfg.Start; loop <= Appcfg.Until; loop++ {
			index := loop
			tmp := helper.GetStrtimestamp()
			clientID := fmt.Sprintf("1" + tmp[4:8] + fmt.Sprintf("%07d", index))
			DoCreateItem(f, clientID)
			// DoCreateItemDvp(f, clientID)

		}

	}
	strOut = fmt.Sprintf(helper.GetStrtimestamp() + " Finish")
	f.WriteString(strOut + fmt.Sprintln(""))
	fmt.Print(helper.GetStrDateTime(), " CreateClient done.\n")
	fmt.Print(helper.GetStrDateTime(), " File Output:", strFileName, "\n\n")

}

// func DoUpdate(filestrDate string) {
// 	var wg sync.WaitGroup
// 	Appcfg := config.LoadConfig()
// 	wg.Add(Appcfg.Until - Appcfg.Start + 1)

// 	strFileName := fmt.Sprintf("%s.log", "./"+filestrDate+"/log")
// 	f, err := os.Create(strFileName)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer f.Close()

// 	fmt.Println(helper.GetStrDateTime(), ">> UpdateAccount")
// 	strOut := fmt.Sprintf("Start")
// 	f.WriteString(strOut + fmt.Sprintln(""))

// 	if Appcfg.Start > 0 {
// 		for loop := 1; loop <= 2144; loop++ {
// 			index := loop
// 			// tmp := helper.GetStrtimestamp()
// 			accountID := fmt.Sprintf("1" + "123" + fmt.Sprintf("%07d", index))
// 			DoUpdateDepositAccount(f, accountID)

// 		}

// 	}

// 	fmt.Print(helper.GetStrDateTime(), " UpdateAccount done.\n")
// 	fmt.Print(helper.GetStrDateTime(), " File Output:", strFileName, "\n\n")

// }

func DoCreateItem(pfile *os.File, clientID string) {

	strEndpoint := config.LoadConfig().Mambu.Endpoint
	// requestBody, _ := json.Marshal("")
	strrequestBody := `{
        "id": "{{id}}",
        "firstName": "FSCIF {{id}}",
        "lastName": "Dummy",
        "homePhone": "0212329389283",
        "mobilePhone": "0878121212121",
        "mobilePhone2": "0878121212122",
        "emailAddress": "fery.setianto@gmail.com",
        "preferredLanguage": "ENGLISH",
        "birthDate": "2000-06-01",
        "gender": "MALE",
        "notes": "",
        "loanCycle": 0,
        "groupLoanCycle": 0,
        "groupKeys": [],
        "addresses": [],
        "idDocuments": [],
        "assignedBranchKey": "8a8e8fab786e635c0178863b7911431e",
        "_personalIdData": {
            "personalIdDateOfIssue": "2018-02-01",
            "urlImageSelfie": "URL Selfi",
            "personalReferralCode": "ReffCode",
            "personalIdExpireDate": "2023-03-31",
            "personalIdNumber": "237268376218762381",
            "personalNpwpNumber": "8237498279482739487",
            "personalIdType": "KTP",
            "personalBranch": "001",
            "urlImageKtp": "URL_KTP"
        },
        "_mailingAddress": {
            "mailingAddressCity": "Bekasi",
            "mailingRecipientPhoneNumber": "0878121212121",
            "mailingAddress": "Alamat Surat Menyurat",
            "mailingAddressVillage": "Duren Jaya",
            "mailingAddressProvince": "Jawa Barat",
            "mailingAddressRtRw": "07/16",
            "mailingAddressPostalCode": "17111",
            "mailingAddressSubDistrict": "Bekasi Timur"
        },
        "_occupationInfo": {
            "personalOccupation": "Karyawan BUMN",
            "companyName": "PT.  Sumber Jaya",
            "personalMonthlyIncome": "Kurang dari Rp5JT",
            "personalTypeOfWork": "Kesehatan"
        },
        "_customerData": {
            "birthPlace": "Jakarta",
            "personalReligion": "Islam",
            "personalMothersName": "Ibuku",
            "personalMaritalStatus": "Menikah",
            "personalScoringCust": "9",
            "personalCitizenship": "WNI"
        },
        "_ktpAddress": {
            "ktpRtRw": "07/016",
            "ktpSubDistrict": "Bekasi Timur",
            "ktpVillage": "Duren Jaya",
            "ktpAddress": "Alamat KTP",
            "ktpCity": "Bekasi ",
            "ktpProvince": "Jawa Barat",
            "ktpPostalCode": "17111"
        },
        "_companyAddress": {
            "companyAddressCity": "Kabupaten (Office)",
            "companyAddressRtRw": "RT/RW (Office)",
            "companyAddressVillage": "Kelurahan (Office)",
            "companyAddress": "Alamat Kantor",
            "companyAddressSubDistrict": "Kecamatan (Office)",
            "companyAddressPostalCode": "23231",
            "companyAddressProvince": "Provinsi (Office)",
            "companyAddressPhoneNumber": "021873872873"
        }
        
	}`
	strrequestBody = strings.ReplaceAll(strrequestBody, "{{id}}", clientID)
	requestBody := []byte(strrequestBody)
	timeout := time.Duration(10 * time.Second)
	baseUrl := strEndpoint + `/api/clients/`
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
		strerror := hasil["errors"]
		strerror = strerror.([]interface{})[0]
		strerrordata := strerror.(map[string]interface{})
		strLog := fmt.Sprint("errorCode:", strerrordata["errorCode"]) + ";"
		fmt.Println("errorCode:", strerrordata["errorCode"])
		if strerrordata["errorSource"] != nil {
			strLog = strLog + fmt.Sprint("errorSource:", strerrordata["errorSource"]) + ";"
		}
		strLog = strLog + fmt.Sprint("errorReason:", strerrordata["errorReason"]) + ";"

		strOut := fmt.Sprintf(clientID + " " + strLog)
		fmt.Println(strOut)
		pfile.WriteString(helper.GetStrtimestamp() + " " + strOut + fmt.Sprintln(""))

		// productEndcodeKey := "8a8e862a78c9cd600178ca689c2800e3"
		// strEncodedKey := "8a8e86387918509501793aa619c02df0"
		// accountID := "10" + clientID[3:12]
		// DoCreateDepositAccount(pfile, clientID, strEncodedKey, "SAVING", productEndcodeKey, accountID)
		// strFixedDepositEndcodeKey := "8a8e865978ca233a0178ca6adb4a0019"
		// accountID = "11" + clientID[3:12]
		// DoCreateDepositAccount(pfile, clientID, strEncodedKey, "FIXED", strFixedDepositEndcodeKey, accountID)

	}

	if err != nil {
		// fmt.Println(strlog)
	}

	if statusCode >= 200 && statusCode < 300 {
		EncodedKey := hasil["encodedKey"]
		strEncodedKey := fmt.Sprintf("%v", EncodedKey)
		if responseBody == "" {
			fmt.Println(responseBody)
		}
		strOut := fmt.Sprintf(clientID + " Client Created SUCCESS.")
		pfile.WriteString(helper.GetStrtimestamp() + " " + strOut + fmt.Sprintln(""))

		// var wg sync.WaitGroup
		// wg.Add(1)
		// go func() {
		strSavingEndcodeKey := "8a8e862a78c9cd600178ca689c2800e3"
		accountID := "10" + clientID[3:12]
		go DoCreateDepositAccount(pfile, clientID, strEncodedKey, "SAVING", strSavingEndcodeKey, accountID)
		// defer wg.Done()
		// }()

		// go func() {
		strFixedDepositEndcodeKey := "8a8e865978ca233a0178ca6adb4a0019"
		accountID = "11" + clientID[3:12]
		go DoCreateDepositAccount(pfile, clientID, strEncodedKey, "FIXED", strFixedDepositEndcodeKey, accountID)
		// defer wg.Done()
		// }()
		// wg.Wait()

	}
	return
}

// func DoCreateItemDvp(pfile *os.File, clientID string) {

// 	strEndpoint := config.LoadConfig().Mambu.Endpoint
// 	// requestBody, _ := json.Marshal("")
// 	strrequestBody := `{
//         "id": "{{id}}",
// 		"firstName": "Dummy CIF",
// 		"lastName": "Dummy",
// 		"homePhone": "0212329389283",
// 		"mobilePhone": "{{id}}",
// 		"emailAddress": "fery.setianto@gmail.com",
// 		"preferredLanguage": "ENGLISH",
// 		"birthDate": "2000-06-01",
// 		"gender": "MALE",
// 		"notes": "",
// 		"loanCycle": 0,
// 		"groupLoanCycle": 0,
// 		"groupKeys": [],
// 		"addresses": [],
// 		"idDocuments": []

// 		}`
// 	strrequestBody = strings.ReplaceAll(strrequestBody, "{{id}}", clientID)
// 	requestBody := []byte(strrequestBody)
// 	timeout := time.Duration(10 * time.Second)
// 	baseUrl := strEndpoint + `/api/clients/`
// 	header := map[string]string{
// 		"Content-Type":  "application/json",
// 		"Connection":    "keep-alive",
// 		"Authorization": config.LoadConfig().Mambu.Authorization,
// 		"Accept":        "application/vnd.mambu.v2+json",
// 	}
// 	var hasil map[string]interface{}
// 	statusCode, responseBody, err := util.SendHTTP(http.MethodPost, baseUrl, int(timeout), &hasil, header, bytes.NewBuffer(requestBody))

// 	fmt.Println("Statuscode:", statusCode)
// 	if statusCode >= 400 {
// 		strerror := hasil["errors"]
// 		strerror = strerror.([]interface{})[0]
// 		strerrordata := strerror.(map[string]interface{})
// 		strLog := fmt.Sprint("errorCode:", strerrordata["errorCode"]) + ";"
// 		fmt.Println("errorCode:", strerrordata["errorCode"])
// 		if strerrordata["errorSource"] != nil {
// 			strLog = strLog + fmt.Sprint("errorSource:", strerrordata["errorSource"]) + ";"
// 		}
// 		strLog = strLog + fmt.Sprint("errorReason:", strerrordata["errorReason"]) + ";"

// 		strOut := fmt.Sprintf(clientID + " " + strLog)
// 		fmt.Println(strOut)
// 		f.WriteString(strOut + fmt.Sprintln(""))
// 	}

// 	if err != nil {
// 		// fmt.Println(strlog)
// 	}

// 	if statusCode >= 200 && statusCode < 300 {
// 		EncodedKey := hasil["encodedKey"]
// 		strEncodedKey := fmt.Sprintf("%v", EncodedKey)
// 		if responseBody == "" {
// 			fmt.Println(responseBody)
// 		}
// 		strOut := fmt.Sprintf(clientID + " Create SUCCESS.")
// 		f.WriteString(strOut + fmt.Sprintln(""))

// 		strSavingEndcodeKey := "8a8e87b0791850e101792ff142b71397"

// 		accountID := "10" + clientID[3:12]
// 		go DoCreateDepositAccountDvp(pfile, clientID, strEncodedKey, "SAVING", strSavingEndcodeKey, accountID)

// 		strFixedDepositEndcodeKey := "8a8e87b0791850e101792db95954132e"
// 		accountID = "11" + clientID[3:12]
// 		go DoCreateDepositAccountDvp(pfile, clientID, strEncodedKey, "FIXED", strFixedDepositEndcodeKey, accountID)

// 	}
// 	return
// }
func DoCreateDepositAccount(pfile *os.File, clientID, clientEncodedKey string, productType string, productEndcodeKey string, accountID string) {
	var strEncodedKey string
	Appcfg := config.LoadConfig()
	MaturityDate := Appcfg.MaturityDate
	strEndpoint := config.LoadConfig().Mambu.Endpoint
	// requestBody, _ := json.Marshal("")
	var strrequestBody string
	if productType == "FIXED" {
		strrequestBody = `
        {
			"id": "{{id}}",
			"accountHolderKey": "{{accountHolderKey}}",	
			"accountHolderType": "CLIENT",
			"name": "Dummy FIXED",
			"productTypeKey": "{{productTypeKey}}",
			"accountState": "APPROVED",
			"_otherInformation": {
				"nisbahAkhir": "0",
				"nisbahCounter": "0",
				"nisbahZakat": "0",
				"nisbahPajak": "0",
				"purpose": "Tabungan",
				"sourceOfFund": "Gaji",
				"tncVersion": "1.0",
				"tenor":"1",
				"aroNonAro":"."
			}
		}`
	} else {
		strrequestBody = `
        {
			"id": "{{id}}",
			"accountHolderKey": "{{accountHolderKey}}",	
			"accountHolderType": "CLIENT",
			"name": "Dummy SAVING",
			"productTypeKey": "{{productTypeKey}}",
			"accountState": "APPROVED",
			"_otherInformation": {
				"nisbahAkhir": "0",
				"nisbahCounter": "0",
				"nisbahZakat": "0",
				"nisbahPajak": "0",
				"purpose": "Tabungan",
				"sourceOfFund": "Gaji",
				"tncVersion": "1.0"
			}
		}`
	}

	strrequestBody = strings.ReplaceAll(strrequestBody, "{{id}}", accountID)
	strrequestBody = strings.ReplaceAll(strrequestBody, "{{accountHolderKey}}", clientEncodedKey)
	strrequestBody = strings.ReplaceAll(strrequestBody, "{{productTypeKey}}", productEndcodeKey)
	requestBody := []byte(strrequestBody)
	timeout := time.Duration(10 * time.Second)
	baseUrl := strEndpoint + `/api/deposits`
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

			strOut := fmt.Sprintf(clientID + " " + strLog)
			fmt.Println(strOut)
			pfile.WriteString(helper.GetStrtimestamp() + " " + strOut + fmt.Sprintln(""))

		}

	}

	if err != nil {
		// fmt.Println(strlog)
	}

	if statusCode >= 200 && statusCode < 300 {
		EncodedKey := hasil["encodedKey"]
		strEncodedKey = fmt.Sprintf("%v", EncodedKey)
		if responseBody == "" {
			fmt.Println(responseBody)
		}
		strOut := fmt.Sprintf(clientID + ";" + accountID + " Create DepositAccount SUCCESS.")
		pfile.WriteString(helper.GetStrtimestamp() + " " + strOut + fmt.Sprintln(""))
		if EncodedKey != "" {
			DoDepositTransaction(pfile, strEncodedKey, "CASH", "10000000")
			if productType == "FIXED" {
				SetMaturityDate(pfile, accountID, MaturityDate)
			}
		}

	}
	return
}

// func DoCreateDepositAccountDvp(pfile *os.File, clientID, clientEncodedKey string, productType string, productEndcodeKey string, accountID string) {
// 	var strEncodedKey string
// 	Appcfg := config.LoadConfig()
// 	MaturityDate := Appcfg.ValueDate
// 	strEndpoint := config.LoadConfig().Mambu.Endpoint
// 	// requestBody, _ := json.Marshal("")
// 	var strrequestBody string
// 	if productType == "FIXED" {
// 		strrequestBody = `
//         {
// 			"id": "{{id}}",
// 			"accountHolderKey": "{{accountHolderKey}}",
// 			"accountHolderType": "CLIENT",
// 			"name": "FS dummy FIXED",
// 			"productTypeKey": "{{productTypeKey}}",
// 			"accountState": "APPROVED",
// 			"interestSettings": {
// 				"interestRateSettings": {
// 					"encodedKey": "8a8e8638791850950179307b533d1596",
// 					"interestRate": 3,
// 					"interestChargeFrequency": "ANNUALIZED",
// 					"interestChargeFrequencyCount": 1,
// 					"interestRateTerms": "FIXED"
// 				},
// 				"interestPaymentSettings": {
// 					"interestPaymentPoint": "ON_ACCOUNT_MATURITY",
// 					"interestPaymentDates": []
// 				}
// 			}
// 		}`
// 	} else {
// 		strrequestBody = `
//         {
// 			"id": "{{id}}",
// 			"accountHolderKey": "{{accountHolderKey}}",
// 			"accountHolderType": "CLIENT",
// 			"name": "FS dummy SAVING",
// 			"productTypeKey": "{{productTypeKey}}",
// 			"accountState": "APPROVED",
// 			"interestSettings": {
// 			"interestRateSettings": {
// 					"encodedKey": "8a8e860477b2bbde0177b408a64c5576",
// 					"interestRate": 1,
// 					"interestChargeFrequency": "ANNUALIZED",
// 					"interestChargeFrequencyCount": 1,
// 					"interestRateTerms": "FIXED"
// 				},
// 				"interestPaymentSettings": {
// 					"interestPaymentPoint": "FIRST_DAY_OF_MONTH",
// 					"interestPaymentDates": []
// 				}
// 			}
// 		}`
// 	}

// 	strrequestBody = strings.ReplaceAll(strrequestBody, "{{id}}", accountID)
// 	strrequestBody = strings.ReplaceAll(strrequestBody, "{{accountHolderKey}}", clientEncodedKey)
// 	strrequestBody = strings.ReplaceAll(strrequestBody, "{{productTypeKey}}", productEndcodeKey)
// 	requestBody := []byte(strrequestBody)
// 	timeout := time.Duration(10 * time.Second)
// 	baseUrl := strEndpoint + `/api/deposits`
// 	header := map[string]string{
// 		"Content-Type":  "application/json",
// 		"Connection":    "keep-alive",
// 		"Authorization": config.LoadConfig().Mambu.Authorization,
// 		"Accept":        "application/vnd.mambu.v2+json",
// 	}
// 	var hasil map[string]interface{}
// 	statusCode, responseBody, err := util.SendHTTP(http.MethodPost, baseUrl, int(timeout), &hasil, header, bytes.NewBuffer(requestBody))

// 	fmt.Println("Statuscode:", statusCode)
// 	if statusCode >= 400 {
// 		if hasil["errors"] != nil {
// 			strerror := hasil["errors"]
// 			strerror = strerror.([]interface{})[0]
// 			strerrordata := strerror.(map[string]interface{})
// 			strLog := fmt.Sprint("errorCode:", strerrordata["errorCode"]) + ";"
// 			fmt.Println("errorCode:", strerrordata["errorCode"])
// 			if strerrordata["errorSource"] != nil {
// 				strLog = strLog + fmt.Sprint("errorSource:", strerrordata["errorSource"]) + ";"
// 			}
// 			strLog = strLog + fmt.Sprint("errorReason:", strerrordata["errorReason"]) + ";"

// 			strOut := fmt.Sprintf(clientID + " " + strLog)
// 			fmt.Println(strOut)
// 			f.WriteString(strOut + fmt.Sprintln(""))
// 		}

// 	}

// 	if err != nil {
// 		// fmt.Println(strlog)
// 	}

// 	if statusCode >= 200 && statusCode < 300 {
// 		EncodedKey := hasil["encodedKey"]
// 		strEncodedKey = fmt.Sprintf("%v", EncodedKey)
// 		if responseBody == "" {
// 			fmt.Println(responseBody)
// 		}
// 		strOut := fmt.Sprintf(clientID + ";" + accountID + " Create DepositAccount SUCCESS.")
// 		f.WriteString(strOut + fmt.Sprintln(""))
// 		fmt.Println(strOut)
// 		if EncodedKey != "" {
// 			DoDepositTransaction(pfile, strEncodedKey, "cash", "10000000")
// 			if productType == "FIXED" {

// 				SetMaturityDate(pfile, accountID, MaturityDate)
// 			}
// 		}

// 	}
// 	return
// }

func DoDepositTransaction(pfile *os.File, accountEncodedKey string, transactionChannelId string, amount string) string {
	var strEncodedKey string
	strEndpoint := config.LoadConfig().Mambu.Endpoint
	// requestBody, _ := json.Marshal("")
	var strrequestBody string
	strrequestBody = `
	{
		"bookingDate":"{{valueDate}}T10:00:00+07:00",
		"valueDate":"{{valueDate}}T10:00:00+07:00",
		"amount": {{amount}},
		"transactionDetails": {
		"transactionChannelId": "{{transactionChannelId}}"
		}
	}`

	Appcfg := config.LoadConfig()
	valueDate := Appcfg.ValueDate
	strrequestBody = strings.ReplaceAll(strrequestBody, "{{valueDate}}", valueDate)
	strrequestBody = strings.ReplaceAll(strrequestBody, "{{amount}}", amount)
	strrequestBody = strings.ReplaceAll(strrequestBody, "{{transactionChannelId}}", transactionChannelId)
	requestBody := []byte(strrequestBody)
	timeout := time.Duration(10 * time.Second)
	baseUrl := strEndpoint + `/api/deposits/{{accountEncodedKey}}/deposit-transactions`
	baseUrl = strings.ReplaceAll(baseUrl, "{{accountEncodedKey}}", accountEncodedKey)
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

			strOut := fmt.Sprintf("Deposit Transaction" + " " + strLog)
			fmt.Println(strOut)
			pfile.WriteString(helper.GetStrtimestamp() + " " + strOut + fmt.Sprintln(""))

		}

	}

	if err != nil {
		// fmt.Println(strlog)
	}

	if statusCode >= 200 && statusCode < 300 {
		EncodedKey := hasil["encodedKey"]
		strEncodedKey = fmt.Sprintf("%v", EncodedKey)
		if responseBody == "" {
			fmt.Println(responseBody)
		}
		strOut := fmt.Sprintf("Deposit Transaction" + " SUCCESS.")
		pfile.WriteString(helper.GetStrtimestamp() + " " + strOut + fmt.Sprintln(""))
		fmt.Println(strOut)

	}
	return strEncodedKey
}

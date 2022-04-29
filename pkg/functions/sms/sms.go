package sms

import (
	"fmt"
	"io/ioutil"
	"log"
	"project/skillbox/Diplom/pkg/data"
	"strings"
)

type SMSData struct {
	Country      string
	Bandwidth    string
	ResponseTime string
	Provider     string
}

func Sms() {
	storage := parseSmsFile()
	printStore(storage)
}

func parseSmsFile() []SMSData {
	//Read file
	content, err := ioutil.ReadFile("./simulator/sms.data")
	if err != nil {
		log.Fatal(err)
	}

	//Parse and check file
	var storage []SMSData

	parseStr := strings.Split(string(content), "\n")

	for _, line := range parseStr {
		if strings.Count(line, ";") == 3 {
			str := strings.Split(line, ";")
			if _, ok := data.Providers[str[3]]; ok {
				if _, okk := data.Countries[str[0]]; okk {
					smsData := SMSData{Country: str[0], Bandwidth: str[1], ResponseTime: str[2], Provider: str[3]}
					storage = append(storage, smsData)
				}
			}
		}
	}
	return storage
}

func printStore(storage []SMSData) {
	for _, str := range storage {
		fmt.Printf("%v\n", str)
	}
}

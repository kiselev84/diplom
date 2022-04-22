package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type SMSData struct {
	Country      string
	Bandwidth    string
	ResponseTime string
	Provider     string
}
type Alpha2 struct {
	Country     string
	Cod    		string
}

func main() {
	storage := parseSmsFile()
	printStore(storage)
}

func parseAlpha2() map[int]*Alpha2 {
	//Read file
	content, err := ioutil.ReadFile("Alpha-2.csv")
	if err != nil {
		log.Fatal(err)
	}

	//Parse file
	parseStr := strings.Split(string(content), "\n")
	storage := make(map[int]*Alpha2)
	index := 0
	for _, data := range parseStr {
		str := strings.Split(data, ";")
		if len(str) == 2 {
			storage[index] = &Alpha2{Country: str[0], Cod: str[1]}
			index++
		}
	}
	return storage
}

func parseSmsFile() map[int]*SMSData {
	//Read file
	content, err := ioutil.ReadFile("./clone/sms.data")
	if err != nil {
		log.Fatal(err)
	}

	//Parse file
	checkAlpha2 := parseAlpha2()
	parseStr := strings.Split(string(content), "\n")
	storage := make(map[int]*SMSData)
	index := 0
	
	for _, data := range parseStr {
		str := strings.Split(data, ";")
		if len(str) == 4 {
			if str[3] == "Rond"|| str[3] == "Topolo" || str[3] == "Kildy" {
				for i, _ := range checkAlpha2 {
					if str[0] == checkAlpha2[i].Cod {
						storage[index] = &SMSData{Country: str[0], Bandwidth: str[1], ResponseTime: str[2], Provider: str[3]}
						index++
					}
				}
			}
		}
	}
	return storage

}

func printStore(storage map[int]*SMSData) {
	//Print file
	for i, _ := range storage {
		fmt.Println(storage[i])
	}
}

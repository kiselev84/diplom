package mms

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"project/skillbox/Diplom/pkg/controller"
	"project/skillbox/Diplom/pkg/data"

	"sort"
	"strings"
)

type MMSData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}

func Result() [][]MMSData {
	result := make([][]MMSData, 2)
	result[0] = mmsSortByCountry(mmsContent())
	result[1] = mmsSortByProvider(mmsContent())
	return result
}

func mmsContent() []MMSData {
	resp := controller.GetResponse("http://127.0.0.1:8383/mms")
	mmsData := make([]MMSData, 0)
	defer resp.Body.Close()
	if resp.StatusCode == 200 {

		jsonStream, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		dec := json.NewDecoder(strings.NewReader(string(jsonStream)))

		_, err = dec.Token()
		if err != nil {
			log.Fatal(err)
		}

		for dec.More() {
			var m MMSData
			err := dec.Decode(&m)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(m)
			if _, ok := data.Countries[m.Country]; ok {
				if _, ok := data.Providers[m.Provider]; ok {
					mmsData = append(mmsData, m)
				}
			}
		}
	}
	return mmsData
}

func mmsSortByCountry(b []MMSData) []MMSData {
	sort.Slice(b, func(i, j int) bool { return b[i].Country < b[j].Country })
	return b
}

func mmsSortByProvider(b []MMSData) []MMSData {
	sort.Slice(b, func(i, j int) bool { return b[i].Provider < b[j].Provider })
	return b
}

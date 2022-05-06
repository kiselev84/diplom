package support

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"project/skillbox/Diplom/pkg/controller"
	"strings"
)

type SupportData struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}

func Result() []int {
	a := supportContent()
	b := make([]int, 2)
	var load int
	for _, v := range a {
		load += v.ActiveTickets
	}
	switch {
	case load > 16:
		b[0] = 3
	case load >= 9 && load <= 16:
		b[0] = 2
	default:
		b[0] = 1
	}
	b[1] = int(float64(load) * 60 / 18)
	return b
}

func supportContent() []SupportData {
	resp := controller.GetResponse("http://127.0.0.1:8383/support")
	dataSlice := make([]SupportData, 0)
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
			var tmp SupportData
			err := dec.Decode(&tmp)
			if err != nil {
				log.Fatal(err)
			}
			dataSlice = append(dataSlice, tmp)
		}

	}
	return dataSlice
}

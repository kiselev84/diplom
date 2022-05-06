package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"project/skillbox/Diplom/pkg/result"
)

func HandleConnection(res *result.ResultT) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		json, err := json.MarshalIndent(res, "", "")
		if err != nil {
			log.Fatal(err)
		}
		w.Write(json)
	}
}

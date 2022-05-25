package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"project/skillbox/Diplom/pkg/result"
	"project/skillbox/Diplom/simulator"
	"time"
)

func main() {
	r := mux.NewRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8383"
	}
	ticker := time.NewTicker(30 * time.Second)
	res := result.GetRes()
	go func() {
		for range ticker.C {
			simulator.Shuffle()
			res = result.GetRes()
		}
	}()
	r.HandleFunc("/api", handleConnection(&res))
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("../pkg/static")))
	s := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}
	log.Fatal(s.ListenAndServe())

}

func handleConnection(res *result.ResultT) func(w http.ResponseWriter, r *http.Request) {
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

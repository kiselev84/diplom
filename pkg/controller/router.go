package controller

import (
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/mux"
	"net/http"
	"project/skillbox/Diplom/pkg/result"
	"project/skillbox/Diplom/simulator"
	"time"
)

func Build(router *mux.Router) {
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	ticker := time.NewTicker(30 * time.Second)
	res := result.GetRes()
	go func() {
		for range ticker.C {
			simulator.Shuffle()
			res = result.GetRes()
		}
	}()
	router.HandleFunc("/api", HandleConnection(&res))
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))

}

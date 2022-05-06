package app

import (
	"net/http"
	"project/skillbox/Diplom/pkg/controller"
	"project/src/github.com/gorilla/mux"
)

func Run() {
	router := mux.NewRouter()
	controller.Build(router)
	http.ListenAndServe(":3000", router)

}

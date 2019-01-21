package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type HTTPPool struct {
}

func HTTPConnect() {

	var httpPool HTTPPool

	r := mux.NewRouter()
	r.HandleFunc("/create", httpPool.Create).Methods("POST")
	r.HandleFunc("/generate", httpPool.Generate).Methods("POST")
	r.HandleFunc("/list", httpPool.List).Methods("GET")
	r.HandleFunc("/delete", httpPool.Delete).Methods("POST")

	err := http.ListenAndServe("localhost:"+config.httpPort, r)

	if err != nil {
		panic(err)
	}

}

var (
	app = new(App)
)

func (http *HTTPPool) Create(rw http.ResponseWriter, rq *http.Request) {
	fmt.Println("HTTP:Create")
}

func (http *HTTPPool) Generate(rw http.ResponseWriter, rq *http.Request) {
	fmt.Println("HTTP:Generate")
}

func (http *HTTPPool) List(rw http.ResponseWriter, rq *http.Request) {
	fmt.Println("HTTP:List")
}

func (http *HTTPPool) Delete(rw http.ResponseWriter, rq *http.Request) {
	fmt.Println("HTTP:Delete")
}

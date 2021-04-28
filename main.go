package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func smplHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "time to hit the hay.")
}

func smplDate(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, time.Now())
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", smplHandler).Methods("GET")
	r.HandleFunc("/date", smplDate).Methods("GET")

	println("smplsrv is listening on 3090...")
	log.Fatal(http.ListenAndServe(":3090", r))
}

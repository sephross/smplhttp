package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func getStatus(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "status:okay")
}

func getDateTime(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, time.Now())
}

func exitContd(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "status:exit")
	os.Exit(1)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", getStatus).Methods("GET")
	r.HandleFunc("/time", getDateTime).Methods("GET")
	r.HandleFunc("/exit", exitContd).Methods("GET")

	println("smplsrv is listening on 9090...")
	log.Fatal(http.ListenAndServe(":9090", r))
}

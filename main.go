package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func headerLogger(w http.ResponseWriter, r *http.Request) {
	var headerString string
	for k, v := range r.Header {
		for _, i := range v {
			headerString += fmt.Sprintf("[%s:%s] ", k, i)
		}
	}

	log.Info("HEADERS: ", headerString)

}

func main() {
	port := ":8080"
	http.HandleFunc("/header", headerLogger)

	log.Info("Starting server on port 8080")

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Error("Error in http.ListenAndServe", err)
	}

	log.Info("closing server")
}

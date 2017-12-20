package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func makeHandler(fn func(http.ResponseWriter, *http.Request, *Processor), p *Processor) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL)
		fn(w, r, p)
	}
}

func apiReadInstances(w http.ResponseWriter, r *http.Request, p *Processor) {
	instances, err := p.ReadInstances()
	if err != nil {
		log.Println(err)
		return
	}

	payload := Payload{Status: "success", Instances: instances}

	data, err := json.MarshalIndent(&payload, "", "\t")
	if err != nil {
		log.Println(err)
	}
	fmt.Fprint(w, string(data))
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
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

func apiReadInstanceStatsHistory(w http.ResponseWriter, r *http.Request, p *Processor) {
	r.ParseForm()
	var id int
	var idStr string
	if _, ok := r.Form["id"]; !ok {
		return
	}
	idStr = r.Form["id"][0]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		return
	}
	statsHistory, err := p.ReadInstanceStatsHistory(id)
	if err != nil {
		log.Println(err)
		return
	}

	payload := Payload{Status: "success", StatsHistory: statsHistory}

	data, err := json.MarshalIndent(&payload, "", "\t")
	if err != nil {
		log.Println(err)
	}
	fmt.Fprint(w, string(data))
}

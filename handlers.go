package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"io"
	"bytes"
)

type JobBuild struct {
	Number int `json:"number"`
	QueueId int `json:"queue_id"`
	Phase string `json:"phase"`
	Status string `json:"status"`
	Url string `json:"url"`
	Log string `json:"log"`
}

type JobUpdate struct {
	Name string `json:"name"`
	Url string `json:"url"`
	Build JobBuild `json:"build"`
}

func JobUpdates(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", body)

	decoder := json.NewDecoder(bytes.NewReader(body))
	var update JobUpdate
	if err := decoder.Decode(&update); err != nil {
		panic(err)
	}

	//if update.Build.Phase == "COMPLETED" {
	//	job := GetByPath(update.Url)
	//
	//	UpdateStatus(job.Id, update.Build.Status)
	//}

	h.broadcast <- body
}

func WSTest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ws test\n")

	h.broadcast <- []byte("Hey this is a message")
}
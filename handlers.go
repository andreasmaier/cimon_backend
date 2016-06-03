package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"io"
	"encoding/json"
)

type JobBuild struct {
	Number int `json:"number"`
	QueueId int `json:"queue_id"`
	Phase string `json:"phase"`
	Url string `json:"url"`
	Log string `json:"log"`
}

type JobUpdate struct {
	Name int `json:"name"`
	//Url string `json:"url"`
	//Build JobBuild `json:"build"`
}

func JobUpdates(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", body)

	//decoder := json.NewDecoder(r.Body)
	//var update JobUpdate
	//if err := decoder.Decode(&update); err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(update)

	h.broadcast <- body
}

func WSTest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ws test\n")

	h.broadcast <- []byte("Hey this is a message")
}

func Jobs(w http.ResponseWriter, r *http.Request) {
	//jenkins, err := gojenkins.CreateJenkins("http://192.168.99.100:8080/", "andi", "changeme").Init()
	//
	//if err != nil {
	//	panic("Something Went Wrong")
	//}
	//
	//jobs, err := jenkins.GetAllJobNames()
	//if err != nil {
	//	panic(err)
	//}
	//
	//for _, job := range jobs {
	//	if job.Color == "" {
	//		nestedJob, err := jenkins.GetJob(job.Name);
	//		if err != nil {
	//			panic(err)
	//		}
	//		b, _ := json.Marshal(nestedJob.GetDetails())
	//		s := string(b)
	// 		fmt.Println(s)
	//	}
	//}
	//
	//fmt.Printf("length %v", len(jobs))
	//
	//if err := json.NewEncoder(w).Encode(jobs); err != nil {
	//	panic(err)
	//}
	json.NewEncoder(w).Encode(getAllJobs())
}

func AddJob(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		fmt.Println("options")

		w.WriteHeader(http.StatusOK)
	} else {
		fmt.Println("add job")

		decoder := json.NewDecoder(r.Body)
		var job JenkinsJob
		if err := decoder.Decode(&job); err != nil {
			panic(err)
		}

		storeJob(job)

		w.WriteHeader(http.StatusCreated)
	}
}

func GetAllJob(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all job")

	json.NewEncoder(w).Encode(getAllJobs())

	//w.WriteHeader(http.StatusOK)
}
package main

import (
	"database/sql"
	_ "github.com/ziutek/mymysql/godrv"
	"fmt"
)

type JenkinsJob struct {
	Id int `json:"id"`
	Server string `json:"server"`
	Path string `json:"path"`
	Status string `jsonL:"status"`
}

func storeJob(job JenkinsJob) {
	con, err := sql.Open("mymysql", "cimon_dev/cimon/changeme")
	defer con.Close()

	if err != nil {
		panic(err)
	}

	var path string
	if err := con.QueryRow("SELECT path from jobs where path=?", job.Path).Scan(&path); err != nil {
		fmt.Println("No job found with that path")

		if _, err := con.Exec("INSERT INTO jobs (path, status) values (?, 'undefined')", job.Path); err != nil {
			panic(err)
		} else {
			fmt.Println("Added job to database")
		}
	}
}

//func isJobWatched(path string) bool {
//	con, err := sql.Open("mymysql", "cimon_dev/cimon/changeme")
//	defer con.Close()
//
//	if err != nil {
//		panic(err)
//	}
//
//	return (con.QueryRow("SELECT path from jobs where path=?", path).Scan(&path) == nil)
//}

func getAllJobs() []*JenkinsJob {
	con, err := sql.Open("mymysql", "cimon_dev/cimon/changeme")
	defer con.Close()

	if err != nil {
		panic(err)
	}

	var jobs []*JenkinsJob
	rows, err := con.Query("SELECT * from jobs")

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		job := new(JenkinsJob)
		if err = rows.Scan(&job.Id, &job.Path, &job.Status); err != nil {
			panic(err)
		}
		jobs = append(jobs, job)
	}

	return jobs
}
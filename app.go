package main

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
)

func connect() (redis.Conn, error) {
	return redis.Dial("tcp", ":6379")
}

const (
	noTimeout = 0
)

func ListenForJobs(jobs chan Job) {
	conn, err := connect()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("Waiting for jobs...")

	for {
		// pop a value from the queue
		reply, err := redis.Values(conn.Do("blpop", "x:queue:default", noTimeout))
		if err != nil {
			log.Fatal(err)
		}

		// unpackage the reply
		var queue string
		var body string
		if _, err := redis.Scan(reply, &queue, &body); err != nil {
			log.Println(err)
		}

		job := new(Job)
		bytes := []byte(body)
		err = json.Unmarshal(bytes, &job)
		if err != nil {
			log.Print(err)
		}

		fmt.Printf("Found job: %s(%s)\n", job.Class, job.Jid)
		jobs <- *job
	}
}

func PerformJobs(jobs chan Job) {
	workers := make(map[string]func(Job) Worker)

	// register our worker
	workers["PlainOldRuby"] = func(job Job) Worker {
		h := Hardwork{&job, "fail", 0}
		return &h
	}

	for {
		job := <-jobs
		createWorker := workers[job.Class]
		worker := createWorker(job)
		go worker.Perform()
	}
}

func main() {
	jobs := make(chan Job)
	go ListenForJobs(jobs)
	go PerformJobs(jobs)

	// wait for user input
	fmt.Println("Press Enter to Exit.")
	var userInput string
	fmt.Scanln(&userInput)
}

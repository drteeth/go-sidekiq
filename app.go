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

func ListenForJobs(jobs chan Job) {
	// connect to redis
	conn, err := connect()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	fmt.Println("Waiting for jobs...")

	for {
		// pop a value from the queue
		reply, err := redis.Values(conn.Do("blpop", "x:queue:default", 0))
		if err != nil {
			log.Fatal(err)
		}

		// unpackage the reply
		var queue string
		var body string
		if _, err := redis.Scan(reply, &queue, &body); err != nil {
			log.Println(err)
		}

		// parse the json
		job := new(Job)
		bytes := []byte(body)
		err = json.Unmarshal(bytes, &job)
		if err != nil {
			log.Print(err)
		}

		// stuff it down the channel
		fmt.Printf("Found job: %s(%s)\n", job.Class, job.Jid)
		jobs <- *job
	}
}

type WorkerFactory func(Job) Worker

func PerformJobs(jobs chan Job) {
	// keep a map of worker factories by name
	workers := make(map[string]WorkerFactory)

	// register our worker
	workers["PlainOldRuby"] = NewHardwork

	for {
		// wait for a job
		job := <-jobs

		// create an instance of the appropriate worker
		factory := workers[job.Class]
		worker := factory(job)

		// do the work asynchronously
		go worker.Perform()
	}
}

func main() {
	jobs := make(chan Job)
	go ListenForJobs(jobs)
	go PerformJobs(jobs)

	fmt.Println("Press Enter to Exit.")
	var userInput string
	fmt.Scanln(&userInput)
}

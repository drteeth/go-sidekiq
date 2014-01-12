package main

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
)

// {
//   "retry": true,
//   "queue": "default",
//   "class": "PlainOldRuby",
//   "args": ["like a dog",3],
//   "jid": "f6e2ab138b6c591989ede8c4",
//   "enqueued_at": 1389458108.727456
// }

type Job struct {
	Retry       bool          `json:retry`
	Queue       string        `json:queue`
	Class       string        `json:class`
	Args        []interface{} `json:args`
	Jid         string        `json:jid`
	Enqueued_at float64       `json:enqueued_at` //float32??
}

func connect() (redis.Conn, error) {
	return redis.Dial("tcp", ":6379")
}

const (
	noTimeout = 0
)

func main() {
	// Connect
	conn, err := connect()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// connect to redis
	// watch the user provided queue (or the default queue)
	// pop a job off and process it.
	// put failing jobs back and decrement

	// use blpop to get items
	// pull from x:queue:default
	// default queue ^^
	for {
		fmt.Println("waiting for redis...")
		reply, err := redis.Values(conn.Do("blpop", "x:queue:default", 0))

		if err != nil {
			log.Printf("xxx %s\n", err)
		}

		var queue string
		var body string
		if _, err := redis.Scan(reply, &queue, &body); err != nil {
			log.Println(err)
		}

		fmt.Println(body)

		job := new(Job)
		bytes := []byte(body)
		err = json.Unmarshal(bytes, &job)
		if err != nil {
			log.Print(err)
		}
	}
}

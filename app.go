package main

import (
	// "encoding/json"
	// "fmt"
	// "github.com/garyburd/redigo/redis"
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
	Retry       bool     `json:retry`
	Queue       string   `json:queue`
	Class       string   `json:class`
	Args        []string `json:args`
	Jid         string   `json:jid`
	Enqueued_at float64  `json:enqueued_at` //float32??
}

func connect() (redis.PubSubConn, error) {
	var conn redis.PubSubConn
	c, err := redis.Dial("tcp", ":6379")

	if err != nil {
		return conn, err
	}

	conn = redis.PubSubConn{c}
	return conn, nil
}

func main() {
	// Connect
	conn, err := connect()
	if err != nil {
		log.Fatal(err)
	}

	// connect to redis
	// watch the user provided queue (or the default queue)
	// pop a job off and process it.
	// put failing jobs back and decrement

	// use blpop to get items
}

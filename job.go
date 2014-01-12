package main

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

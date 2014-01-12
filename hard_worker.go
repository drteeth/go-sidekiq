package main

import (
	"fmt"
	"time"
)

type Hardwork struct {
	*Job
	HowHard string
	HowLong float64
}

func (work *Hardwork) Perform() {
	fmt.Printf("Working %s for %g days/week\n", work.HowHard, work.HowLong)
	time.Sleep(time.Duration(work.HowLong) * time.Second)
	fmt.Printf("Finished: %s\n", work.Jid)
}

func NewHardwork(job Job) Worker {
	howHard := job.Args[0].(string)
	howLong := job.Args[1].(float64)
	return &Hardwork{&job, howHard, howLong}
}

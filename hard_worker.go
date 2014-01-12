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

func (work *Hardwork) ParseArgs() {
	work.HowHard = work.Args[0].(string)
	work.HowLong = work.Args[1].(float64)
}

func (work *Hardwork) Perform() {
	work.ParseArgs()
	fmt.Printf("Working %s for %g days/week\n", work.HowHard, work.HowLong)
	time.Sleep(time.Duration(work.HowLong) * time.Second)
	fmt.Printf("Finished: %s\n", work.Jid)
}

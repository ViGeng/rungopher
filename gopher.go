package main

import "time"

func run(coreload *CoreLoad) {
	timeUnit := time.Millisecond
	idleDuration := timeUnit * time.Duration(100-coreload.LoadPercent)
	runningDuration := timeUnit * time.Duration(coreload.LoadPercent)

	go gopher(idleDuration, runningDuration)
	time.Sleep(time.Duration(coreload.Duration)) // fixme: add a channel to close the gopher
	// todo: close the gopher
}

func gopher(idleDuration, runningDuration time.Duration) {
	for {
		start := time.Now()
		for {
			if time.Now().Sub(start) > runningDuration {
				break
			}
		}
		time.Sleep(idleDuration)
	}
}
package main

import (
	"golang.org/x/sys/unix"
	"time"
)

func run(coreload *CoreLoad) {
	timeUnit := time.Millisecond
	idleDuration := timeUnit * time.Duration(100-coreload.LoadPercent)
	runningDuration := timeUnit * time.Duration(coreload.LoadPercent)

	cpuSet := new(unix.CPUSet)
	cpuSet.Set(int(coreload.CoreNO))

	go gopher(idleDuration, runningDuration, cpuSet)
	time.Sleep(time.Duration(coreload.Duration)) // fixme: add a channel to close the gopher
	// todo: close the gopher
}

func gopher(idleDuration, runningDuration time.Duration, cpuSet *unix.CPUSet) {
	unix.SchedSetaffinity(0, cpuSet)
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
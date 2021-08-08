package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("os.Args less than 2")
	}
	configFilePath := os.Args[1]
	cpuLoad := ReadConfig(configFilePath)
	for _, core := range cpuLoad.Coreloads {
		run(core)
	}
	fmt.Println(cpuLoad)
}


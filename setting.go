package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type CoreLoad struct {
	CoreNO      uint
	LoadPercent uint
	Duration    uint
}

type CPULoad struct {
	Coreloads []*CoreLoad
}

func (c *CoreLoad) isValid() bool {
	return true
}

func (c *CPULoad) isValid() bool {
	for _, c := range c.Coreloads {
		if !c.isValid() {
			return false
		}
	}
	return true
}

func ReadConfig(configFilePath string) *CPULoad {
	cpuLoad := new(CPULoad)

	configJsonBytes, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		log.Fatalln("Reading config file err: ", err)
	}

	if err := json.Unmarshal(configJsonBytes, cpuLoad); err != nil {
		log.Fatalln("Unmarshal config err: ", err)
	}

	if !cpuLoad.isValid() {
		log.Fatalln("Config parameters unvalid.")
	}
	return cpuLoad
}

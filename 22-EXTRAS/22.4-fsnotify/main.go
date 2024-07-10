package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
)

type DBConfig struct {
	Host     string `json:"host"`
	DB       string `json:"db"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var config DBConfig

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	defer watcher.Close()

	MarshalConfig("config.json")

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Write == fsnotify.Write {
					MarshalConfig("config.json")
					log.Println("Config file changed: ", event.Name)
					log.Println(config)
				}
			case err := <-watcher.Errors:
				if err != nil {
					panic(err)
				}
			}
		}
	}()
	err = watcher.Add("config.json")
	if err != nil {
		panic(err)
	}
	<-done
}

func MarshalConfig(file string) {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}
}

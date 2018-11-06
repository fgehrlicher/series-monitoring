package main

import (
	"bitbucket.org/fabian_gehrlicher/series-watcher-v3/Serieswatcher"
	"time"
)

func main() {
	 //wait for the mysql docker container to start
	 //@TODO replace this 💩 with proper checking if the mysql container is running
	time.Sleep(5 * time.Second)

	Server.Init()
}

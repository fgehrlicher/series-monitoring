package main

import (
	"gitea.fge.cloud/fabian_gehrlicher/series-watcher-v3/Serieswatcher"
	"time"
)

func main() {
	 //wait for the mysql docker container to start
	 //@TODO replace this 💩 with proper checking if the mysql container is running
	time.Sleep(4 * time.Second)

	Server.Init()
}

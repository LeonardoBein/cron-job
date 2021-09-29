package main

import (
	"time"

	"github.com/LeonardoBein/cron-job/core"
)

func main() {

	core.Run()

	for {
		time.Sleep(2 * time.Second)
	}

}

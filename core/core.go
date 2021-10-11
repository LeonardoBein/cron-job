package core

import (
	"context"
	"flag"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/LeonardoBein/cron-job/config"
	"github.com/LeonardoBein/cron-job/entity"
	"github.com/LeonardoBein/cron-job/lang"
	"github.com/robfig/cron"
)

func createFunc(config entity.CronJob, logger *log.Logger) func() {
	logger.Println(config.Command)
	return func() {

		if config.IsRunning && config.MultiProcessingLimit <= 1 {
			logger.Printf("WARN\tScript %s isRunning\n", config.Command)
			return
		}

		if config.MultiProcessingLimit != 0 && config.MultiProcessingCount == config.MultiProcessingLimit {
			logger.Printf("WARN\tScript %s running in %d processes, in order to wait for some process to end\n", config.Command, config.MultiProcessingLimit)
			return

		}

		config.IsRunning = true
		config.MultiProcessingCount++

		ctx := context.Background()

		if config.TimeOut > 0 {
			var cancel context.CancelFunc
			ctx, cancel = context.WithTimeout(context.Background(), time.Duration(config.TimeOut)*time.Second)
			defer cancel()
		}

		cmd := exec.CommandContext(ctx, config.Command[0])

		cmd.Args = config.Command

		startCommand := time.Now()

		result, err := cmd.Output()

		if err != nil {
			logger.Printf("ERROR\t%s\t%s\t%s\n", time.Since(startCommand), strings.Join(config.Command, " "), err)
			config.IsRunning = false
			config.MultiProcessingCount--
			return
		}

		config.IsRunning = false
		config.MultiProcessingCount--

		logger.Printf("INFO\t%s\t%s\t\"%s\"\n", time.Since(startCommand), strings.Join(config.Command, " "), string(result))

	}
}

func Run() {

	pathFile := flag.String("path", "/usr/local/etc/cron-job", "a string var")

	flag.Parse()

	if pathFile == nil {
		panic(lang.ErrConfigPathRequired)
	}

	allConfig := config.GetConfig(*pathFile)

	c := cron.New()

	for i := 0; i < len(allConfig); i++ {

		for j := 0; j < len(allConfig[i].Scripts); j++ {

			var writerLog io.Writer

			writerLog = io.MultiWriter(os.Stdout)

			if allConfig[i].LogFile != nil {
				logFile, err := os.Create(*allConfig[i].LogFile)

				if err != nil {
					panic(err)
				}

				writerLog = io.MultiWriter(os.Stdout, logFile)
			}

			logger := log.New(writerLog, "", log.LstdFlags)

			c.AddFunc(allConfig[i].Scripts[j].Spec, createFunc(allConfig[i].Scripts[j], logger))
		}

	}

	c.Start()

}

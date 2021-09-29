package core

import (
	"context"
	"flag"
	"fmt"
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

func createFunc(config *entity.CronJob, logger *log.Logger) func() {
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

	for _, cnf := range allConfig {
		fmt.Println(cnf.Scripts)

		for i := 0; i < len(cnf.Scripts); i++ {

			var writerLog io.Writer

			writerLog = io.MultiWriter(os.Stdout)

			if cnf.LogFile != nil {
				logFile, err := os.Create(*cnf.LogFile)

				if err != nil {
					panic(err)
				}

				writerLog = io.MultiWriter(os.Stdout, logFile)
			}

			logger := log.New(writerLog, "", log.LstdFlags)

			c.AddFunc(cnf.Scripts[i].Spec, createFunc(&cnf.Scripts[i], logger))
		}

	}

	c.Start()

}

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/Mikhalevich/repeater"
)

type logger struct {
}

func (l logger) Infof(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func main() {
	var (
		count    = flag.Int("c", 3, "repeat count attempt")
		duration = flag.Duration("d", time.Second*1, "wait timeout between failed attemps")
	)

	flag.Parse()

	cmdArg := flag.Arg(0)
	if cmdArg == "" {
		log.Println("command not specified")
		os.Exit(1)
	}

	cmd := exec.Command(cmdArg)

	if err := repeater.Do(
		func() error {
			return cmd.Run()
		},
		repeater.WithCount(*count),
		repeater.WithTimeout(*duration),
		repeater.WithLogger(logger{}),
		repeater.WithLogMessage(fmt.Sprintf("run commnad \"%s\"", cmdArg)),
	); err != nil {
		log.Printf("unable to run command \"%s\" error: %v\n", cmdArg, err)
		os.Exit(1)
	}
}

package main

import (
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/exec"
	"strings"
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
		//nolint:mnd
		count    = flag.Int("c", 3, "repeat count attempts")
		duration = flag.Duration("t", time.Second*1, "wait timeout between failed attempts")
	)

	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		slog.Error("command not specified")
		os.Exit(1)
	}

	var (
		cmd     = args[0]
		cmdArgs = args[1:]
	)

	var out []byte

	if err := repeater.Do(
		func() error {
			var err error

			out, err = exec.Command(cmd, cmdArgs...).Output()

			return fmt.Errorf("execute %s: %w", cmd, err)
		},
		repeater.WithAttempts(*count),
		repeater.WithTimeout(*duration),
		repeater.WithLogger(logger{}),
	); err != nil {
		slog.Error("unable to run command", slog.String("cmd", strings.Join(args, " ")), slog.String("error", err.Error()))
		os.Exit(1)
	}

	slog.Info(string(out))
}

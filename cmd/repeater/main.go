package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/Mikhalevich/repeater"
	"github.com/Mikhalevich/repeater/logger"
)

func main() {
	var (
		//nolint:mnd
		count    = flag.Int("c", 3, "repeat count attempts")
		duration = flag.Duration("t", time.Second*1, "wait timeout between failed attempts")
		factor   = flag.Int("f", 0, "factor for exponential timeout backoff")
		jitter   = flag.Bool("j", false, "jitter for exponential timeout backoff")
		log      = slog.New(slog.NewTextHandler(os.Stdout, nil))
	)

	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		log.Error("command not specified")
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
			if err != nil {
				return fmt.Errorf("execute %s: %w", cmd, err)
			}

			return nil
		},
		repeater.WithAttempts(*count),
		repeater.WithTimeout(*duration),
		repeater.WithFactor(*factor),
		repeater.WithJitter(*jitter),
		repeater.WithLogger(logger.NewSLog(log)),
	); err != nil {
		log.Error("unable to run command", slog.String("cmd", strings.Join(args, " ")), slog.String("error", err.Error()))
		os.Exit(1)
	}

	fmt.Fprintln(os.Stdout, string(out))
}

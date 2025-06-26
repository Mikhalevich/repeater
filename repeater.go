package repeater

import (
	"fmt"
	"math"
	"math/rand/v2"
	"time"

	"github.com/Mikhalevich/repeater/logger"
)

const (
	defaultFailureAttempts = 3
)

// Do executes function while receiving error or attempts were exeded.
func Do(doFn func() error, opts ...Option) error {
	params := options{
		Attempts: defaultFailureAttempts,
		Logger:   logger.NewNoop(),
	}

	for _, o := range opts {
		o(&params)
	}

	var err error

	for attempt := range params.Attempts {
		err = doFn()
		if err == nil {
			break
		}

		timeout := calculateTimeout(attempt, params.Timeout, params.Factor, params.IsJitterEnabled)

		params.Logger.WithError(err).Error("failure attempt", "attempt", attempt+1, "timeout", timeout)

		if timeout > 0 {
			time.Sleep(timeout)
		}
	}

	if err != nil {
		return fmt.Errorf("attempts exeded: %w", err)
	}

	return nil
}

func calculateTimeout(
	attempt int,
	timeout time.Duration,
	factor int,
	withJitter bool,
) time.Duration {
	timeoutSecs := calculateTimeoutSeconds(attempt, int(timeout.Seconds()), factor, withJitter)

	return time.Second * time.Duration(timeoutSecs)
}

func calculateTimeoutSeconds(
	attempt int,
	timeout int,
	factor int,
	withJitter bool,
) int {
	if timeout == 0 {
		return 0
	}

	if factor == 0 {
		return timeout
	}

	attemptFactor := int(math.Pow(float64(factor), float64(attempt)))

	timeout *= attemptFactor

	if withJitter {
		//nolint:gosec
		timeout += int(float32(timeout) * rand.Float32())
	}

	return timeout
}

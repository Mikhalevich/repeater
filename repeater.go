package repeater

import (
	"fmt"
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

		params.Logger.WithError(err).Error("failure attempt", "attempt", attempt+1)

		if params.Timeout > 0 {
			time.Sleep(params.Timeout)
		}
	}

	if err != nil {
		return fmt.Errorf("attempts exeded: %w", err)
	}

	return nil
}

package repeater

import (
	"fmt"
	"time"
)

const (
	defaultFailureAttempts = 3
)

// Logger interface for external implementation for logger support.
type Logger interface {
	Infof(format string, args ...interface{})
}

// Do executes function while receiving error.
func Do(doFn func() error, opts ...Option) error {
	params := &options{
		Attempts: defaultFailureAttempts,
	}

	for _, o := range opts {
		o(params)
	}

	var err error
	for attempt := range params.Attempts {
		err = doFn()
		if err == nil {
			break
		}

		if params.Logger != nil {
			params.Logger.Infof("repeating attempt: %d, err: %v\n", attempt+1, err)
		}

		if params.Timeout > 0 {
			time.Sleep(params.Timeout)
		}
	}

	if err != nil {
		return fmt.Errorf("attempts exeded: %w", err)
	}

	return nil
}

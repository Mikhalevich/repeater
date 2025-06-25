package repeater

import (
	"time"
)

type options struct {
	Attempts int
	Timeout  time.Duration
	Logger   Logger
}

// Option is a specific option for repeater.
type Option func(opts *options)

// WithAttempts set max number of attempts in case of failure(3 by default).
func WithAttempts(count int) Option {
	return func(opts *options) {
		opts.Attempts = count
	}
}

// WithTimeout set a timeout between failure attempts(no timeout by default).
func WithTimeout(d time.Duration) Option {
	return func(opts *options) {
		opts.Timeout = d
	}
}

// WithLogger set custom logger (no logger by default).
func WithLogger(logger Logger) Option {
	return func(opts *options) {
		opts.Logger = logger
	}
}

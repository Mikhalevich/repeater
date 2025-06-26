package repeater

import (
	"time"

	"github.com/Mikhalevich/repeater/logger"
)

type options struct {
	Attempts        int
	Timeout         time.Duration
	Factor          int
	IsJitterEnabled bool
	Logger          logger.Logger
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

// WithFactor set factor for exponential backoff timeout calculation(0 by default).
func WithFactor(factor int) Option {
	return func(opts *options) {
		opts.Factor = factor
	}
}

// WithJitter enable/disable jitter for exponential backoff timeout calculation(false by default).
func WithJitter(enabled bool) Option {
	return func(opts *options) {
		opts.IsJitterEnabled = enabled
	}
}

// WithLogger set custom logger (no logger by default).
func WithLogger(log logger.Logger) Option {
	return func(opts *options) {
		opts.Logger = log
	}
}

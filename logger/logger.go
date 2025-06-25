package logger

// Logger interface for external implementation to enable logging for outbox.
type Logger interface {
	Debug(msg string, args ...any)
	Error(msg string, args ...any)

	WithError(err error) Logger
}

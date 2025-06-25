package logger

import "log/slog"

// SLog is slog implementation of Logger interface.
type SLog struct {
	log *slog.Logger
}

func NewSLog(log *slog.Logger) *SLog {
	return &SLog{
		log: log,
	}
}

func (s *SLog) Error(msg string, args ...any) {
	s.log.Error(msg, args...)
}

func (s *SLog) WithError(err error) Logger {
	return &SLog{
		log: s.log.With("error", err.Error()),
	}
}

package logger

// Noop is no operations imlementation of Logger interface.
type Noop struct {
}

func NewNoop() *Noop {
	return &Noop{}
}

func (nw *Noop) Debug(msg string, args ...any) {
}

func (nw *Noop) Error(msg string, args ...any) {
}

func (nw *Noop) WithError(err error) Logger {
	return nw
}

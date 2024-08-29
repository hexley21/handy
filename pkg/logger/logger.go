package logger

type Logger interface {
	Info(msg string, args ...any)
	Debug(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(err error, args ...any)
	ErrorCause(err error, cause any)
	Fatal(err error, args ...any)
}

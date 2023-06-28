package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

var l *logrus.Logger

func init() {
	l = logrus.New()
	l.Out = os.Stdout
	l.SetFormatter(&logrus.TextFormatter{})
}

type Logger interface {
	Info(...any)
	Warn(...any)
	Error(...any)
	Panic(...any)
	L() *logrus.Logger
}

// TODO
type logger struct {
	l *logrus.Logger
}

func (l *logger) Info(args ...any) {
	l.l.Info(args...)
}
func (l *logger) Warn(args ...any) {
	l.l.Warn(args...)
}
func (l *logger) Error(args ...any) {
	l.l.Error(args...)
}
func (l *logger) Panic(args ...any) {
	l.l.Panic(args...)
}
func (l *logger) L() *logrus.Logger{ 
	return l.l
}

func NewLogger() Logger {
	l := logrus.New()
	// l.SetFormatter(&logrus.JSONFormatter{})
	l.SetFormatter(&logrus.TextFormatter{
		ForceColors:               true,
		FullTimestamp:             true,
	})
	// l.SetReportCaller(true)

	l.Out = os.Stdout

	return &logger{
		l: l,
	}
}

// todo: testing
func Info(args ...any) {
	l.Info(args...)
}

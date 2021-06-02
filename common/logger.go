package common

import (
	"log"
	"os"
	"time"
)

type Logger struct {
	Service string
	Verb string
	Start time.Time
	Duration time.Duration
}

func StartLog(service, verb string) *Logger {
	log := Logger{Service: service, Verb: verb}

	log.Start = time.Now()
	log.Log("start")

	return &log
}

func (l *Logger) NormalReturn() {
	l.Duration = time.Since(l.Start)
	l.LogWithDuration("ok")
}

func (l *Logger) FailedReturn() {
	l.Duration = time.Since(l.Start)
	l.LogWithDuration("fail")
}

func (l *Logger) Log(message string) {
	host, err := os.Hostname()
	if err != nil {
		host = "unknown"
	}

	log.Printf("%s.%s.%s.%s", host, l.Service, l.Verb, message)
}

func (l *Logger) LogWithDuration(message string) {
	host, err := os.Hostname()
	if err != nil {
		host = "unknown"
	}

	log.Printf("%s.%s.%s.%s (%s)", host, l.Service, l.Verb, message, l.Duration.String())
}


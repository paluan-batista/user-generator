package log

import (
	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"os"
)

func NewLogger() *log.Logger {
	return &log.Logger{
		Handler: cli.New(os.Stdout),
		Level:   1,
	}
}

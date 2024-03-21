package h

import (
	"io"
	"log/slog"
	"os"
	"sync"
)

var Log *slog.Logger
var once sync.Once

type LoggerOptions struct {
	Channel io.Writer
}

func LogInit(opt LoggerOptions) *slog.Logger {
	once.Do(func() {
		if opt.Channel == nil {
			opt.Channel = os.Stdout
		}

		Log = slog.New(slog.NewJSONHandler(opt.Channel, nil))
		Log.Debug("log created")
	})

	return Log
}

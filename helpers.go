package jig

import (
	"fmt"
	"os"

	"github.com/bontusss/jig/log"
)

// CreateDirIfNotExist creates a new directory if it does not exist
func (g *Jig) createDirIfNotExist(path string) error {
	const mode = 0755
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, mode)
		if err != nil {
			return err
		}
	}

	return nil
}

// CreateFileIfNotExists creates a new file at path if it does not exist
func (g *Jig) createFileIfNotExists(path string) error {
	var _, err = os.Stat(path)
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			return err
		}

		defer func(file *os.File) {
			_ = file.Close()
		}(file)
	}
	return nil
}

func (g *Jig) checkDotEnv(path string) error {
	err := g.createFileIfNotExists(fmt.Sprintf("%s/.env", path))
	if err != nil {
		return err
	}
	return nil
}

// InitializeLogger initializes the logger for the Jig application with a console handler and a file handler.
// It sets the log level based on the debug flag of the Jig instance.
// Usage:
//        err := jig.InitializeLogger()
//        if err != nil {
//            // handle error
//        }
//		  jig.JigLogger.Debug("This is a debug message from Jig")
// 		  jig.JigLogger.Info("This is an info message from Jig")
// 		  jig.JigLogger.Warn("This is a warning message from Jig")
// 		  jig.JigLogger.Error("This is an error message from Jig")
func (j *Jig) InitializeLogger(handlerTypes []string) error {
	var handlers []log.Handler
	for _, h := range handlerTypes {
		switch h {
		case "console":
			handlers = append(handlers, &log.ConsoleHandler{})
		case "file":
			fileHandler, err := log.NewFileHandler("app.log")
			if err != nil {
				return err
			}
			handlers = append(handlers, fileHandler)
		default:
			// Handle unknown handler
			return fmt.Errorf("unknown handler: %s", h)
		}
	}

	// slackHandler := log.NewSlackHandler("https://hooks.slack.com/services/YOUR/WEBHOOK/URL")
	var logLevel log.LogLevel
	// When Debug is set to true, the logger should capture and display more detailed (debug-level) log messages.
	if j.Debug {
		logLevel = log.DEBUG
	} else {
		logLevel = log.INFO
	}

	j.Logger = log.NewLogger(logLevel, handlers)
	return nil

}

package log

import (
	"fmt"
	"github.com/patrykjadamczyk/go-status-server/config"
	"log"
	"os"
)

type Level int

const Debug = Level(0b10000)
const Verbose = Level(0b1000)
const Info = Level(0b100)
const Warn = Level(0b10)
const Error = Level(0b1)

type Output int

const OutputStdout = Output(0b0)
const OutputLogfile = Output(0b1)

type Logger struct {
	logLevel  Level
	output    Output
	logFile   string
	logLogger *log.Logger
}

type ParametersArray []interface{}

func (l *Logger) canLog(level Level) bool {
	return l.logLevel >= level
}

func (l *Logger) levelToString(level Level) string {
	switch level {
	case Debug:
		return "Debug"
	case Verbose:
		return "Verbose"
	case Info:
		return "Info"
	case Warn:
		return "Warn"
	case Error:
		return "Error"
	default:
		return ""
	}
}

func (l *Logger) prefix(level Level) string {
	levelString := l.levelToString(level)
	if levelString == "" {
		return "GSS: "
	}
	return fmt.Sprintf("GSS[%s]: ", levelString)
}

func (l *Logger) prepareLog() *log.Logger {
	if l.logLogger != nil {
		return l.logLogger
	}
	logFile, err := os.OpenFile(l.logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Logging Error:", err)
		panic(err)
	}
	defer logFile.Close()
	logger := log.New(logFile, string(nil), int(nil))
	l.logLogger = logger
	return logger
}

func (l *Logger) Log(level Level, params ...interface{}) {
	if !l.canLog(level) {
		return
	}
	prefix := l.prefix(level)
	parameters := ParametersArray{prefix}
	parameters = append(parameters, params...)
	if l.output == OutputLogfile {
		logLogger := l.prepareLog()
		logLogger.Println(parameters...)
	} else {
		fmt.Println(parameters...)
	}
}

func (l *Logger) Debug(params ...interface{}) {
	l.Log(Debug, params...)
}

func (l *Logger) Verbose(params ...interface{}) {
	l.Log(Verbose, params...)
}

func (l *Logger) Info(params ...interface{}) {
	l.Log(Info, params...)
}

func (l *Logger) Warn(params ...interface{}) {
	l.Log(Warn, params...)
}

func (l *Logger) Error(params ...interface{}) {
	l.Log(Error, params...)
}

func (l *Logger) VerboseLevelToLogLevel(level config.VerboseLevel) {
	var newLevel Level
	switch level {
	case config.VerboseLevelDebug:
		newLevel = Debug
	case config.VerboseLevelVerbose:
		newLevel = Verbose
	case config.VerboseLevelWarnings:
		newLevel = Warn
	case config.VerboseLevelError:
		newLevel = Error
	default:
		newLevel = Warn
	}

	l.logLevel = newLevel
}

func NewLogManager(level config.VerboseLevel, logFile string) Logger {
	logManager := Logger{}
	if logFile != "" {
		logManager.logFile = logFile
		logManager.output = OutputLogfile
	} else {
		logManager.output = OutputStdout
	}
	logManager.VerboseLevelToLogLevel(level)
	return logManager
}

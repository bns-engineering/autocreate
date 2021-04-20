package logging

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	log "github.com/Sirupsen/logrus"
)

var dir string

//SetVariable returns logger object with variables
func SetVariable(v map[string]interface{}) ILogger {
	varLog := &Logger{
		fields: v,
	}

	return varLog
}

// InfoLn output Info
func (lg *Logger) InfoLn(msg ...interface{}) {
	fmt.Println(msg...)
	log.WithFields(lg.getFields()).Infoln(msg...)
}

// Infof info with format
func (lg *Logger) Infof(msgFormat string, msg ...interface{}) {
	fmt.Println(fmt.Sprintf(msgFormat, msg...))
	log.WithFields(lg.getFields()).Infof(msgFormat, msg...)
}

// Error output error
func (lg *Logger) Error(err error, msg ...interface{}) {
	fmt.Println(msg...)
	log.WithFields(lg.getFields(err)).Errorln(msg...)
}

// Errorf output error with format
func (lg *Logger) Errorf(err error, msgFormat string, msg ...interface{}) {
	fmt.Printf(msgFormat, msg...)
	log.WithFields(lg.getFields(err)).Errorf(msgFormat, msg...)
}

// Panic log
func (lg *Logger) Panic(msg ...interface{}) {
	fmt.Println(msg...)
	log.WithFields(lg.getFields()).Panicln(msg...)
}

// Fatal log
func (lg *Logger) Fatal(msg ...interface{}) {
	fmt.Println(msg...)
	log.WithFields(lg.getFields()).Fatalln(msg...)
}

// Debugln log debug
func (lg *Logger) Debugln(msg ...interface{}) {
	fmt.Println(msg...)
	log.WithFields(lg.getFields()).Debugln(msg...)
}

//Debugf debug with format
func (lg *Logger) Debugf(msgFormat string, msg ...interface{}) {
	fmt.Printf(msgFormat, msg...)
	log.WithFields(lg.getFields()).Debugf(msgFormat, msg...)
}

// InfoLn output Info
func InfoLn(msg ...interface{}) {
	logger.InfoLn(msg...)
}

// Infof info with format
func Infof(msgFormat string, msg ...interface{}) {
	logger.Infof(msgFormat, msg...)
}

// Error output error
func Error(err error, msg ...interface{}) {
	logger.Error(err, msg...)
}

// Errorf output error with format
func Errorf(err error, msgFormat string, msg ...interface{}) {
	logger.Errorf(err, msgFormat, msg...)
}

// Panic log
func Panic(msg ...interface{}) {
	logger.Panic(msg...)
}

// Fatal log
func Fatal(msg ...interface{}) {
	logger.Fatal(msg...)
}

// Debugln log debug
func Debugln(msg ...interface{}) {
	logger.Debugln(msg...)
}

//Debugf debug with format
func Debugf(msgFormat string, msg ...interface{}) {
	logger.Debugf(msgFormat, msg...)
}

//OpenFileLog open file location based on env
func OpenFileLog() {
	var filename string
	env := os.Getenv("ENV")
	filename = strings.Replace(filepath.Base(os.Args[0]), ".exe", "", 1)
	if len(env) == 0 {
		filename = filename + ".log"
	} else if env == "DEV" {
		filename = filename + ".log"
	} else if env == "STAGING" {
		filename = filename + ".log"
	} else if env == "PRODUCTION" {
		filename = filename + ".log"
	} else {
		log.Error("please define env")
	}
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	Formatter := new(log.TextFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05.9999"
	Formatter.FullTimestamp = true
	log.SetFormatter(Formatter)
	if err != nil {
		// Cannot open log file. Logging to stderr
		fmt.Println(err)
	} else {
		log.SetOutput(f)
	}
}

func (lg *Logger) getFields(err ...error) (fields log.Fields) {
	fields = log.Fields{}

	if len(err) > 0 {
		fields["err"] = err[0]
	}

	callDepth := 4
	if lg.fields != nil {
		for k, v := range lg.fields {
			fields[k] = v
		}
		callDepth = 3
	}

	fields["file"] = getParentCaller(callDepth)

	return fields
}

func getParentCaller(callDepth int) string {
	var buffer bytes.Buffer

	pc, file, line, ok := runtime.Caller(callDepth)
	fnc := runtime.FuncForPC(pc)
	if ok {
		if len(strings.Split(file, "/bns-engineering/platformbanking-presenter/")) > 1 {
			buffer.WriteString(strings.Split(file, "/bns-engineering/platformbanking-presenter/")[1])
		} else {
			buffer.WriteString(file)
		}
		buffer.WriteString(":")
		buffer.WriteString(strconv.Itoa(line))
		buffer.WriteString(" @")

		funcName := fnc.Name()
		funcName = funcName[strings.LastIndex(funcName, ".")+1:] // remove the function detailed path
		buffer.WriteString(funcName)
	}

	return buffer.String()
}

// EnableDebugMode as it is
func EnableDebugMode() {
	log.SetLevel(log.DebugLevel)
}

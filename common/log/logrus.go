package log

import (
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

// Log custom logger
type Log struct {
	log    *logrus.Logger
	fields Fields
}

// NewLogrus return log object
func NewLogrus(accessFile string) Logger {
	logObj := &Log{}
	accLogger := logrus.New()
	formatter := &logrus.JSONFormatter{}

	//Set Info File
	infoFile, err := os.OpenFile(accessFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println("Could not open log file - switching to normal output")
	}

	if infoFile != nil {
		accLogger.Out = infoFile
	}
	accLogger.Formatter = formatter
	logObj.log = accLogger
	return logObj
}

// Println output Info
func (lg *Log) Println(msg ...interface{}) {
	lg.log.WithFields(lg.getFields()).Infoln(msg...)
}

// Printf info with format
func (lg *Log) Printf(fmt string, msg ...interface{}) {
	lg.log.WithFields(lg.getFields()).Infof(fmt, msg...)
}

// Error output error
func (lg *Log) Error(err error, msg ...interface{}) {
	lg.log.WithFields(lg.getFields(err)).Errorln(msg...)
}

// Errorf output error with format
func (lg *Log) Errorf(err error, fmt string, msg ...interface{}) {
	lg.log.WithFields(lg.getFields(err)).Errorf(fmt, msg...)
}

// Panic log
func (lg *Log) Panic(msg ...interface{}) {
	lg.log.WithFields(lg.getFields()).Panicln(msg...)
}

// Fatal log
func (lg *Log) Fatal(msg ...interface{}) {
	lg.log.WithFields(lg.getFields()).Fatalln(msg...)
}

// Debugln log debug
func (lg *Log) Debugln(msg ...interface{}) {
	lg.log.WithFields(lg.getFields()).Debugln(msg...)
}

// Debugf debug with format
func (lg *Log) Debugf(fmt string, msg ...interface{}) {
	lg.log.WithFields(lg.getFields()).Debugf(fmt, msg...)
}

func (lg *Log) getFields(err ...error) (fields logrus.Fields) {
	fields = logrus.Fields{}

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

// SetVariable returns logger object with variables
func (lg *Log) SetVariable(v Fields) Logger {
	varLog := &Log{
		fields: v,
		log:    lg.log,
	}

	return varLog
}

// EnableDebugMode as it is
func (lg *Log) EnableDebugMode() {
	lg.log.SetLevel(logrus.DebugLevel)
}

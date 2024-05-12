package log

import (
	"bytes"
	"runtime"
	"strconv"
	"strings"
)

type (
	//Logger logging wrapper
	Logger interface {
		Println(...interface{}) //INFO
		Printf(string, ...interface{})
		Error(error, ...interface{})
		Errorf(error, string, ...interface{})
		Panic(...interface{})
		Fatal(...interface{})
		Debugln(...interface{})
		Debugf(string, ...interface{})
		SetVariable(Fields) Logger
	}

	//Fields type for standardize custom field input
	Fields map[string]interface{}
)

var (
	logger Logger
)

// Init loggerj
func Init(accessFile string) {
	logger = NewLogrus(accessFile)
}

// EnableDebugMode - currently implemented by logrus logger
func EnableDebugMode() {
	logger.(*Log).EnableDebugMode()
}

// InitMock logger mock
func InitMock() {
	Init("")
}

// InitNoOutput init without visible output
func InitNoOutput() {
	Init("/dev/null")
}

// Println output Info
func Println(msg ...interface{}) {
	logger.Println(msg...)
}

// Printf info with format
func Printf(fmt string, msg ...interface{}) {
	logger.Printf(fmt, msg...)
}

// Error output error
func Error(err error, msg ...interface{}) {
	logger.Error(err, msg...)
}

// Errorf output error with format
func Errorf(err error, fmt string, msg ...interface{}) {
	logger.Errorf(err, fmt, msg...)
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

// Debugf debug with format
func Debugf(fmt string, msg ...interface{}) {
	logger.Debugf(fmt, msg...)
}

// SetVariable returns logger object with variables
func SetVariable(v map[string]interface{}) Logger {
	return logger.SetVariable(v)
}

func getParentCaller(callDepth int) string {
	var buffer bytes.Buffer

	pc, file, line, ok := runtime.Caller(callDepth)
	fnc := runtime.FuncForPC(pc)
	if ok {
		if len(strings.Split(file, "/go-gin-intikom/")) > 1 {
			buffer.WriteString(strings.Split(file, "/go-gin-intikom/")[1])
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

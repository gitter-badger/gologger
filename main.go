package logger

import (
	"errors"
	"io"
	"fmt"
	"log"
)

var logDirectory ,processName, logLevel, hostName  string
var processID int


//Debug Debug log without formatting
func Debug(message ...interface{}) {
	_logger.Log(1, DEBUG, fmt.Sprint(message...))
}

//Error function for error logs without formatting
func Error(message ...interface{}) {
	_logger.Log(1, ERROR, fmt.Sprint(message...))
}

//Info info level logs without formatting
func Info(message ...interface{}) {
	_logger.Log(1, INFO, fmt.Sprint(message...))
}

//Warning Warning level logs without formatting
func Warning(message ...interface{}) {
	_logger.Log(1, WARNING, fmt.Sprint(message...))
}

//Errorf Prints log with formatting
func Errorf(message ...interface{}) {
	_logger.Log(1, ERROR, fmt.Sprintf(message[0].(string), message[1:]...))
}

//Debugf Prints log with formatting
func Debugf(message ...interface{}) {
	_logger.Log(1, DEBUG, fmt.Sprintf(message[0].(string), message[1:]...))
}

//Warningf Prints log with formatting
func Warningf(message ...interface{}) {
	_logger.Log(1, WARNING, fmt.Sprintf(message[0].(string), message[1:]...))
}

//Infof Prints log with formatting
func Infof(message ...interface{}) {
	_logger.Log(1, INFO, fmt.Sprintf(message[0].(string), message[1:]...))
}


//InitLogger initialise logger object with logWriter and log level
func InitLogger(level, directory, process string) error {
	logLevel = level
	logDirectory = directory
	processName = process 

	_log := &Logger{}
	logWriter, err := GetLogWriter()
	if err != nil {
		log.Println("Failed getting log writer", err.Error())
		return err
	}

	_log.Writer = logWriter
	_log.LogLevel = LogLevels[logLevel]
	_logger = _log
	return nil
}


//GetLogger  returns the current default logger instance
func GetLogger() (*Logger, error) {
	if _logger == nil {
		return nil, errors.New("Nil logger ")
	}
	return _logger,nil
}

//SetLogWriter sets default log writer
//This works only if the default logger is being used
func SetLogWriter(writer io.Writer) error {
	if writer == nil {
		return errors.New("Nil writer")
	}
    logWriter = writer
	return _logger.SetLogWriter(writer)
}

//Flush flushes the data logs to log writer
func Flush() {
	_logger.Flush()
}

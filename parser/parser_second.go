package parser

import (
	"errors"
	"logtomongo/config"
	"logtomongo/db"
	"strings"
	"time"
)

type ParserSecond struct{}

func (p ParserSecond) Parse(fileInfo config.FileInfo, line string) (db.MongoItemLog, error) {
	var mongoDoc = db.MongoItemLog{}

	logTime, err := p.logTimeParse(line)
	if err != nil {
		return mongoDoc, err
	}
	logMsg, err := p.logMsgParse(line)
	if err != nil {
		return mongoDoc, err
	}
	mongoDoc = db.MongoItemLog{
		Log_time:   logTime,
		Log_msg:    logMsg,
		File_name:  fileInfo.Path,
		Log_format: fileInfo.Type,
	}
	return mongoDoc, nil
}

func (p ParserSecond) logTimeParse(line string) (time.Time, error) {
	timeStr := strings.Split(line, " | ")[0]
	return time.Parse("2006-01-02T15:04:05Z", timeStr)
}

func (p ParserSecond) logMsgParse(line string) (string, error) {
	var logMsg string
	var logSlice = strings.Split(line, " | ")
	if len(logSlice) > 1 {
		logMsg = logSlice[1]
		return logMsg, nil
	}
	err := errors.New("Error parse log message")
	return logMsg, err
}

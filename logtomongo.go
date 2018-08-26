package main

import (
	"fmt"
	"logtomongo/config"
	"logtomongo/db"
	"logtomongo/parser"
)

func main() {
	var conf config.Configuration
	conf = config.Config{}
	conf.Init()
	logFiles := conf.GetListFileInfo()
	collection := db.LogCollection()

	ch := make(chan db.MongoItemLog)
	parser.ParseListOfFiles(logFiles, ch)

	for mongoDoc := range ch {
		collection.Upsert(mongoDoc, &mongoDoc)
		fmt.Println(mongoDoc.Log_format, mongoDoc.File_name, mongoDoc.Log_time, mongoDoc.Log_msg)
	}
}

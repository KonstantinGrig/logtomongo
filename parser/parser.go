package parser

import (
	"bufio"
	"log"
	"logtomongo/config"
	"logtomongo/db"
	"os"
	"time"
)

type Parser interface {
	Parse(fileInfo config.FileInfo, line string) (db.MongoItemLog, error)
}

func ParseListOfFiles(logFiles config.ListOfFilesInfo, ch chan db.MongoItemLog) {
	for _, fileInfo := range logFiles {
		go ParseFile(fileInfo, ch)
	}
}

func ParseFile(fileInfo config.FileInfo, ch chan db.MongoItemLog) {
	var lastModTime time.Time
	var lastLine int
	var parser Parser
	switch fileInfo.Type {
	case "first_format":
		parser = ParserFirst{}
	case "second_format":
		parser = ParserSecond{}
	default:
		panic("Undefined log file format")
	}
	for {
		file, err := os.Open(fileInfo.Path)
		if err != nil {
			log.Fatal(err)
		}

		info, _ := os.Stat(fileInfo.Path)
		modTime := info.ModTime()
		if modTime.After(lastModTime) {
			lastModTime = modTime
			scanner := bufio.NewScanner(file)
			var curentLine int
			for scanner.Scan() {
				curentLine += 1
				if curentLine > lastLine {
					line := scanner.Text()
					if line != "" {
						mongoDoc, err := parser.Parse(fileInfo, line)
						if err == nil {
							ch <- mongoDoc
						}
						lastLine = curentLine
					}
				}
			}
			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}
		}
		file.Close()
		time.Sleep(1000 * time.Millisecond)
	}
}

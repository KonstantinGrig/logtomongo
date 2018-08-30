package parser

import (
	"bufio"
	"github.com/fsnotify/fsnotify"
	"log"
	"logtomongo/config"
	"logtomongo/db"
	"os"
	//"time"
)

type Parser interface {
	Parse(fileInfo config.FileInfo, line string) (db.MongoItemLog, error)
}

func ParseListOfFiles(logFiles config.ListOfFilesInfo, ch chan db.MongoItemLog) {
	for _, fileInfo := range logFiles {
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			log.Fatal(err)
		}
		err = watcher.Add(fileInfo.Path)
		if err != nil {
			log.Fatal(err)
		}
		go ParseFile(fileInfo, ch, watcher)
	}
}

func ParseFile(fileInfo config.FileInfo, ch chan db.MongoItemLog, watcher *fsnotify.Watcher) {
	defer watcher.Close()
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
		select {
		case event := <-watcher.Events:
			if event.Op&fsnotify.Write == fsnotify.Write && event.Name == fileInfo.Path {

				log.Println("modified file:", event.Name)

				file, err := os.Open(fileInfo.Path)
				if err != nil {
					log.Fatal(err)
				}

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
						}
					}
				}
				lastLine = curentLine
				if err := scanner.Err(); err != nil {
					log.Fatal(err)
				}
				file.Close()
			}
		case err := <-watcher.Errors:
			log.Println("error:", err)
		}
	}
}

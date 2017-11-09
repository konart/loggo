package main

import (
	"bufio"
	"fmt"
	"github.com/labstack/gommon/log"
	"os"
	"strings"
)

//type logLevel int

//const (
//	INFO logLevel = iota
//	DEBUG
//	WARN
//	ERR
//)
//
//var logLevels []logLevel = []logLevel{INFO, DEBUG, WARN, ERR}
//
//func (ll logLevel) String() string  {
//	logLevels := []string{"INFO", "DEBUG", "WARN", "ERR"}
//	return logLevels[ll]
//}

type record struct {
	level   string
	message string
}

type journal []*record

func checkLogLevel(s string) bool {
	logLevelsMap := map[string]bool{
		"[ERROR]": true,
		"[INFO]":  true,
		"[DEBUG]": true,
		"[WARN]":  true,
	}
	result := logLevelsMap[s]
	return result
}

func parseFile(f *os.File) journal {
	var journal journal
	scanner := bufio.NewScanner(bufio.NewReader(f))
	for scanner.Scan() {
		//fmt.Println(scanner.Text()) // Println will add back the final '\n'
		record := getLogRecord(scanner.Text())
		if record.level != "" {
			//fmt.Println(record.level)
			journal = append(journal, record)
		} else {
			journal[len(journal)-1].message += " " + record.message
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return journal
}

func getLogRecord(s string) *record {
	record := &record{}
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		if checkLogLevel(word) {
			record.level = word
			continue
		}
		if word == "--" {
			continue
		}
		if record.message == "" {
			record.message += word
		} else {
			record.message += " " + word
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return record
}

func main() {
	f, err := os.Open("C:\\logs\\test.log")
	if err != nil {
		log.Fatal(err)
	}
	j := parseFile(f)
	for _, rec := range j {
		fmt.Printf("%s -- %s\n", rec.level, rec.message)
	}
}

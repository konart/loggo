package main

import (
	"bufio"
	"crypto/md5"
	"flag"
	"fmt"
	"github.com/labstack/gommon/log"
	"io"
	"os"
	"regexp"
	"strings"
	"time"
)

type record struct {
	message     string
	dateTime    time.Time
	lineNumbers []int
}

var logStatuses = []string{"[ERROR]", "[INFO]", "[DEBUG]", "[WARN]"}

type newJournalHash map[string]map[string]*record

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

func (r *record) getMd5String() string {
	h := md5.New()
	io.WriteString(h, r.message)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (r *record) setTime(s string) {
	s = strings.Replace(s, ",", ".", -1)
	layout := "2006-01-02 15:04:05.000"
	r.dateTime, _ = time.Parse(layout, s)
}

func cutTimestamp(s string) (string, string) {
	timeRe := regexp.MustCompile("((2\\d\\d\\d)-(0\\d|1[012])-([0-2][1-9]|3[01]) (([01]?\\d)|([2][0-3])):([0-5]?\\d)(:([0-5]?\\d))(,\\d\\d\\d))?")
	timeString := timeRe.FindString(s)
	if timeString != "" {
		s = strings.Replace(s, timeString, "", 1)
	}
	return timeString, s
}

func parseFile(f *os.File) newJournalHash {
	newJournalHash := map[string]map[string]*record{}
	for _, status := range logStatuses {
		newJournalHash[status] = map[string]*record{}
	}
	scanner := bufio.NewScanner(bufio.NewReader(f))
	var lastRecord *record
	i := 0
	for scanner.Scan() {
		i += 1
		rec, logLevel := getLogRecord(scanner.Text())
		if logLevel != "" {

			var timeString string
			timeString, rec.message = cutTimestamp(rec.message)
			if timeString != "" {
				rec.setTime(timeString)
			}

			md5String := rec.getMd5String()
			existingRecord := newJournalHash[logLevel][md5String]

			if existingRecord == nil {
				newJournalHash[logLevel][md5String] = rec
			} else {
				existingRecord.lineNumbers = append(existingRecord.lineNumbers, i)
			}

			rec.lineNumbers = append(rec.lineNumbers, i)
			lastRecord = rec
		} else {
			lastRecord.message += " " + rec.message
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return newJournalHash
}

func getLogRecord(s string) (*record, string) {
	record := &record{}
	var logLevel string
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		if checkLogLevel(word) {
			logLevel = word
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
	return record, logLevel
}

func main() {
	var filename = flag.String("file", "C:\\log.txt", "Full path to a log file")
	flag.Parse()
	f, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}

	j := parseFile(f)

	times := 0
	for _, rec := range j["[ERROR]"] {
		fmt.Println(rec.message)
		fmt.Println("Частота возниктовения ошибки: ", len(rec.lineNumbers))
		fmt.Println("Места возниктовения ошибки: ", rec.lineNumbers)
		times += 1
	}
	fmt.Println("Всего ошибок: ", times)
}

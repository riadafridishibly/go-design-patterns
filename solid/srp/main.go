package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

// Logs keep track of logs
type Logs struct {
	logs []string
}

// Add adds new entry to journal
func (lg *Logs) Add(line string) {
	lg.logs = append(lg.logs, fmt.Sprintf("%s: %s", time.Now().Format(time.UnixDate), line))
}

func (lg *Logs) String() string {
	return strings.Join(lg.logs, "\n")
}

// WriteToDisk should not be implemented like this
func (lg *Logs) WriteToDisk() error {
	panic("Violates SRP")
}

// NewReader Returns a LogReader which implements reader interface
func (lg *Logs) NewReader() *LogReader {
	return &LogReader{
		lines: &lg.logs,
		i:     0,
	}
}

// LogReader reads Logs
type LogReader struct {
	lines *[]string
	i     int
}

func (lr *LogReader) Read(p []byte) (int, error) {
	if lr.lines == nil || lr.i >= (len(*lr.lines)) {
		return 0, io.EOF
	}

	n := copy(p, []byte((*lr.lines)[lr.i]+"\n"))
	lr.i++
	return n, nil
}

func main() {
	lg := Logs{}
	lg.Add("Hello")
	lg.Add("World")

	// Read from string interface
	fmt.Println(lg.String())

	// Read from the reader
	r := lg.NewReader()

	data, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
	// ioutil.WriteFile("somefile.txt", data, 0644)
}

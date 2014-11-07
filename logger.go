package log

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var Stderr io.Writer = os.Stderr
var Discard io.Writer = ioutil.Discard

var topics map[string]struct{} = make(map[string]struct{})

func Println(topic string, v ...interface{}) bool {
	if IsOn(topic) {
		log.Println(v...)
		return true
	}
	return false
}

func SetOutput(w io.Writer) { log.SetOutput(w) }

func On(topic string) {
	for _, t := range strings.Split(topic, " ") {
		topics[t] = struct{}{}
	}
}

func Off(topic string) {
	delete(topics, topic)
}

func IsOn(topic string) bool {
	_, ok := topics[topic]
	return ok
}

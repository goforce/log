package log

import (
	"fmt"
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

func Panic(v ...interface{}) bool {
	s := fmt.Sprint(v...)
	log.Println(s)
	panic(s)
}

func SetOutput(w io.Writer) { log.SetOutput(w) }

func On(topic ...string) {
	for _, ts := range topic {
		for _, t := range strings.Split(ts, " ") {
			if t != "" {
				topics[t] = struct{}{}
			}
		}
	}
}

func Off(topic ...string) {
	for _, ts := range topic {
		for _, t := range strings.Split(ts, " ") {
			if t != "" {
				delete(topics, t)
			}
		}
	}
}

func IsOn(topic string) bool {
	_, ok := topics[topic]
	return ok
}

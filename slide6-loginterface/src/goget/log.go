package goget

import (
	"log"
)

type LogInterface interface{
	Fatal(...interface{})
}

type LogData struct {}

func GetLogInterface() LogInterface {
	return &LogData{}
}

func (_ *LogData) Fatal(i ...interface{}) {
	log.Fatal(i...)
}

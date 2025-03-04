package clog

import "log"

func Error(context string, err error) {
	log.Printf("\033[31m[ERROR]\033[97m %s: %v\n", context, err)
}

func Info(context, msg string) {
	log.Printf("\033[34m[INFO]\033[97m %s: %v\n", context, msg)
}

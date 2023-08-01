package helper

import "log"

func CustomPanic(msg string, err error) {
	log.Panicf("%s: %s\n", msg, err.Error())
}

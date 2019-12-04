package debug

import "log"

var isDebug = false

func EnableDebug() {
	isDebug = true
}

func Printf(format string, args ...interface{}) {
	if isDebug {
		log.Printf(format, args)
	}
}

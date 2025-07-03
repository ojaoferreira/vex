package utils

import "os"

func Info(message string) {
	// Log the info message
	println("[ 💡 INFO ] " + message)
}

func Error(message string) {
	// Log the error message
	println("[ 💥 ERROR ] " + message)
	os.Exit(1)
}

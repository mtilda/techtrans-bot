// Provide tools to communicate with the user via CLI

package main

import (
	"fmt"
	"time"
)

// Print formatted information to the console with a timestamp
func Inform(message string, args ...interface{}) {
	dt := time.Now()

	Printfc("\033[36m%s \033[35m> ", dt.Format(time.RFC1123Z))

	// fg: white, bg: none
	Printfc("\033[0m"+message+"\r\n", args...)
}

// Print a formatted warning to the console with a timestamp
func Warn(message string, args ...interface{}) {
	dt := time.Now()

	Printfc("\033[36m%s \033[35m> ", dt.Format(time.RFC1123Z))

	// fg: bold bright yellow, bg: none
	Printfc("\033[1;33mWARNING: "+message+"\r\n", args...)
}

// Print a formatted error to the console with a timestamp
func Error(err error) {
	dt := time.Now()

	Printfc("\033[36m%s \033[35m> ", dt.Format(time.RFC1123Z))

	// fg: bold bright red, bg: none
	Printfc("\033[1;31mERROR: " + err.Error() + "\r\n")
}

// Print a formatted message and reset color
func Printfc(message string, args ...interface{}) {
	fmt.Printf(message+"\033[0m", args...)
}

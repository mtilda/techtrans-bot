// Provide tools to communicate with the user via CLI

package main

import (
	"fmt"
	"time"
)

// Print formatted information to the console with a timestamp
func Inform(message string, args ...interface{}) {
	dt := time.Now()

	fmt.Printf("\033[36m%s \033[35m> ", dt.Format(time.RFC1123Z))

	// fg: white, bg: none
	fmt.Printf("\033[0m"+message+"\r\n\033[0m", args...)
}

// Print a formatted warning to the console with a timestamp
func Warn(message string, args ...interface{}) {
	dt := time.Now()

	fmt.Printf("\033[36m%s \033[35m> ", dt.Format(time.RFC1123Z))

	// fg: bold bright yellow, bg: none
	fmt.Printf("\033[1;33mWARNING: "+message+"\r\n\033[0m", args...)
}

// Print a formatted error to the console with a timestamp
func Error(message string, args ...interface{}) {
	dt := time.Now()

	fmt.Printf("\033[36m%s \033[35m> ", dt.Format(time.RFC1123Z))

	// fg: bold bright red, bg: none
	fmt.Printf("\033[1;31mERROR: "+message+"\r\n\033[0m", args...)
}

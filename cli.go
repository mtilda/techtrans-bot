// Provide tools to communicate with the user via CLI

package main

import (
	"fmt"
	"time"
)

// Print formatted information to the console with a timestamp
func Inform(message string, args ...interface{}) string {
	// fg: white, bg: none
	return PrintTimestamp("\033[0m"+message, args...)
}

// Print a formatted warning to the console with a timestamp
func Warn(message string, args ...interface{}) string {
	// fg: bold bright yellow, bg: none
	return PrintTimestamp("\033[1;33mWARNING: "+message, args...)
}

// Print a formatted error to the console with a timestamp
func Error(err error) string {
	// fg: bold bright red, bg: none
	return PrintTimestamp("\033[1;31mERROR: " + err.Error())
}

// Print a formatted message with a timestamp
// Reset color and start a new line
func PrintTimestamp(message string, args ...interface{}) string {
	dt := time.Now()
	formattedString := fmt.Sprintf("\033[36m"+dt.Format(time.StampMicro)+" \033[35m~ "+message+"\033[0m\r\n", args...)
	fmt.Print(formattedString)
	return formattedString
}

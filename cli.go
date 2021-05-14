package main

import (
	"fmt"
	"time"
)

func LogInfo(message string, args ...interface{}) {
	dt := time.Now()

	fmt.Printf("\033[36m%s \033[35m> ", dt.Format(time.RFC1123Z))

	fmt.Printf("\033[0m"+message+"\r\n", args...)
}

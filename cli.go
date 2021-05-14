package main

import (
	"fmt"
	"time"
)

func LogInfo(message string) {
	dt := time.Now()
	fmt.Printf("\033[36m%s \033[35m> \033[0m"+message+"\r\n", dt.Format(time.RFC1123Z))
}

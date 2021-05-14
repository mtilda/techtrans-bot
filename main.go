package main

import (
	"fmt"
)

func main() {

	bot := Bot{
		Name:            "techtrans",
		Port:            "8888",
		CredentialsPath: "./credentials.json",
		Server:          "irc.chat.twitch.tv",
	}

	fmt.Println(bot)
}

package main

import (
	"fmt"
)

func main() {

	bot := Bot{
		Name:            "techtrans",
		Port:            "6667",
		CredentialsPath: "./credentials.json",
		Server:          "irc.chat.twitch.tv",
	}

	bot.ReadCredentials()

	fmt.Println(bot.credentials)
}

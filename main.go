package main

func main() {

	bot := Bot{
		Nick:            "techtrans",
		CredentialsPath: "./credentials.json",
		Port:            "6667",
		Server:          "irc.chat.twitch.tv",
		Channel:         "twitch",
	}

	bot.Start()
}

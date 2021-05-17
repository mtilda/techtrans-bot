package main

func main() {

	bot := Bot{
		CredentialsPath: "./credentials.json",
		Port:            "6667",
		Server:          "irc.chat.twitch.tv",
		Channel:         "chillhopmusic",
	}

	bot.Start()
}

package main

func main() {

	bot := Bot{
		Name:            "techtrans",
		CredentialsPath: "./credentials.json",
		Port:            "6667",
		Server:          "irc.chat.twitch.tv",
		Channel:         "twitch",
	}

	bot.ReadCredentials()
	bot.Connect()
	bot.JoinChannel()
}

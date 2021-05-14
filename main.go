package main

func main() {

	bot := Bot{
		Name:            "techtrans",
		CredentialsPath: "./credentials.json",
		Port:            "6667",
		Server:          "irc.chat.twitch.tv",
	}

	bot.ReadCredentials()
	bot.Connect()
}

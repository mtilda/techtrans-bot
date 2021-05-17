// Twitch IRC ChatBot

package main

import (
	"errors"
)

type Bot struct {

	// Domain of the IRC server
	Server string

	// Port of the IRC server
	Port string

	// Name of channel to connect to
	Channel string

	// Path to a json file containing the bot's OAuth credentials
	CredentialsPath string

	// Twitch OAuth credentials
	Credentials OAuthCred

	// IRC server controller
	irc IRC
}

/* Start bot
   Connect to twitch, join channel, and handle chat
*/
func (bot *Bot) Start() {
	err := bot.Credentials.Read(bot.CredentialsPath)
	if nil != err {
		Error(err)
		Inform("Failed to read bot credentials. Aborting.")
		return
	}

	bot.irc.Connect(bot.Server + ":" + bot.Port)
	defer bot.irc.Disconnect(bot.Server + ":" + bot.Port)

	bot.JoinChannel()

	bot.irc.Listen(bot)
	if nil != err {
		Error(err)
	}
}

/* Respond to incoming messages
 */
func (bot *Bot) Chat(content string, sender string, isWhisper bool) error {
	if content == "!tech" {
		tech, err := FetchTech()
		if err != nil {
			return err
		}

		// if isWhisper {
		// 	bot.Whisper(tech, sender)
		// 	return nil
		// }

		bot.Say(tech)
		return nil
	}

	return nil
}

/* Send a message to the chat
 */
func (bot *Bot) Say(message string) error {
	if message == "" {
		return errors.New("Message empty")
	}

	// check if message is too large for IRC
	if len(message) > 512 {
		return errors.New("Message exceeds 512 bytes")
	}

	err := bot.irc.Write("PRIVMSG #" + bot.Channel + " :" + message)
	if nil != err {
		return err
	}

	Inform("(me): " + message)

	return nil
}

/* Send a whisper to a user
   Bot must be verified in order for whispering to be allowed
*/
func (bot *Bot) Whisper(message string, recipient string) error {
	if "" == message {
		return errors.New("Message empty")
	}

	// check if message is too large for IRC
	if len(message) > 512 {
		return errors.New("Message exceeds 512 bytes")
	}

	err := bot.irc.Write("PRIVMSG #" + bot.Channel + " :/w " + recipient + " " + message)
	if nil != err {
		return err
	}

	Inform("(I) \033[48:5:92;37m*whisper to (" + recipient + ")\033[0m: " + message)

	return nil
}

/* Join a channel
   Will panic if connection is not established
*/
func (bot *Bot) JoinChannel() (err error) {
	Inform("Joining #%s as @%s...", bot.Channel, bot.Credentials.Nick)

	err = bot.irc.Write("PASS oauth:" + bot.Credentials.Password + "\r\n")
	err = bot.irc.Write("NICK " + bot.Credentials.Nick + "\r\n")
	err = bot.irc.Write("CAP REQ :twitch.tv/commands\r\n") // enable reading whispers
	err = bot.irc.Write("JOIN #" + bot.Channel + "\r\n")

	return
}

// Provide tools for connecting to and interacting with Twitch IRC

package main

import (
	"bufio"
	"errors"
	"net"
	"net/textproto"
	"strings"
	"time"
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

	// Reference to the IRC connection
	connection net.Conn
}

// Start bot
// Connect to twitch, join channel, and handle chat
func (bot *Bot) Start() {
	err := bot.Credentials.Read(bot.CredentialsPath)
	if nil != err {
		Error(err)
		Inform("Failed to read bot credentials. Aborting.")
		return
	}

	bot.Connect()
	defer bot.Disconnect()

	bot.JoinChannel()

	bot.Monitor()
	if nil != err {
		Error(err)
	}
}

// Monitor the current IRC connection forever
func (bot *Bot) Monitor() error {
	tp := textproto.NewReader(bufio.NewReader(bot.connection))
	for {
		line, err := tp.ReadLine()
		if err != nil {
			Error(err)
			return errors.New("Failed to read buffer")
		}

		formattedLine, err := bot.HandleIRCLine(line)
		if err != nil {
			Error(err)
			return errors.New("Unable to parse IRC line: '" + line + "'")
		}
		if formattedLine != "" {
			Inform(formattedLine)
		}
	}
}

// Parse a single IRC line
// Check for commands
// Format messages
func (bot *Bot) HandleIRCLine(line string) (string, error) {
	var formattedLine string

	if strings.Contains(line, "PING :tmi.twitch.tv") {
		_, err := bot.connection.Write([]byte("PONG :tmi.twitch.tv\r\n"))
		if nil != err {
			Error(err)
			return "", errors.New("Unable to send 'PONG :tmi.twitch.tv'")
		}

		return "ping...pong", nil
	}

	lineSlice := strings.Split(line, " :")

	if len(lineSlice) >= 1 {
		user := strings.Split(lineSlice[0], " ")[0]
		if strings.HasPrefix(user, ":") {
			user = strings.Split(user, ":")[1]
			user = strings.Split(user, "!")[0]
			formattedLine += user + " : "
		}
	}
	if len(lineSlice) >= 2 {
		formattedLine += strings.Join(lineSlice[1:], " :")
	}

	return formattedLine, nil
}

// Make the bot send a message to the chat channel
func (bot *Bot) Say(message string) error {
	if message == "" {
		return errors.New("Message empty")
	}

	// check if message is too large for IRC
	if len(message) > 512 {
		return errors.New("Message exceeds 512 bytes")
	}

	_, err := bot.connection.Write([]byte("PRIVMSG #" + bot.Channel + " :" + message + "\r\n"))
	if nil != err {
		return err
	}

	Inform(bot.Credentials.Nick + " : " + message)

	return nil
}

// Make the bot send a whisper to a user
// Bot must be verified in order for whispering to be allowed
func (bot *Bot) Whisper(message string, recipient string) error {
	if "" == message {
		return errors.New("Message empty")
	}

	// check if message is too large for IRC
	if len(message) > 512 {
		return errors.New("Message exceeds 512 bytes")
	}

	_, err := bot.connection.Write([]byte("PRIVMSG #" + bot.Channel + " :/w " + recipient + " " + message + "\r\n"))
	if nil != err {
		return err
	}

	Inform(bot.Credentials.Nick + " \033[35m<whisper> " + recipient + " : " + message)

	return nil
}

// Connect the bot to the Twitch IRC server
// Retry until successful, with exponential backoff
func (bot *Bot) Connect() {
	Inform("Attempting to connect with %s...", bot.Server)

	delay := 1
	for {
		// Make connection to Twitch IRC server
		var err error
		bot.connection, err = net.Dial("tcp", bot.Server+":"+bot.Port)
		if err != nil {
			Error(err)
			Inform("Connection to %s failed, retrying in %d seconds...", bot.Server, delay)
			time.Sleep(time.Duration(delay) * time.Second)
			delay *= 2
			continue
		}

		break
	}

	Inform("Connected to %s!", bot.Server)
}

// Disconnect bot from the Twitch IRC server
func (bot *Bot) Disconnect() {
	Inform("Disconnecting from IRC server.")
	bot.connection.Close()
	Inform("Closed connection with %s!", bot.Server)
}

// Make the bot join its pre-specified channel
// Will panic if connection is not established
func (bot *Bot) JoinChannel() {
	Inform("Attempting to join #%s...", bot.Channel)

	bot.connection.Write([]byte("PASS oauth:" + bot.Credentials.Password + "\r\n"))
	bot.connection.Write([]byte("NICK " + bot.Credentials.Nick + "\r\n"))
	bot.connection.Write([]byte("CAP REQ :twitch.tv/commands\r\n")) // enable reading whispers
	bot.connection.Write([]byte("JOIN #" + bot.Channel + "\r\n"))

	Inform("Joined #%s as @%s!", bot.Channel, bot.Credentials.Nick)
}

// Provide tools for connecting to and interacting with Twitch IRC

package main

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/textproto"
	"os"
	"time"
)

type OAuthCredentials struct {
	// The bot account's OAuth password
	Password string `json:"password"`

	// The developer application client ID
	// Used for API calls to Twitch
	ClientID string `json:"client_id"`
}

type Bot struct {
	// Name the bot goes by
	Name string

	// Path to a json file containing the bot's OAuth credentials
	CredentialsPath string

	// Port of the IRC server
	Port string

	// Domain of the IRC server
	Server string

	// Name of channel to connect to
	Channel string

	/* Private */

	// Twitch OAuth credentials
	credentials *OAuthCredentials

	// Reference to the IRC connection
	connection net.Conn
}

// Start bot
// Connect to twitch, join channel, and handle chat
func (bot *Bot) Start() {
	err := bot.ReadCredentials()
	if nil != err {
		Error(err)
		Inform("Failed to read bot credentials. Aborting.")
		return
	}

	bot.Connect()
	bot.JoinChannel()

	tp := textproto.NewReader(bufio.NewReader(bot.connection))
	for {
		line, err := tp.ReadLine()
		if nil != err {
			Error(err)
			Inform("Failed to read buffer. Disconnecting from IRC server.")
			bot.Disconnect()
			return
		}

		Inform(line)
	}
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
	bot.connection.Close()
	Inform("Closed connection with %s!", bot.Server)
}

// Make the bot join its pre-specified channel
// Will panic if connection is not established
func (bot *Bot) JoinChannel() {
	Inform("Attempting to join #%s...", bot.Channel)

	bot.connection.Write([]byte("NICK " + bot.Name + "\r\n"))
	bot.connection.Write([]byte("PASS " + bot.credentials.Password + "\r\n"))
	bot.connection.Write([]byte("JOIN #" + bot.Channel + "\r\n"))

	Inform("Joined #%s as @%s!", bot.Channel, bot.Name)
}

// Read from the private credentials json file
// Stores the data in the bot's Credentials field
func (bot *Bot) ReadCredentials() error {
	bot.credentials = &OAuthCredentials{}

	credJSON, err := os.Open(bot.CredentialsPath)
	if nil != err {
		return err
	}

	credByte, err := ioutil.ReadAll(credJSON)
	if nil != err {
		return err
	}

	err = json.Unmarshal(credByte, &bot.credentials)
	if nil != err {
		return err
	}

	return nil
}

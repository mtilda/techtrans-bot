package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

type OAuthCred struct {
	// The bot account's OAuth password
	Password string `json:"password"`

	// The developer application client ID
	// Used for API calls to Twitch
	ClientID string `json:"client_id"`
}

type Bot struct {
	// Name the bot goes by
	Name string

	// Port of the IRC server
	Port string

	// Path to a json file containing the bot's OAuth credentials
	CredentialsPath string

	// The domain of the IRC server
	Server string

	/* Private */

	// Twitch OAuth credentials
	credentials *OAuthCred

	// Reference to the IRC connection
	connection net.Conn
}

// Connects the bot to the Twitch IRC server
// Retry until successful
func (bot *Bot) Connect() {
	fmt.Printf("Connecting to %s...\n", bot.Server)

	for {
		// Make connection to Twitch IRC server
		var err error
		bot.connection, err = net.Dial("tcp", bot.Server+":"+bot.Port)
		if err == nil {
			break
		}
	}

	fmt.Printf("Connected to %s!\n", bot.Server)
}

// Reads from the private credentials file and stores the data in the bot's Credentials field
func (bot *Bot) ReadCredentials() error {

	credJSON, err := os.Open(bot.CredentialsPath)
	if nil != err {
		return err
	}

	bot.credentials = &OAuthCred{}

	credByte, err := ioutil.ReadAll(credJSON)
	if nil != err {
		return err
	}

	// Unmarshal credByte into bot.Credentials
	// Case and underscore insensitive
	json.Unmarshal(credByte, &bot.credentials)

	return nil
}

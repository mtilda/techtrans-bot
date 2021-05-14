package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
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

	// The domain of the IRC server
	Server string

	/* Private */

	// Twitch OAuth credentials
	credentials *OAuthCredentials

	// Reference to the IRC connection
	connection net.Conn
}

// Connect the bot to the Twitch IRC server
// Retry until successful, with exponential backoff
func (bot *Bot) Connect() {
	fmt.Printf("Connecting to %s...\n", bot.Server)

	delay := 1
	for {
		// Make connection to Twitch IRC server
		var err error
		bot.connection, err = net.Dial("tcp", bot.Server+":"+bot.Port)
		if err == nil {
			break
		}

		fmt.Printf("Failed to connect to %s, retrying in %d seconds...\n", bot.Server, delay)
		time.Sleep(time.Duration(delay) * time.Second)
		delay *= 2
	}

	fmt.Printf("Connected to %s!\n", bot.Server)
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

	json.Unmarshal(credByte, &bot.credentials)

	return nil
}

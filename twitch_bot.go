package main

import (
	"net"
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

// Tools for interacting with the Twitch IRC server

package main

import (
	"bufio"
	"errors"
	"net"
	"net/textproto"
	"strings"
	"time"
)

type IRC struct {
	// Address of IRC server
	Address string

	// Reference to server connection
	conn net.Conn
}

type Chatter interface {
	Chat(content string, sender string, isWhisper bool) error
}

// Listen to current IRC connection forever
func (irc *IRC) Listen(chatter Chatter) error {
	tp := textproto.NewReader(bufio.NewReader(irc.conn))
	for {
		line, err := tp.ReadLine()
		if err != nil {
			Error(err)
			return errors.New("Failed to read buffer")
		}

		err = irc.HandleLine(line, chatter)
		if err != nil {
			Error(err)
			return errors.New("Unable to parse IRC line: '" + line + "'")
		}
	}
}

// Parse a single IRC line
// Check for commands
// Invoke chatter.Chat() on messages
func (irc *IRC) HandleLine(line string, chatter Chatter) error {

	// Avoid premature disconnections by playing ping-pong with server
	if strings.Contains(line, "PING :tmi.twitch.tv") {
		err := irc.Write("PONG :tmi.twitch.tv")
		if nil != err {
			Error(err)
			return errors.New("Unable to send 'PONG :tmi.twitch.tv'")
		}

		Inform("ping...pong")
	}

	/* line is probably a message! */

	// content comes after " :"
	// prefix comes before " :"
	ll := strings.Split(line, " :")

	// if message has any content
	if len(ll) >= 2 {
		prefix, content, sender, isWhisper := ll[0], ll[1], "", false

		// extract sender
		sender = strings.Split(prefix, " ")[0]
		if strings.HasPrefix(sender, ":") {
			sender = strings.Split(sender, ":")[1]
			sender = strings.Split(sender, "!")[0]
		}

		// check if message is a whisper
		isWhisper = strings.Contains(prefix, "WHISPER")

		if isWhisper {
			Inform("(" + sender + ") \033[48:5:92;37m*whispers*\033[0m: " + content)
		} else {
			Inform("(" + sender + "): " + content)
		}

		// let the bot decide what it wants to do with this
		err := chatter.Chat(content, sender, isWhisper)
		if err != nil {
			Error(err)
		}
	}

	return nil
}

/* Write a line to the IRC server
 */
func (irc *IRC) Write(line string) (err error) {
	_, err = irc.conn.Write([]byte(line))
	if err != nil {
		Error(err)
		err = errors.New("Could not write line to IRC server: \"" + line + "\"")
		return
	}

	return
}

/* Initiate connection to IRC server
   Retry until successful, with exponential backoff
*/
func (irc *IRC) Connect(address string) {
	Inform("Attempting to connect with %s...", address)

	delay := 1
	for {
		// Make connection to Twitch IRC server
		var err error
		irc.conn, err = net.Dial("tcp", address)
		if err != nil {
			Error(err)
			Inform("Connection to %s failed, retrying in %d seconds...", address, delay)
			time.Sleep(time.Duration(delay) * time.Second)
			delay *= 2
			continue
		}

		break
	}

	Inform("Connected to %s!", address)
}

// Disconnect from IRC server
func (irc *IRC) Disconnect(address string) {
	Inform("Disconnecting from IRC server.")
	irc.conn.Close()
	Inform("Closed connection with %s!", address)
}

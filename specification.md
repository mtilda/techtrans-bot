# Interview Assignment - Twitch Bot #

## Overview ##
---
Create an automated [Twitch](https://dev.twitch.tv/docs/irc) chat bot console application that can be run from a command line interface (CLI).


## Requirements
---
The bot application should be able to:
* Console output all interactions - legibly formatted, with timestamps.
* Connect to Twitch IRC over SSL.
* Join a channel.
* Read a channel.
* Read a private message.
* Write to a channel
* Reply to a private message.
* Avoid premature disconnections by handling Twitch courier ping / pong requests.
* Publicly reply to a user-issued string command within a channel (!YOUR_COMMAND_NAME).
* Reply to the "!tech" command by dynamically returning details related to a randomly chosen tech from the [NASA](https://api.nasa.gov/) [TechTransfer API](https://api.nasa.gov/techtransfer/patent/?engine&api_key=DEMO_KEY). If a tech description includes HTML elements, insure that they have been removed before returning output to a chat channel.


## Caveats ##
---
* The application must be written in Go using the [standard library](https://golang.org/pkg/) - absolutely no third-party module dependencies.
* All interactions should be asynchronous.
* The application should account for Twitch API rate limits.
* The application should not exit prematurely.

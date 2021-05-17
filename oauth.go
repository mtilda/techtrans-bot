package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Twitch IRC OAuth Credentials
type OAuthCred struct {
	Nick     string `json:"nick"`
	Password string `json:"password"`
}

/* Read secrete credentials json file
   Unmarshal values into destinationCred
*/
func (destinationCred *OAuthCred) Read(sourcePath string) error {
	credJSON, err := os.Open(sourcePath)
	if nil != err {
		return err
	}

	credByte, err := ioutil.ReadAll(credJSON)
	if nil != err {
		return err
	}

	err = json.Unmarshal(credByte, &destinationCred)
	if nil != err {
		return err
	}

	return nil
}

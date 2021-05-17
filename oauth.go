package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type OAuthCred struct {
	Nick     string `json:"nick"`
	Password string `json:"password"`
}

func (cred *OAuthCred) Read(sourcePath string) error {
	credJSON, err := os.Open(sourcePath)
	if nil != err {
		return err
	}

	credByte, err := ioutil.ReadAll(credJSON)
	if nil != err {
		return err
	}

	err = json.Unmarshal(credByte, &cred)
	if nil != err {
		return err
	}

	return nil
}

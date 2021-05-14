# TechTrans

## Overview
A Twitch bot that tells you about cool inventions by pulling data from NASA's Technology Transfer API

---

## Instructions
---
### Run
1. Change into the project directory
    ```
    cd techtrans-bot
    ```
1. Build the project
    ```
    go build
    ```
1. Run the executable binary
    ```
    ./techtrans-bot
    ```
---
### [Credentials](https://dev.twitch.tv/console/apps)

Create a `credentials.json` file in the root with the following structure. Fill in the values with the appropriate data. Password should contain the OAuth access token.
```
{
  "password": "",
  "client_id": ""
}
```

Obtain a new OAuth access token by sending the following HTTP request.
```
method: POST
url: 	https://id.twitch.tv/oauth2/token
        ?client_id=<your client ID>
        &client_secret=<your client secret>
        &grant_type=client_credentials
        &scope=<space-separated list of scopes>
```

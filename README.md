# TechTrans

## Overview
A Twitch bot that tells you about cool inventions by pulling data from NASA's Technology Transfer API

---

## `!tech` command
Returns the name and image URL of a NASA patent.

Example output:

    Aircraft Engine Icing Event Avoidance and Mitigation
    https://tinyurl.com/yf4oa2ab

---

![Computer generated plane flying through a cloudy sky](https://ntts-prod.s3.amazonaws.com/t2p/prod/t2media/tops/img/LEW-TOPS-125/iStock-157730835_LEW-19309-1_airplane-storm_1388x1050-300dpi.jpg)

## Developer Instructions
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

Create a `credentials.json` file in the root with the following structure. Fill in the values with the appropriate data. Password should contain the OAuth access token, which you can obtain here: https://twitchapps.com/tmi/
```
{
  "password": "",
  "client_id": ""
}
```
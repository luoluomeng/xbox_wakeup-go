
## xbox_wakeup-go 


[![Build Status](https://travis-ci.org/luoluomeng/xbox_wakeup-go.svg?branch=master)](https://travis-ci.org/luoluomeng/xbox_wakeup-go)
[![GitHub release](http://img.shields.io/github/release/luoluomeng/xbox_wakeup-go.svg?style=flat-square)](https://github.com/luoluomeng/xbox_wakeup-go/releases)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](https://github.com/luoluomeng/xbox_wakeup-go/blob/master/LICENSE)


This tool is a golang port from xbox-remote-power(https://github.com/Schamper/xbox-remote-power)

You need to forward 5050 port to your XBOX ONE on you router if you need to wake it on WAN

If you need to stream your XBOX ONE over internet, you need also forward port 4838,49000-65000 to your XBOX ONE

Or set up a vpn connection to your XBOX's network


## How to use
```
xbox_wakeup-go.exe <ip address> <live id>
```

If you are waking your XBOX through WAN, "ip address" is your external IP

If you are waking your XBOX through Lan and VPN, "ip address" is the XBOX's LAN IP

Your Live device ID can be found at Settings -> System -> Console info on your XBOX

NOTE: Please don't share this ID to anyone



xbox_wakeup-go [![Build Status](https://travis-ci.org/luoluomeng/xbox_wakeup-go.svg?branch=master)](https://travis-ci.org/luoluomeng/xbox_wakeup-go)

======================

This tool is a golang port from xbox-remote-power(https://github.com/Schamper/xbox-remote-power)

You need to forward 5050 port to your XBOX one on you router if you need to wake it on WAN

Or set up a vpn to your home network


## How to use
```
xbwake <ip address> <live id>
```

If you are waking your XBOX through WAN, the "ip address" is your external IP

If you are waking your XBOX through Lan and VPN, the "ip address" is the XBOX's LAN IP

Your Live device ID can be found at Settings -> System -> Console info on your XBOX

NOTE: Please don't share this information to anyone


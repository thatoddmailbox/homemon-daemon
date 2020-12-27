# homemon-daemon
This program runs on a device and sends reports to a copy of [homemon-server](https://github.com/thatoddmailbox/homemon-server) or [homemon-receiver](https://github.com/thatoddmailbox/homemon-receiver).

## Development
To build the code and run it on a connected MW41:

1. `GOOS=linux GOARCH=arm go build`
2. `adb push homemon-daemon /media/card`
3. `adb shell "cd /media/card && ./homemon-daemon"`

## Configuration
Create a `config.toml` file, in the same directory you run the daemon from. This file should look something like:
```
Interval = "10s"
InitialDelay = "5s"
Host = "1.2.3.4"
Port = 7890
Token = "token goes here"
Transport = "HTTP"
```

### Options
* Interval - the time to wait between reports
* InitialDelay - the time to wait before sending the first report
* Host - the host to send reports to
* Port - the port on the host to send reports to (only applies to UDP transport)
* Token - the token to use for authencation, must match the token in homemon-server or homemon-receiver
* Transport - the transport to use (either HTTP or UDP)
# homemon-daemon
This program runs on a device and sends regular status reports to a copy of [homemon-server](https://github.com/thatoddmailbox/homemon-server) or [homemon-receiver](https://github.com/thatoddmailbox/homemon-receiver). It is intended to be used with an Alcatel MW41 mobile hotspot, which you need to first [get root on](https://alex.studer.dev/2021/01/04/mw41-1).

## Development
To build the code and run it on a connected MW41:

1. `GOOS=linux GOARCH=arm go build`
2. `adb push homemon-daemon /media/card`
3. `adb shell "cd /media/card && ./homemon-daemon"`

## Configuration
Create a `config.toml` file in the same directory you run the daemon from. This file should look something like:
```toml
Interval = "10m"
InitialDelay = "1m"
RestartCount = 144
Host = "1.2.3.4"
Port = 7890
Token = "token goes here"
Transport = "HTTP"
```

### Options
* `Interval` - the time to wait between reports
* `InitialDelay` - the time to wait before sending the first report (this is useful if you want to wait for, say, the network to come up after your device has restarted)
* `RestartCount` - the number of reports between reboots of the device (can be set to 0 to disable rebooting)
* `Host` - the host to send reports to (if using the UDP transport, it's recommended to make this an IP address so you don't have to do a DNS lookup, but a hostname will work)
* `Port` - the port on the host to send reports to (only applies to UDP transport)
* `Token` - the token to use for authentication, must match the token in homemon-server or homemon-receiver
* `Transport` - the transport to use (either HTTP or UDP)

Durations (like `Interval` and `InitialDelay`) must be in a format supported by Go's [time.ParseDuration](https://golang.org/pkg/time/#ParseDuration).
# homemon-daemon
This program runs on a device and sends reports to a copy of [homemon-server](https://github.com/thatoddmailbox/homemon-server) or [homemon-receiver](https://github.com/thatoddmailbox/homemon-receiver).

## Development
To build the code and run it on a connected MW41:

1. `GOOS=linux GOARCH=arm go build`
2. `adb push homemon-daemon /media/card`
3. `adb shell "cd /media/card && ./homemon-daemon"`
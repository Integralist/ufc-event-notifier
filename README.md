<h1 align="center">UFC Event Notifier</h1>

<p align="center">
  <img src="https://img.shields.io/badge/Completed-100%25-green.svg?style=flat-square">
</p>

<p align="center">
  Long running binary (written in Go) for Mac OS X that scrapes the UFC website for the next main UFC event and triggers a native OS notification
</p>

## Build

```bash
GOOS=darwin GOARCH=386   go build -o ufc-386   ufc.go # 32 bit MacOSX
GOOS=darwin GOARCH=amd64 go build -o ufc-amd64 ufc.go # 64 bit MacOSX
```

## Execute

```
mv ./ufc /usr/local/bin # update $PATH reference
which ufc               # /usr/local/bin/ufc
ufc &                   # run program in a background process
```

## Launch

You can utilise the Mac OS X Launchd service, which controls which programs are launched when the OS boots.

> Reference: [developer.apple.com/launchd.plist](https://developer.apple.com/library/mac/documentation/Darwin/Reference/ManPages/man5/launchd.plist.5.html)

Modify the `ufc-notifier.plist` file inside this repo to include details specific for your needs and move it to the following location:

```
mv ufc-notifier.plist $HOME/Library/LaunchAgents
```

Then execute:

```bash
cd "$HOME/Library/LaunchAgents"
launchctl load ufc-notifier.plist
```

> Note: this has been tested with Mac OS X version: `10.11.1`

To stop this service:

```bash
cd "$HOME/Library/LaunchAgents"
launchctl unload ufc-notifier.plist
rm -i ufc-notifier.plist
```

### Let's get serious

If you find the service isn't stopping, then it's time to get heavy handed...

- Check the program is indeed handled by Launchd: `launchctl list | grep ufc`
- Remove the program: `launchctl remove ufc-notifier`
- Verify it's gone: `launchctl list | grep ufc`

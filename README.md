<h1 align="center">UFC Event Notifier</h1>

<p align="center">
  <img src="https://img.shields.io/badge/Completed-100%25-green.svg?style=flat-square">
</p>

<p align="center">
  Long running binary (written in Go) for Mac OS X that scraps the UFC website for the next main UFC event and triggers a native OS notification
</p>

## Build

```bash
GOOS=darwin GOARCH=386   go build -o ufc-386   ufc.go # 32 bit MacOSX
GOOS=darwin GOARCH=amd64 go build -o ufc-amd64 ufc.go # 64 bit MacOSX
```

## Execute

```
mv ./ufc /usr/local/bin # update $PATH reference
which ufc # /usr/local/bin/ufc
ufc
```

## Launch

Modify the `ufc-notifier.plist` file inside this repo and move it to:

```
/Library/LaunchAgents/ufc-notifier.plist
```

Then execute:

```bash
sudo launchctl load /Library/LaunchAgents/ufc-notifier.plist
```

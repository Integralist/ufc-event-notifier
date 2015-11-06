<h1 align="center">UFC Event Notifier</h1>

<p align="center">
  <img src="https://img.shields.io/badge/Completed-100%25-green.svg?style=flat-square">
</p>

<p align="center">
  Long running binary (written in Go) for Mac OS X that scraps the UFC website for the next main UFC event and triggers a native OS notification
</p>

## Build

```bash
GOOS=darwin GOARCH=386   go build ufc.go # 32 bit MacOSX
GOOS=darwin GOARCH=amd64 go build ufc.go # 64 bit MacOSX
```

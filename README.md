# share.me
share.me is a basic file sharing tool

You can install share.me via go directly
```bash
$ go install github.com/pale-whale/share.me@latest
```

## Basic usage
There is only one command at the moment and it's the share command 
```bash
$ share.me share --help
Start a server to share a single file (or directory)
or start an UI that let you select which files to share

Usage:
  share.me share [flags] [file]

Flags:
  -h, --help   help for share

Global Flags:
      --config string   config file (default is $HOME/.share.me.yaml)
  -p, --port string     port used for the server (default 0)
```

You can invoke share with an argument to share it directly or without to start the sharing server also, invoking share.me without command is a shortcut to the share command

## Sharing a single file
To share a single file invoke share.me with the path to said file 
```bash
$ share.me example.txt 
addr: http://192.168.1.31:33297/
file: example.txt
```
share.me will start to listen to a random port and serve the file to anyone requesting /
## Sharing Server
share.me can also be used as a sharing server, invoke it without a file and an ui will be started.
From here you can browse your current directory and share any files here

## Config
share.me can be configurated, it will look for a config in `$HOME/.share.me.yaml` keep in mind that this location WILL change.
This file is a yaml file where you can set your default options.
For example setting a port here
```yaml
port: 8080
```
will have for result that share.me will try to bind on port 8080 on startup
```bash
$ share.me share.me 
Using config file: /home/palewhale/.share.me.yaml
addr: http://192.168.1.31:8080/
file: share.me
```

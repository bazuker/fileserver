[![Go Report Card](https://goreportcard.com/badge/github.com/kisulken/fileserver)](https://goreportcard.com/report/github.com/kisulken/fileserver)

# Go fileserver
Simple file server that shares access to a directory over HTTP using a selected port.

By default, the current working directory is hosted on port `8080` if no flags were specified.

## Installation
Build from the source
```bash
$ go install github.com/bazuker/fileserver
```
Or check out the [pre-built binaries](https://github.com/bazuker/fileserver/releases)

## CLI Usage

```bash
$ fileserver -port=5555 -dir="/users/username/documents/public"
```

This command will host the directory `/users/username/documents/public` on the port 5555.
Now, you should be able to access it via a web browser at http://localhost:5555/

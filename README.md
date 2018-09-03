# Simple file server written in Golang

## Installation
```bash
go get github.com/kisulken/fileserver
```

Simple file server that shares access to a directory on a selected port

## CLI Usage

```bash
fileserver 5555 "/users/username/documents/public"
```

This command will share the access of /users/username/documents/public on the port 5555.
Now, you should be able to access that directory using a web browser at http://localhost:5555/

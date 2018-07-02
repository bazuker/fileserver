# Simple file server written in Golang
Simple file server that shares access to a directory on a selected port

Example of use:

"/../HttpFileServer.exe" 5555 "C:/folder/data"

This command will host the folder C:/folder/data using the port 5555.
Now, you should be able to access that directory using a web browser at http://localhost:5555/

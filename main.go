package main

import (
	"net/http"
	"fmt"
	"strconv"
	"os"
	"github.com/kardianos/osext"
)

func main() {
	fmt.Println("Simple file server v1.0")
	// getting the port from the user
	fmt.Print("Enter port: ")

	var port string
	fmt.Scanln(&port)

	// validating the port
	_, err := strconv.Atoi(port)
	if err != nil {
		fmt.Println("Invalid port" + err.Error())
		os.Exit(1)
	}

	// getting the executable path
	wd, _ := osext.ExecutableFolder()

	fmt.Println("Working directory is " + wd)

	// setting up the file server
	fs := http.FileServer(http.Dir(wd))
	http.Handle("/", fs)

	// running the file server
	fmt.Println("Serving the files...")
	http.ListenAndServe(":" + port, nil)
}
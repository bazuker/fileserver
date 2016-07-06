package main

import (
	"net/http"
	"fmt"
	"strconv"
	"os"
	"github.com/kardianos/osext"
)

func main() {
	// setting the default settings
	port := "1234"
	wd, _ := osext.ExecutableFolder()
	customSettings := false

	// checking for custom settings
	if len(os.Args) > 1 {
		customSettings = true

		port = os.Args[1]

		if len(os.Args) > 2 {
			wd = os.Args[2]
		}
	}

	// validating the settings
	if customSettings {
		if _, err := strconv.Atoi(port); err != nil {
			fmt.Println("Invalid port!")
			os.Exit(1)
		}
		if _, err := os.Stat(wd); os.IsNotExist(err) {
			fmt.Println("Invalid working directory!")
			os.Exit(2)
		}
	}

	fmt.Println("Port is " + port)
	fmt.Println("Working directory is " + wd)

	// setting up the file server
	fs := http.FileServer(http.Dir(wd))
	http.Handle("/", fs)

	// running the file server
	fmt.Println("Serving the files...")
	http.ListenAndServe(":" + port, nil)
}
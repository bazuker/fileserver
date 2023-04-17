package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

var (
	port = flag.Int("port", 1234, "port to use")
	// By default, the current directory will be used
	wd = flag.String("dir", "", "directory to serve")
)

func main() {
	flag.Parse()

	// setting the default settings
	if wd == nil || *wd == "" {
		detectedWD, err := os.Getwd()
		if err != nil {
			fmt.Println("failed to determine working directory:", err)
			return
		}
		wd = &detectedWD
	}
	if port == nil || *port == 0 {
		fmt.Printf("Invalid port provided")
		return
	}

	if _, err := os.Stat(*wd); os.IsNotExist(err) {
		fmt.Println("Invalid working directory", *wd)
		os.Exit(2)
	}

	fmt.Println("Using port", *port)
	fmt.Println("Using directory", *wd)

	// setting up the file server
	fs := http.FileServer(http.Dir(*wd))
	http.Handle("/", fs)

	// running the file server
	fmt.Println("Serving the files...")
	if err := http.ListenAndServe(":"+strconv.Itoa(*port), nil); err != nil {
		fmt.Println("Error:", err)
	}
}

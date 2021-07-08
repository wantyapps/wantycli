package main

import (
	"fmt"
	"flag"
	"os"
	"io"
	"io/ioutil"
	"net/http"
)

func usage() {
	fmt.Printf("Usage:\n\n")
	fmt.Printf("API\n")
	fmt.Printf("    api --url {string} --username {string} --password {string} | (optional) --action {string}\n")
}

func main() {
	// Set up command-line arguments
	apiCmd := flag.NewFlagSet("api", flag.ExitOnError)
	apiURL := apiCmd.String("url", "", "")
	apiUsername := apiCmd.String("username", "", "")
	apiPassword := apiCmd.String("password", "", "")
	apiAction := apiCmd.String("action", "", "")
	apiCmd.Usage = usage
	apiCmd.SetOutput(io.Discard)
	// client := &http.Client{}
	// req, err := http.NewRequewt("GET", *apiRequestTool, nil)
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	switch os.Args[1] {
		case "api":
			apiCmd.Parse(os.Args[2:])
			if *apiURL != "" {
				client := &http.Client{}
				req, _ := http.NewRequest("GET", *apiURL, nil)
				req.Header.Set("username", *apiUsername)
				req.Header.Set("password", *apiPassword)
				if *apiAction != "" {
					req.Header.Set("action", *apiAction)
				}
				
				res, err := client.Do(req)
				if err != nil {
					fmt.Println("[\033[91mERROR\033[0m] Connection error.")
					os.Exit(1)
				}
				defer res.Body.Close()
				body, err := ioutil.ReadAll(res.Body)
				if err != nil {
					fmt.Println("[\033[91mERROR\033[0m] Error reading response body.")
				}

				fmt.Println(string(body))
			} else {
				usage()
				os.Exit(1)
			}
			
		default:
			usage()
			os.Exit(1)
	}
}

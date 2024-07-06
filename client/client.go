package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	endpoint := flag.String("e", "http://localhost:8080/items", "The endpoint to request.")
	method := flag.String("m", "GET", "The HTTP method to use.")

	flag.Parse()

	if flag.NArg() == 0 {

		switch *method {
		case "GET":
			fmt.Println("GET")
			response, err := http.Get(*endpoint)

			if err != nil {
				fmt.Print(err.Error())
				os.Exit(1)
			}

			responseData, err := io.ReadAll(response.Body)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(string(responseData))
		case "POST":
			fmt.Println("POST")
			response, err := http.Post(*endpoint, "application/json", nil)

			if err != nil {
				fmt.Print(err.Error())
				os.Exit(1)
			}

			responseData, err := io.ReadAll(response.Body)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(string(responseData))
		case "PUT":
			fmt.Println("PUT")
		}

	} else if flag.Arg(0) == "list" {
		files, _ := os.Open(".")
		defer files.Close()

		fileInfo, _ := files.Readdir(-1)
		for _, file := range fileInfo {
			fmt.Println(file.Name())
		}
	} else {
		fmt.Printf("Hello, %s!\n", *endpoint)
		fmt.Printf("Hello, %s!\n", *method)
		fmt.Println("Helloasd")
	}

}

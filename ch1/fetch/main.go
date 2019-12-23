package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	for _, url := range os.Args[1:] {

		resp, err := http.Get(url)
		handleError(err)

		body, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		handleError(err)
		fmt.Printf("%s", body)

	}
}

func handleError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
}

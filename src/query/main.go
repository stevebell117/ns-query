package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	response := Query(os.Getenv("NIGHTSCOUT_ENDPOINT")))

	fmt.Println(response)
}

/* Send a Query that then returns the result. */
func Query(queryStr string) string {
	response, err := http.Get(queryStr)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(responseData)
}

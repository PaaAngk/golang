package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	client := http.Client{}

	resp, err := client.Get("http://localhost/1")
	if err != nil {
		log.Fatal(err)
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(buf))

}

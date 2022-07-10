package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {

	client := http.Client{}

	//var timeArray []time.Duration

	// resp, err := client.Get("http://localhost/1")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// buf, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(string(buf))

	for i := 0; i < 2; i++ {
		cPost(client, "async")
	}
	cPost(client, "sync")

}

func cPost(client http.Client, qType string) {
	var _, errPost = client.PostForm("http://localhost/add", url.Values{"time": {"0h0m10s"}, "type": {qType}})

	if errPost != nil {
		log.Fatal(errPost)
	}
	//fmt.Println("respPost", respPost)
	fmt.Println("Query type: ", qType)
}

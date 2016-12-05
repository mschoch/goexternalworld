package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// START OMIT
	resp, err := http.Get("http://localhost:8081/") // HL
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // HL
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", body)
	// END OMIT
}

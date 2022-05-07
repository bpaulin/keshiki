package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/go-resty/resty/v2"

	"gopkg.in/yaml.v3"
)

func main() {
	// Yaml load
	yfile, err := ioutil.ReadFile("config.yaml")

	if err != nil {
		log.Fatal(err)
	}

	data := make(map[interface{}]interface{})

	err2 := yaml.Unmarshal(yfile, &data)

	if err2 != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", data)

	// rest client
	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		Get("https://bitbucket.org/")

	// Explore response object
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println()
}

package services

import (
	"log"
	"net/http"

	"github.com/nautilus/gateway"
	"github.com/nautilus/graphql"
)

const BookApiURL = "http://user-api:1112/api"
const HttpPort = "1111"

func Start() {
	schemas, err := graphql.IntrospectRemoteSchemas(BookApiURL)
	if err != nil {
		log.Fatalf("Error in IntrospectRemoteSchemas: %v", err)
	}

	gw, err := gateway.New(schemas)
	if err != nil {
		log.Fatalf("Error when creating gateway: %v", err)
	}
	http.HandleFunc("/api", gw.GraphQLHandler)
	http.HandleFunc("/", gw.PlaygroundHandler)

	log.Printf("Starting https server on port " + HttpPort)
	log.Printf("Starting graphql gateway !")

	err = http.ListenAndServe(":"+HttpPort, nil)
	if err != nil {
		log.Fatalf("Error https server on port %v: %v", HttpPort, err)
	}
}

func checkApis(urls ...string) []string {
	var workingUrls []string
	for _, url := range urls {
		res, err := http.Get(url)
		if err != nil {
			log.Printf("Error reaching %v api : %v", url, err)
			continue
		}
		log.Println(res)

		workingUrls = append(workingUrls, url)
	}

	return workingUrls
}

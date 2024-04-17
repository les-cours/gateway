package service

import (
	"log"
	"net/http"
)

func checkApis(urls ...string) []string {
	var workingUrls []string
	for _, url := range urls {
		_, err := http.Get(url)
		if err != nil {
			log.Printf("Error reaching %v api : %v", url, err)
			continue
		}

		workingUrls = append(workingUrls, url)
	}

	return workingUrls
}

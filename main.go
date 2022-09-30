package main

import (
	"log"
	"os"

	"github.com/natasha-h-s/AnchorHTMLCrawler/pkg/util"
)

func main() {
	log.Println("Starting webcrawler service...")

	url, exists := os.LookupEnv("URL")

	if !exists {
		log.Fatalf("Cannot find env var URL")
	}

	log.Printf("Retrieving information from %s", url)

	u := util.RetrieveURL(url)
	s := string(u[:])

	m := util.FindAnchors(s)

	if len(m) == 0 {
		log.Println("No links found, please try another URL. Exiting program...")
		return
	}

	for _, val := range m {
		log.Printf("HREF: %v \n", val)
	}
}

package main

import (
	"fmt"
	"http"
	"os"
	"log"
	"strings"

	urlshortener "google-api-go-client.googlecode.com/hg/urlshortener/v1"
)

func init() {
	registerDemo("urlshortener", urlshortener.UrlshortenerScope, urlShortenerMain)
}

func urlShortenerMain(client *http.Client, argv []string) {
	if len(argv) != 1 {
		fmt.Fprintf(os.Stderr, "Usage: urlshortener http://goo.gl/xxxxx     (to look up details)\n")
		fmt.Fprintf(os.Stderr, "       urlshortener http://example.com/long (to shorten)\n")
		return
	}

	svc, _ := urlshortener.New(client)
	urlstr := argv[0]

	// short -> long
	if strings.HasPrefix(urlstr, "http://goo.gl/") || strings.HasPrefix(urlstr, "https://goo.gl/") {
		url, err := svc.Url.Get(urlstr).Do()
		if err != nil {
			log.Fatalf("URL Get: %v", err)
		}
		fmt.Printf("Lookup of %s: %s\n", urlstr, url.LongUrl)
		return
	}

	// long -> short
	url, err := svc.Url.Insert(&urlshortener.Url{
		Kind:    "urlshortener#url", // Not really needed
		LongUrl: urlstr,
	}).Do()
	if err != nil {
		log.Fatalf("URL Insert: %v", err)
	}
	fmt.Printf("Shortened %s => %s\n", urlstr, url.Id)
}

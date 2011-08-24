package main

import (
	"flag"
	"http"
	"os"
	"log"

	buzz "google-api-go-client.googlecode.com/hg/buzz/v1"
)

func init() {
	registerDemo("buzz", buzz.BuzzScope, buzzMain)
}

func buzzMain(client *http.Client, argv []string) {
	fs := flag.NewFlagSet("buzz", flag.ExitOnError)
	includeMedia := fs.Bool("media", false, "Include an image")
	fs.Parse(argv)

	svc, _ := buzz.New(client)
	call := svc.Activities.Insert("@me", &buzz.Activity{
		Title: "a buzz post",
		Object: &buzz.ActivityObject{
			Type:    "http://activitystrea.ms/schema/1.0/note",
			Content: "the content of a buzz post",
		},
	})
	if *includeMedia {
		if f, err := findGopherFile(); err == nil {
			call.Media(f)
			defer f.Close()
		}
	}
	activity, err := call.Do()
	log.Printf("Buzz Activities.Insert: (%#v, %v)", activity, err)
}

func findGopherFile() (*os.File, os.Error) {
	f, err := os.Open("gopher.png")
	if err == nil {
		return f, nil
	}
	return os.Open("examples/gopher.png")
}

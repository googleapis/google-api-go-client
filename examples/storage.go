package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"google.golang.org/api/googleapi"
	storage "google.golang.org/api/storage/v1"
)

func init() {
	registerDemo("storage", storage.DevstorageReadWriteScope, storageMain)
}

func storageMain(client *http.Client, argv []string) {
	if len(argv) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: storage filename bucket (to upload an object)")
		return
	}

	service, err := storage.New(client)
	if err != nil {
		log.Fatalf("Unable to create Storage service: %v", err)
	}

	filename := argv[0]
	bucket := argv[1]

	goFile, err := os.Open(filename)
	if err != nil {
		log.Fatalf("error opening %q: %v", filename, err)
	}
	storageObject, err := service.Objects.Insert(bucket, &storage.Object{Name: filename}).Media(goFile).Do()
	log.Printf("Got storage.Object, err: %#v, %v", storageObject, err)

	// Test If-None-Match - should get a "HTTP 304 Not Modified" response.
	obj, err := service.Objects.Get(bucket, filename).IfNoneMatch(storageObject.Etag).Do()
	log.Printf("Got obj, err: %#v, %v", obj, err)
	if googleapi.IsError(err, http.StatusNotModified) {
		log.Printf("Success. Object not modified since upload.")
	} else {
		log.Printf("Error: expected object to not be modified since upload.")
	}
}

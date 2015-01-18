package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

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
	if err != nil {
		return
	}

	resp, err := service.Objects.Get(bucket, filename).Download()
	if err != nil {
		log.Fatalf("error downloading %q: %v", filename, err)
	}
	defer resp.Body.Close()

	n, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		log.Fatalf("error downloading %q: %v", filename, err)
	}

	log.Printf("Downloaded %d bytes", n)
}

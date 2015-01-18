package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	storage "google.golang.org/api/storage/v1"
)

func init() {
	registerDemo("storage", storage.DevstorageRead_writeScope, storageMain)
}

func storageMain(client *http.Client, argv []string) {
	if len(argv) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: storage filename bucket (to upload an object)")
		return
	}

	service, _ := storage.New(client)
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

	httpResp, err := service.Objects.Get(bucket, filename).Download()
	var downloaded int
	if err == nil {
		defer httpResp.Body.Close()
		body, _ := ioutil.ReadAll(httpResp.Body)
		downloaded = len(body)
	}
	log.Printf("Downloaded %d bytes, err: %v", downloaded, err)
}

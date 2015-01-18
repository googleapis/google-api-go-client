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
	if err != nil {
		log.Fatalf("error downloading %q: %v", filename, err)
	}
	defer httpResp.Body.Close()

	downloaded, err := io.Copy(ioutil.Discard, httpResp.Body)
	if err != nil {
		log.Fatalf("error downloading %q: %v", filename, err)
	}

	log.Printf("Downloaded %d bytes", downloaded)
}

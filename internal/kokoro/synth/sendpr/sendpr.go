// Copyright 2020 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	title := flag.String("title", "", "title for PR")
	flag.Parse()

	payload := map[string]interface{}{
		"title": *title,
		"body":  "", // TODO
		"head":  "synth-update",
		"base":  "master",
	}
	body, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", "https://api.github.com/repos/googleapis/google-api-go-client/pulls", bytes.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Authorization", "token "+os.Getenv("GITHUB_TOKEN"))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode > 299 {
		log.Fatalf("non-OK response creating PR: %s", bodyBytes)
	}

	log.Printf("PR created: %s", bodyBytes)
}

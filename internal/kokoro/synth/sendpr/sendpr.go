// Copyright 2020 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
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
	req := http.NewRequest("POST", "https://api.github.com/repos/googleapis/google-api-go-client/pulls", bytes.NewReader(body))
	req.Headers.Set("Accept", "application/vnd.github.v3+json")
	req.Headers.Set("Authorization", "token "+os.Getenv("GITHUB_TOKEN"))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode > 299 {
		log.Fatal("non-OK response creating PR: %s", body)
	}

	log.Print("PR created: %s", body)
}

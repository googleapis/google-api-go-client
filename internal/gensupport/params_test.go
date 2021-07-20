package gensupport

import (
	"testing"

	"google.golang.org/api/googleapi"
)

func TestSetOptionsGetMulti(t *testing.T) {
	co := googleapi.QueryParameter("key", "foo", "bar")
	urlParams := make(URLParams)
	SetOptions(urlParams, co)
	if got, want := urlParams.Encode(), "key=foo&key=bar"; got != want {
		t.Fatalf("URLParams.Encode() = %q, want %q", got, want)
	}
}

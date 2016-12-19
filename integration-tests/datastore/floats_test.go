// +build integration

// Test float marshal/unmarshal.

package datastore

import (
	"encoding/json"
	"math"
	"testing"

	ds "google.golang.org/api/datastore/v1"
)

func TestUnmarshal(t *testing.T) {
	for _, test := range []struct {
		in   string
		want ds.LatLng
	}{
		{`{"latitude": 0, "longitude": 0}`, ds.LatLng{Latitude: 0, Longitude: 0}},
		{`{"latitude": 1.2, "longitude": "Infinity"}`, ds.LatLng{Latitude: 1.2, Longitude: math.Inf(1)}},
		{`{"latitude": "NaN", "longitude": "-Infinity"}`, ds.LatLng{Latitude: math.NaN(), Longitude: math.Inf(-1)}},
	} {
		var got ds.LatLng
		if err := json.Unmarshal([]byte(test.in), &got); err != nil {
			t.Fatal(err)
		}
		if !fleq(got.Latitude, test.want.Latitude) || !fleq(got.Longitude, test.want.Longitude) {
			t.Errorf("got\n%+v\nwant\n%+v", got, test.want)
		}
	}
}

func fleq(f1, f2 float64) bool {
	return f1 == f2 || (math.IsNaN(f1) && math.IsNaN(f2))
}

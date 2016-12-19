// +build integration

// Test float marshal/unmarshal with *float64.

package monitoring

import (
	"encoding/json"
	"math"
	"reflect"
	"testing"

	mon "google.golang.org/api/monitoring/v3"
)

func TestUnmarshal(t *testing.T) {
	x := "x"
	f3 := 3.0
	fInf := math.Inf(1)
	for _, test := range []struct {
		in   string
		want mon.TypedValue
	}{
		{`{"stringValue": "x"}`, mon.TypedValue{StringValue: &x}},
		{`{"doubleValue": 3}`, mon.TypedValue{DoubleValue: &f3}},
		{`{"doubleValue": "Infinity"}`, mon.TypedValue{DoubleValue: &fInf}},
	} {
		var got mon.TypedValue
		if err := json.Unmarshal([]byte(test.in), &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("got\n%+v\nwant\n%+v", got, test.want)
		}
	}
}

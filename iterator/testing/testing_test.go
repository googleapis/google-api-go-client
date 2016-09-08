package testing

import (
	"reflect"
	"testing"
)

func TestCompareSlices(t *testing.T) {
	equal := func(a, b []int) bool {
		_, ok := compareSlices(reflect.ValueOf(a), reflect.ValueOf(b))
		return ok
	}
	for _, test := range []struct {
		a, b      []int
		wantEqual bool
	}{
		{nil, nil, true},
		{nil, []int{}, true},
		{[]int{1, 2}, []int{1, 2}, true},
		{[]int{1}, []int{1, 2}, false},
		{[]int{1, 2}, []int{1}, false},
		{[]int{1, 2}, []int{1, 3}, false},
	} {
		if got, want := equal(test.a, test.b), test.wantEqual; got != want {
			t.Errorf("%v, %v: got %t, want %t", test.a, test.b, got, want)
		}
	}
}

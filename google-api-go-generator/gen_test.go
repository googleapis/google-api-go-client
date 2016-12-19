package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"math"
	"path/filepath"
	"reflect"
	"testing"

	"google.golang.org/api/internal"
)

var updateGolden = flag.Bool("update_golden", false, "If true, causes TestAPIs to update golden files")

func TestAPIs(t *testing.T) {
	names := []string{
		"any",
		"arrayofarray-1",
		"arrayofenum",
		"arrayofmapofobjects",
		"arrayofmapofstrings",
		"blogger-3",
		"getwithoutbody",
		"mapofany",
		"mapofarrayofobjects",
		"mapofobjects",
		"mapofstrings-1",
		"param-rename",
		"quotednum",
		"repeated",
		"resource-named-service", // appengine/v1/appengine-api.json
		"unfortunatedefaults",
		"variants",
		"wrapnewlines",
	}
	for _, name := range names {
		t.Logf("TEST %s", name)
		api, err := apiFromFile(filepath.Join("testdata", name+".json"))
		if err != nil {
			t.Errorf("Error loading API testdata/%s.json: %v", name, err)
			continue
		}
		clean, err := api.GenerateCode()
		if err != nil {
			t.Errorf("Error generating code for %s: %v", name, err)
			continue
		}
		goldenFile := filepath.Join("testdata", name+".want")
		if *updateGolden {
			if err := ioutil.WriteFile(goldenFile, clean, 0644); err != nil {
				t.Fatal(err)
			}
		}
		want, err := ioutil.ReadFile(goldenFile)
		if err != nil {
			t.Error(err)
			continue
		}
		if !bytes.Equal(want, clean) {
			tf, _ := ioutil.TempFile("", "api-"+name+"-got-json.")
			tf.Write(clean)
			tf.Close()
			t.Errorf("Output for API %s differs: diff -u %s %s", name, goldenFile, tf.Name())
		}
	}
}

func TestScope(t *testing.T) {
	tests := [][]string{
		{
			"https://www.googleapis.com/auth/somescope",
			"SomescopeScope",
		},
		{
			"https://mail.google.com/somescope",
			"MailGoogleComSomescopeScope",
		},
		{
			"https://mail.google.com/",
			"MailGoogleComScope",
		},
	}
	for _, test := range tests {
		if got := scopeIdentifierFromURL(test[0]); got != test[1] {
			t.Errorf("scopeIdentifierFromURL(%q) got %q, want %q", test[0], got, test[1])
		}
	}
}

func TestDepunct(t *testing.T) {
	tests := []struct {
		needCap  bool
		in, want string
	}{
		{
			needCap: true,
			in:      "part__description",
			want:    "Part__description",
		},
		{
			needCap: true,
			in:      "Part_description",
			want:    "PartDescription",
		},
		{
			needCap: false,
			in:      "part_description",
			want:    "partDescription",
		},
		{
			needCap: false,
			in:      "part-description",
			want:    "partDescription",
		},
		{
			needCap: false,
			in:      "part.description",
			want:    "partDescription",
		},
		{
			needCap: false,
			in:      "part$description",
			want:    "partDescription",
		},
		{
			needCap: false,
			in:      "part/description",
			want:    "partDescription",
		},
		{
			needCap: true,
			in:      "Part__description_name",
			want:    "Part__descriptionName",
		},
		{
			needCap: true,
			in:      "Part_description_name",
			want:    "PartDescriptionName",
		},
		{
			needCap: true,
			in:      "Part__description__name",
			want:    "Part__description__name",
		},
		{
			needCap: true,
			in:      "Part_description__name",
			want:    "PartDescription__name",
		},
	}
	for _, test := range tests {
		if got := depunct(test.in, test.needCap); got != test.want {
			t.Errorf("depunct(%q,%v) = %q; want %q", test.in, test.needCap, got, test.want)
		}
	}
}

func TestRenameVersion(t *testing.T) {
	tests := []struct {
		version, want string
	}{
		{
			version: "directory_v1",
			want:    "directory/v1",
		},
		{
			version: "email_migration_v1",
			want:    "email_migration/v1",
		},
		{
			version: "my_api_v1.2",
			want:    "my_api/v1.2",
		},
	}
	for _, test := range tests {
		if got := renameVersion(test.version); got != test.want {
			t.Errorf("renameVersion(%q) = %q; want %q", test.version, got, test.want)
		}
	}
}

// The definitions and test below validate our method for handling special float values
// in JSON backwards-compatibly.

// S represents a generated struct that will be filled by unmarshaling JSON, like a
// response. S contains a field of type float64. If we unmarshaled JSON into S
// directly, then the special float values "NaN", "Infinity" and "-Infinity" would
// result in errors, because the encoding/json package doesn't recognize them.
type S struct {
	I int
	F float64 `json:"f,omitempty"`
}

// We generate this UnmarshalJSON function for S.
func (s *S) UnmarshalJSON(data []byte) error {
	type nomethod S

	var s1 struct {
		// This field has the same name as S.F, but it is defined at a higher level
		// (because S.F is embedded), so it hides S.F. The JSON key "f" will unmarshal
		// into this field. The internal.JSONFloat64 type correctly handles the special
		// float values.
		F internal.JSONFloat64 `json:"f"`

		// S is effectively embedded here. All the other fields of S will be visible in
		// s1, so they will be populated normally during unmarshaling.
		*nomethod
	}
	// Now we place s, the struct value that we ultimately want to fill correctly, into
	// the embedded field.
	s1.nomethod = (*nomethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	// s has been populated from the JSON, except for field F, which was hidden by s1.
	// So we copy s1.F into s.F
	s.F = float64(s1.F)
	return nil
}

func TestUnmarshalFloat(t *testing.T) {
	s := S{I: 1, F: 2}
	in := `{"i": 6, "f": "-Infinity"}`
	if err := json.Unmarshal([]byte(in), &s); err != nil {
		t.Fatal(err)
	}
	want := S{I: 6, F: math.Inf(-1)}
	if !reflect.DeepEqual(s, want) {
		t.Errorf("got %+v, want %+v", s, want)
	}
}

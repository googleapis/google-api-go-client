// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"testing"
)

// Files in this package use two separate licenses:

// 1. A BSD license (used for generated files).
var sentinelBSD = regexp.MustCompile(`// Copyright \d\d\d\d (Google LLC|The Go Authors)(\.)*( All rights reserved\.)*
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
`)

// 2. An Apache license (used for handwritten files).
var sentinelApache = regexp.MustCompile(`(//|#) Copyright \d\d\d\d Google LLC
(//|#)
(//|#) Licensed under the Apache License, Version 2\.0 \(the "License"\);
(//|#) you may not use this file except in compliance with the License\.
(//|#) You may obtain a copy of the License at
(//|#)
(//|#)     https://www\.apache\.org/licenses/LICENSE-2\.0
(//|#)
(//|#) Unless required by applicable law or agreed to in writing, software
(//|#) distributed under the License is distributed on an "AS IS" BASIS,
(//|#) WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied\.
(//|#) See the License for the specific language governing permissions and
(//|#) limitations under the License\.`)

const prefix = "// Copyright"

// A few files have to be skipped.
var skip = map[string]bool{
	"tools.go": true, // This file requires another comment above the license.
	"googleapi/internal/uritemplates/uritemplates.go": true, // This file is licensed to an individual.
}

// This test validates that all go files in the repo start with an appropriate license.
func TestLicense(t *testing.T) {
	err := filepath.Walk(".", func(path string, fi os.FileInfo, err error) error {
		if skip[path] {
			return nil
		}

		if err != nil {
			return err
		}

		if filepath.Ext(path) != ".go" && filepath.Ext(path) != ".sh" {
			return nil
		}

		src, err := ioutil.ReadFile(path)
		if err != nil {
			return nil
		}

		// Verify that one of the two valid licenses is matched.
		if !sentinelBSD.Match(src) && !sentinelApache.Match(src) {
			t.Errorf("%v: license header not present", path)
			return nil
		}

		// Also check it is at the top of .go files (but not .sh files, because they must have a shebang first).
		if filepath.Ext(path) == ".go" && !bytes.HasPrefix(src, []byte(prefix)) {
			t.Errorf("%v: license header not at the top", path)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}

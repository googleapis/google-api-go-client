// Copyright 2018 Google LLC
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// deprecatedPkgs is a map that contains packages that should be
// deprecated in favor of another package.
type deprecatedPkgs map[string]*apiReplacement

func (d deprecatedPkgs) Get(name, version string) string {
	ar := d[name]
	if ar == nil {
		return ""
	}
	for _, v := range ar.versions {
		if v == "*" || v == version {
			return ar.replacementPkg
		}
	}
	return ""
}

// replacementPackage is a map from an API package name to the
// import path of the package that replaces it. If an API appears
// in this map, its package doc comment will note that it is deprecated
// and point to the replacement.
// TODO(jba): consider automating this by looking at the structure of the gocloud repo.
var replacementPackage deprecatedPkgs = map[string]*apiReplacement{
	"bigquery": {
		replacementPkg: "cloud.google.com/go/bigquery",
		versions:       []string{"*"},
	},
	"cloudkms": {
		replacementPkg: "cloud.google.com/go/kms/apiv1",
		versions:       []string{"*"},
	},
	"cloudtasks": {
		replacementPkg: "cloud.google.com/go/cloudtasks/apiv2beta2",
		versions:       []string{"*"},
	},
	"dataproc": {
		replacementPkg: "cloud.google.com/go/dataproc/apiv1",
		versions:       []string{"*"},
	},
	"datastore": {
		replacementPkg: "cloud.google.com/go/datastore",
		versions:       []string{"*"},
	},
	"dialogflow": {
		replacementPkg: "cloud.google.com/go/dialogflow/apiv2",
		versions:       []string{"*"},
	},
	"dlp": {
		replacementPkg: "cloud.google.com/go/dlp/apiv2",
		versions:       []string{"*"},
	},
	"firestore": {
		replacementPkg: "cloud.google.com/go/firestore",
		versions:       []string{"*"},
	},
	"language": {
		replacementPkg: "cloud.google.com/go/language/apiv1",
		versions:       []string{"*"},
	},
	"logging": {
		replacementPkg: "cloud.google.com/go/logging",
		versions:       []string{"*"},
	},
	"monitoring": {
		replacementPkg: "cloud.google.com/go/monitoring/apiv3",
		versions:       []string{"v3"},
	},
	"oslogin": {
		replacementPkg: "cloud.google.com/go/oslogin/apiv1",
		versions:       []string{"*"},
	},
	"pubsub": {
		replacementPkg: "cloud.google.com/go/pubsub",
		versions:       []string{"*"},
	},
	"redis": {
		replacementPkg: "cloud.google.com/go/redis/apiv1",
		versions:       []string{"*"},
	},
	"spanner": {
		replacementPkg: "cloud.google.com/go/spanner",
		versions:       []string{"*"},
	},
	"speech": {
		replacementPkg: "cloud.google.com/go/speech/apiv1",
		versions:       []string{"*"},
	},
	"texttospeech": {
		replacementPkg: "cloud.google.com/go/texttospeech/apiv1",
		versions:       []string{"*"},
	},
	"translate": {
		replacementPkg: "cloud.google.com/go/translate",
		versions:       []string{"*"},
	},
	"videointelligence": {
		replacementPkg: "cloud.google.com/go/videointelligence/apiv1",
		versions:       []string{"*"},
	},
	"vision": {
		replacementPkg: "cloud.google.com/go/vision/apiv1",
		versions:       []string{"*"},
	},
	"storage": {
		replacementPkg: "cloud.google.com/go/storage",
		versions:       []string{"*"},
	},
}

//
type apiReplacement struct {
	// replacementPkg is a reference to a package that should be used instead.
	replacementPkg string
	// versions is a slice of API versions for which the replacementPkg should
	// be used. The value `*` means the replacement should take place for all
	// versions
	versions []string
}

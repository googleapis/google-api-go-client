// Copyright 2018 Google LLC
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// deprecatedPkgs is a map that contains packages that should be
// deprecated in favor of another package.
type deprecatedPkgs map[string]apiReplacement

func (d deprecatedPkgs) Get(name, version string) string {
	ar, ok := d[name]
	if !ok {
		return ""
	}
	if len(ar.versions) == 0 {
		return ar.replacementPkg
	}
	for _, v := range ar.versions {
		if v == version {
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
var replacementPackage deprecatedPkgs = map[string]apiReplacement{
	"bigquery": {
		replacementPkg: "cloud.google.com/go/bigquery",
	},
	"cloudkms": {
		replacementPkg: "cloud.google.com/go/kms/apiv1",
	},
	"cloudtasks": {
		replacementPkg: "cloud.google.com/go/cloudtasks/apiv2beta2",
	},
	"dataproc": {
		replacementPkg: "cloud.google.com/go/dataproc/apiv1",
	},
	"datastore": {
		replacementPkg: "cloud.google.com/go/datastore",
	},
	"dialogflow": {
		replacementPkg: "cloud.google.com/go/dialogflow/apiv2",
	},
	"dlp": {
		replacementPkg: "cloud.google.com/go/dlp/apiv2",
	},
	"firestore": {
		replacementPkg: "cloud.google.com/go/firestore",
	},
	"language": {
		replacementPkg: "cloud.google.com/go/language/apiv1",
	},
	"logging": {
		replacementPkg: "cloud.google.com/go/logging",
	},
	"monitoring": {
		replacementPkg: "cloud.google.com/go/monitoring/apiv3",
		versions:       []string{"v3"},
	},
	"oslogin": {
		replacementPkg: "cloud.google.com/go/oslogin/apiv1",
	},
	"pubsub": {
		replacementPkg: "cloud.google.com/go/pubsub",
	},
	"redis": {
		replacementPkg: "cloud.google.com/go/redis/apiv1",
	},
	"spanner": {
		replacementPkg: "cloud.google.com/go/spanner",
	},
	"speech": {
		replacementPkg: "cloud.google.com/go/speech/apiv1",
	},
	"texttospeech": {
		replacementPkg: "cloud.google.com/go/texttospeech/apiv1",
	},
	"translate": {
		replacementPkg: "cloud.google.com/go/translate",
	},
	"videointelligence": {
		replacementPkg: "cloud.google.com/go/videointelligence/apiv1",
	},
	"vision": {
		replacementPkg: "cloud.google.com/go/vision/apiv1",
	},
	"storage": {
		replacementPkg: "cloud.google.com/go/storage",
	},
}

//
type apiReplacement struct {
	// replacementPkg is a reference to a package that should be used instead.
	replacementPkg string
	// versions is a slice of API versions for which the replacementPkg should
	// be used. A zero length slice means the replacement should take place for
	// all versions.
	versions []string
}

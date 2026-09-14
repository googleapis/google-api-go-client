#!/bin/bash

# Copyright 2019 Google LLC.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

# Fail on error, and display commands being run.
set -ex

# Only do the formatting and mod checking for the "latest" version
if [[ $KOKORO_JOB_NAME != *"latest-version"* ]]; then
  # Fail if a dependency was added without the necessary go.mod/go.sum change
  # being part of the commit.
  go mod tidy
  git diff go.mod | tee /dev/stderr | (! read)
  git diff go.sum | tee /dev/stderr | (! read)

  # Easier to debug CI.
  pwd

  gofmt -s -d -l . 2>&1 | tee /dev/stderr | (! read)
  goimports -l . 2>&1 | tee /dev/stderr | (! read)

fi

# Only do staticcheck on our "earliest" version.
if [[ $KOKORO_JOB_NAME != *"earliest-version"* ]]; then
  staticcheck ./... 2>&1 | (
    grep -v "SA1019" |
      grep -v "S1007" |
      grep -v "error var Done should have name of the form ErrFoo" |
      grep -v "examples" || true
  ) | tee /dev/stderr | (! read)
fi

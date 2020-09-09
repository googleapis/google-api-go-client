#!/bin/bash

# Copyright 2020 Google LLC.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

set -e -x

cd $(dirname $0)/../../..
go version
go env

pushd google-api-go-generator
make all
popd

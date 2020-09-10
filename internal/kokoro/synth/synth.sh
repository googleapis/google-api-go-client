#!/bin/bash

# Copyright 2020 Google LLC.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

set -e -x

cd $(dirname $0)/../../..

# Install Go 1.15
rm -rf /usr/local/go
curl -o /tmp/go.tgz https://dl.google.com/go/go1.15.1.linux-amd64.tar.gz &&
    tar -C /usr/local -xzf /tmp/go.tgz &&
    rm /tmp/go.tgz &&
    export PATH=$PATH:/usr/local/go/bin

go version
go env

pushd google-api-go-generator
make all
popd

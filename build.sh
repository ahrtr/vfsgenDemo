#!/bin/sh

set -ex

go generate ./data

TAGS="${TAGS:-}"
TAGS="${TAGS} release"

go build -tags "${TAGS}" -o "bin/vfsgen-demo"  .

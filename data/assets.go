// +build !release
//go:generate go run generate.go

package data

import (
	"net/http"
)

// Assets contains project assets.
var Assets http.FileSystem = http.Dir("data")

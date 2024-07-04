//go:build !release
// +build !release

package supervisord_main

import (
	"net/http"
)

// HTTP auto generated
var HTTP http.FileSystem = http.Dir("./webgui")

package src

import (
	"path"
	"runtime"
)

func Root() (pwd string) {
	_, fullFilename, _, _ := runtime.Caller(0)
	pwd = path.Dir(fullFilename)
	return
}

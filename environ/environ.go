package environ

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

var (
	packageroot string
)

func init() {
	packageroot = findpkgroot()
}

func findpkgroot() string {
	_, current, _, ok := runtime.Caller(0)
	if !ok {
		return getwd()
	}

	for {
		dir := filepath.Dir(current)
		defer func() { current = dir }()

		isRoot := dir == filepath.Dir(dir)
		_, err := os.Stat(filepath.Join(dir, "go.mod"))
		// ✅ Found the directory with go.mod
		if err == nil {
			return dir
		}

		if isRoot {
			return getwd()
		}
	}
}

func getwd() string {
	pwd, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("os.Getwd: %w", err))
	}

	return pwd
}

// PackageRoot returns the nearest directory with a go.mod file. Generally that is the root of the package/project.
func PackageRoot() string {
	return packageroot
}

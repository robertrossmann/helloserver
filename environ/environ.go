package environ

import (
	"os"
	"path/filepath"
	"runtime"
)

var (
	packageroot string
)

func findpkgroot() string {
	_, here, _, _ := runtime.Caller(0)
	dirpath := filepath.Dir(here)

	for {
		candidate := filepath.Join(dirpath, "go.mod")
		isRootDirectory := dirpath == filepath.Dir(dirpath)

		// go.mod is not present, use current working directory
		if isRootDirectory {
			pwd, err := os.Getwd()
			if err != nil {
				panic(err)
			}

			return pwd
		}

		_, err := os.Stat(candidate)
		if err == nil {
			break
		} else {
			dirpath = filepath.Dir(dirpath)
		}
	}

	return dirpath
}

func init() {
	packageroot = findpkgroot()
}

// PackageRoot returns the nearest directory with a go.mod file. Generally that is the root of the package.
func PackageRoot() string {
	return packageroot
}

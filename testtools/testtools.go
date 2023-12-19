package testtools

import (
	"helloserver/environ"
	"path/filepath"

	"github.com/joho/godotenv"
)

// Bootstrap prepares the environment for testing. Call this in every test file. ⚠️
func Bootstrap() {
	pkgroot := environ.PackageRoot()
	err := godotenv.Load(filepath.Join(pkgroot, ".env.test"))
	if err != nil {
		panic(err)
	}
}

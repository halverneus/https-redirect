package version

import (
	"fmt"
	"runtime"
)

// Run print operation.
func Run() error {
	fmt.Printf("%s\n%s\n", VersionText, GoVersionText)
	return nil
}

var (
	// version is the application version set during build.
	version string

	// VersionText for directly accessing the redirect version.
	VersionText = fmt.Sprintf("v%s", version)

	// GoVersionText for directly accessing the version of the Go runtime
	// compiled with the redirect.
	GoVersionText = runtime.Version()
)

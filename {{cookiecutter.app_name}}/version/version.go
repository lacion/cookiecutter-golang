package version

import (
	"fmt"
	"runtime"
)

// GitCommit : The git commit that was compiled. This will be filled in by the compiler.
var GitCommit string

// Version : The main version number that is being run at the moment based on git tag.
var Version string

var BuildDate = ""

var GoVersion = runtime.Version()

var OsArch = fmt.Sprintf("%s %s", runtime.GOOS, runtime.GOARCH)

# Go-Version

Go-Version is a binary version command helper. It creates a `-v` or `--version` flag in your binary that prints contents of a version file, the repo revision and the build time.

## Usage

The command `go-version-helper` does the following:

- Creates a version file (default `.version`) if it does not exists 
- Sets the contents of the create file to the current version (default `v0.1`)
- Creates a build file (default `build.sh`) for the binary package (default `.`)

In your main file you have just to import the package `github.com/klaidliadon/go-version` 
and wrap your main function around a `version.App`.

```go
package main

import (
	"github.com/klaidliadon/go-version"

	"fmt"
)

func main() {
	a := version.NewApp("Sample App")
	a.Run(func() {
		// if -v or --version is specified prints the version and exits
		fmt.Println("app running")
	})
}

```


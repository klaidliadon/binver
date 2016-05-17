# Binver

Binary version command helper.

## Usage

The command `binver` does the following:

- Creates a version file (default `.version`) if it does not exists 
- Sets the contents of the create file to the current version (default `v0.1`)
- Creates a build file (default `build.sh`) for the binary package (default `.`)

In your main file you have just to import the package `github.com/klaidliadon/binver` 
and call the function `binver.Version("YourAppName")`.

```go
package main

import (
	"github.com/klaidliadon/binver"

	"fmt"
)

func main() {
	// if -v or --version is specified 
	// prints the version and exits
	binver.Version("Sample App")
	fmt.Println("Start")
	// ...
	fmt.Println("End")
}

```
package binver

import (
	"flag"
	"fmt"
	"os"
)

var (
	version   = "n/a"
	buildTime = "n/a"
	revison   = "n/a"
)

var (
	short = flag.Bool("v", false, "Show version")
	long  = flag.Bool("version", false, "Show version")
)

func Version(appName string) {
	if !flag.Parsed() {
		flag.Parse()
	}
	if !*short && !*long {
		return
	}
	fmt.Printf("%s %s (%s) %s\n", appName, version, revison[:7], buildTime)
	os.Exit(0)
}

package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	contents = `#!/bin/bash
if [ ! -d ".git" ]; then
	 echo "binver build requires a git repo!"
	 exit 1
fi
go build -ldflags "\
-X github.com/klaidliadon/go-version.version=$(cat ` + vToken + `) \
-X github.com/klaidliadon/go-version.buildTime=$(date -u '+%Y-%m-%dT%H:%M:%S') \
-X github.com/klaidliadon/go-version.revison=$(git rev-parse HEAD)\
" `
	vToken = "[version]"
)

var (
	fBuildFile   = flag.String("o", "build.sh", "output file")
	fPackageName = flag.String("p", ".", "package name")
	fVersionFile = flag.String("f", ".version", "version file")
	fCurrVersion = flag.String("v", "v0.1", "current version")
)

func init() {
	flag.Parse()
}

func main() {
	if _, err := os.Stat(*fVersionFile); err != nil {
		if err := ioutil.WriteFile(*fVersionFile, []byte(*fCurrVersion), 0644); err != nil {
			log.Fatalf("Cannot open version file %q: %s", *fVersionFile, err)
		}
	}
	txt := strings.Replace(contents, vToken, *fVersionFile, 1) + *fPackageName
	if err := ioutil.WriteFile(*fBuildFile, []byte(txt), 0755); err != nil {
		log.Fatalf("Cannot write build file %q: %s", *fBuildFile, err)
	}
}

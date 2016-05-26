package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
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
	txt := fmt.Sprintf(`#!/bin/bash
if [ ! -d ".git" ]; then
	 echo "%s build requires a git repo!"
	 exit 1
fi
go build -ldflags "\
-X github.com/klaidliadon/go-version.version=$(cat %s) \
-X github.com/klaidliadon/go-version.buildTime=$(date -u '+%%Y-%%m-%%dT%%H:%%M:%%S') \
-X github.com/klaidliadon/go-version.revison=$(git rev-parse HEAD)\
" %s`, os.Args[0], *fVersionFile, *fPackageName)
	if err := ioutil.WriteFile(*fBuildFile, []byte(txt), 0755); err != nil {
		log.Fatalf("Cannot write build file %q: %s", *fBuildFile, err)
	}
}

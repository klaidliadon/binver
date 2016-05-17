package version

import (
	"flag"
	"fmt"
	"os"
	"time"
)

const timeFormat = "2006-01-02T15:04:05"

var (
	version   = "n/a"
	buildTime = "n/a"
	revison   = "0000000000000000000000000000000000000000"
)

var (
	short = flag.Bool("v", false, "Show version")
	long  = flag.Bool("version", false, "Show version")
)

type App struct {
	Name      string
	Version   string
	Revision  string
	BuildTime time.Time
}

func NewApp(appName string) *App {
	t, _ := time.Parse(timeFormat, buildTime)
	return &App{Name: appName, Version: version, Revision: revison, BuildTime: t}
}

func (a *App) Run(mainFn func()) {
	if !flag.Parsed() {
		flag.Parse()
	}
	if !*short && !*long {
		mainFn()
		return
	}
	fmt.Print(a)
	os.Exit(0)
}

func (a *App) String() string {
	return fmt.Sprintf("%s %s (%s) %s\n", a.Name, a.Version, a.Revision[:7], a.BuildTime.Format(timeFormat))
}

package main

import (
    "flag"
    "fmt"
    "os"

    "github.com/AcidGo/zabbix-certificate/pkg/target"
)

var (
    // flag vars
    targetMode  string
    targetAddr  string
    showVerbose bool

    // app info
    AppName             string
    AppAuthor           string
    AppVersion          string
    AppGitCommitHash    string
    AppBuildTime        string
    AppGoVersion        string
)

func init() {
    flag.StringVar(&targetMode, "m", "", "the mode for checking target certificate")
    flag.StringVar(&targetAddr, "h", "", "the address for checking target certificate")
    flag.BoolVar(&showVerbose, "v", false, "show the verbose certificates result list")
    flag.Usage = flagUsage
    flag.Parse()
}

func main() {
    tg, err := target.NewTarget(targetMode, targetAddr)
    if err != nil {
        panic(err)
    }

    res, err := tg.Check()
    if err != nil {
        panic(err)
    }

    if showVerbose {
        fmt.Print(res.Verbose())
    } else {
        fmt.Print(res.TopOne())
    }
}

func flagUsage() {
    usageMsg := fmt.Sprintf(`App: %s
Version: %s
Author: %s
GitCommit: %s
BuildTime: %s
GoVersion: %s
Options:
`, AppName, AppVersion, AppAuthor, AppGitCommitHash, AppBuildTime, AppGoVersion)

    fmt.Fprintf(os.Stderr, usageMsg)
    flag.PrintDefaults()
}
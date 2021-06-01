package main

import (
    "flag"
    "fmt"
    "log"
    "strconv"

    "github.com/AcidGo/zabbix-certificate/pkg/target"
)

var (
    // flag vars
    targetMode  string
    targetAddr  string
    showVerbose bool
)

func init() {
    flag.StringVar(&targetMode, "m", "", "the mode for checking target certificate")
    flag.StringVar(&targetAddr, "h", "", "the address for checking target certificate")
    flag.BoolVar(&showVerbose, "v", false, "show the verbose certificates result list")
    flag.Parse()
}

func main() {
    tg, err := target.NewTarget(targetMode, targetAddr)
    if err != nil {
        log.Fatal(err)
    }

    res, err := tg.Check()
    if err != nil {
        log.Fatal(err)
    }

    if showVerbose {
        fmt.Print(res.Verbose())
    } else {
        fmt.Print(res.TopOne())
    }
}
package target

import (
    "fmt"

    "github.com/AcidGo/zabbix-certificate/pkg/check"
)

type Target struct {
    Mode    string
    Addr    string
    f       func(string) (check.Results, error)
}

func NewTarget(mode, addr string) (*Target, error) {
    t := &Target{Mode: mode, Addr: addr}
    return t, t.validate()
}

func (t *Target) validate() (error) {
    const (
        HTTPSMode   = "https"
    )

    var err error

    switch t.Mode {
    case HTTPSMode:
        t.f = check.HTTPSCheck
    default:
        err = fmt.Errorf("cannot support the chekcing mode %s now", t.Mode)
    }

    if err == nil && t.Addr == "" {
        err = fmt.Errorf("address of target is empty")
    }

    return err
}

func (t *Target) Check() (check.Results, error) {
    if f == nil {
        return nil, fmt.Errorf("the inner checking func is nil")
    }

    return t.f(t.Mode)
}
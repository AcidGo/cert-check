package check

import (
    "fmt"
    "strconv"
    "strings"
    "time"
)

type Result struct {
    Issuer          string
    Subject         string
    ExpireUnixtime  int64
}

type Results []*Result

func (rs *Results) Verbose() (string) {
    var sb strings.Builder
    first := true

    for _, r := range *rs {
        if first {
            first = false
        } else {
            sb.WriteString("--------------------\n")
        }

        sb.WriteString(fmt.Sprintf("issuer: %s\n", r.Issuer))
        sb.WriteString(fmt.Sprintf("subject: %s\n", r.Subject))
        sb.WriteString(fmt.Sprintf("expire date: %s\n", time.Unix(r.ExpireUnixtime, 0).Format(time.RFC3339)))
    }

    return sb.String()
}

func (rs *Results) TopOne() (string) {
    const MaxInt64 = int64(^uint(0)>>1)
    cur := MaxInt64

    for _, r := range *rs {
        if r.ExpireUnixtime < cur {
            cur = r.ExpireUnixtime
        }
    }

    if cur == MaxInt64 {
        return ""
    }

    return strconv.FormatInt(cur, 10)
}
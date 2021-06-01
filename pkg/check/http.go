package check

import (
    "crypto/tls"
    "net/http"
    "time"
)

func HTTPSCheck(addr string) (Results, error) {
    client := &http.Client{
        Transport: &http.Transport{
            TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
        },
        Timeout: 5*time.Second,
    }

    resp, err := client.Get(addr)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var res Results
    for _, cert := range resp.TLS.PeerCertificates {
        r := &Result{
            Issuer: cert.Issuer.String(),
            Subject: cert.Subject.String(),
            ExpireUnixtime: NotAfter.Unix(),
        }

        res = append(res, r)
    }

    return res, nil
}
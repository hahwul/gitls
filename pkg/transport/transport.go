package transport

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/hahwul/gitls/pkg/model"
)

// GetTransport is setting timetout and proxy on tranport
func GetTransport(options model.Options) *http.Transport {
	// set timeout
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
			Renegotiation:      tls.RenegotiateOnceAsClient,
		},
		DisableKeepAlives: true,
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			DualStack: true,
		}).DialContext,
	}
	// if use proxy mode , set proxy
	if options.Proxy != "" {
		proxyAddress, err := url.Parse(options.Proxy)
		if err != nil {
			msg := fmt.Sprintf("not running %v from proxy option", err)
			fmt.Println(msg)
		}
		transport.Proxy = http.ProxyURL(proxyAddress)
	}
	if options.UseTor {
		proxyAddress, err := url.Parse("http://localhost:9050")
		if err != nil {
			msg := fmt.Sprintf("not running %v from proxy option", err)
			fmt.Println(msg)
		}
		transport.Proxy = http.ProxyURL(proxyAddress)
	}
	return transport
}

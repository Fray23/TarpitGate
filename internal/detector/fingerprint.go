package detector

import (
	"net"
	"net/http"
	"net/url"
)


type Fingerprint struct {
	IP          string
	RealIP      string
	UserAgent   string
	Headers     http.Header
	Path        string
	Method      string
	QueryParams url.Values
	Proto       string
}


func Extract(r *http.Request) *Fingerprint {
	ip := r.RemoteAddr
	if host, _, err := net.SplitHostPort(ip); err == nil {
		ip = host
	}

	return &Fingerprint{
		IP:          ip,
		RealIP:      r.Header.Get("X-Real-IP"),
		UserAgent:   r.Header.Get("User-Agent"),
		Headers:     r.Header,
		Path:        r.URL.Path,
		Method:      r.Method,
		QueryParams: r.URL.Query(),
		Proto:       r.Proto,
	}
}

func (f *Fingerprint) Analyze() Score {
	var s Score
	checkUserAgent(f, &s)
	checkHeaders(f, &s)
	checkPath(f, &s)
	checkProto(f, &s)
	checkMethod(f, &s)
	return s
}

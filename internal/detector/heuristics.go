package detector

import (
	"strings"
)

// TODO replace to redis
var noBrowserUASubstrings = []string{
	"curl/", "wget/", "python-requests", "python-urllib",
}

var botKeywords = []string{
	"bot", "crawler", "spider", "scraper", "scan", "fetch",
}

func checkUserAgent(f *Fingerprint, s *Score) {
	ua := strings.ToLower(strings.TrimSpace(f.UserAgent))
	if ua == "" {
		s.add(40, "empty User-Agent")
		return
	}
	if len(ua) < 15 {
		s.add(25, "suspiciously short User-Agent")
	}
	for _, sub := range noBrowserUASubstrings {
		if strings.Contains(ua, sub) {
			s.add(35, "client-library UA: "+sub)
			return
		}
	}
	for _, kw := range botKeywords {
		if strings.Contains(ua, kw) {
			s.add(20, "bot keyword in UA: "+kw)
			return
		}
	}
	if !strings.HasPrefix(ua, "mozilla/") {
		s.add(15, "non-browser UA prefix")
	}
}

func checkHeaders(f *Fingerprint, s *Score) {
	h := f.Headers

	if h.Get("Accept") == "" {
		s.add(15, "no Accept header")
	}
	if h.Get("Accept-Language") == "" {
		s.add(15, "no Accept-Language")
	}
	if h.Get("Accept-Encoding") == "" {
		s.add(10, "no Accept-Encoding")
	}
}


var unusualPaths = []string{
	"/.env", "/.git", "/.aws", "/.ssh",
	"/wp-admin", "/wp-login.php", "/xmlrpc.php",
	"/phpmyadmin", "/admin.php", "/administrator",
	"/config.php", "/shell.php", "/.htaccess",
	"/server-status", "/actuator", "/console",
}

func checkPath(f *Fingerprint, s *Score) {
	p := strings.ToLower(f.Path)
	for _, upath := range unusualPaths {
		if strings.HasPrefix(p, upath) {
			s.add(60, "unusual path: "+upath)
			return
		}
	}
	if len(p) > 200 {
		s.add(20, "very long path")
	}
	if strings.Count(p, "..") > 0 {
		s.add(40, "path traversal attempt")
	}
}


func checkProto(f *Fingerprint, s *Score) {
	if f.Proto == "HTTP/1.0" {
		s.add(15, "HTTP/1.0 client")
	}
}

func checkMethod(f *Fingerprint, s *Score) {
	switch f.Method {
	case "TRACE", "CONNECT", "DEBUG":
		s.add(40, "unusual method: "+f.Method)
	}
}

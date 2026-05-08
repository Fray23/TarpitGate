package detector

import (
	"fmt"
	"net/http"
	// "strings"
	"testing"
)

// helper — создаёт минимальный Fingerprint, его поля можно переопределять в тестах
func TestTest(t *testing.T) {
	var a int
	a = 1
	a += 1
	// if a != 0 {
	// 	t.Error("no")
	// }
}


func newFP() *Fingerprint {
	return &Fingerprint{
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
		Headers:   http.Header{},
		Path:      "/",
		Proto:     "HTTP/1.1",
		Method:    "GET",
	}
}

func TestTest2(t *testing.T) {
	fp := newFP()
	fp.UserAgent = ""
	var s Score
	checkUserAgent(fp, &s)

	fmt.Println(s)
	var a int
	a = 1
	a += 1
	// if a != 0 {
	// 	t.Error("no")
	// }
}

// func TestCheckUserAgent(t *testing.T) {
// 	tests := []struct {
// 		name       string
// 		ua         string
// 		wantScore  int
// 		wantReason string // подстрока, которая должна встретиться в одной из причин; "" — причин быть не должно
// 	}{
// 		{"пустой UA", "", 40, "empty User-Agent"},
// 		{"только пробелы", "   ", 40, "empty User-Agent"},
// 		{"короткий UA", "Mozilla/5.0", 25 + 15, "suspiciously short"}, // <15 символов + не браузерный префикс? проверим
// 		{"curl", "curl/7.85.0", 35, "client-library UA"},
// 		{"python-requests", "python-requests/2.31.0", 35, "client-library UA"},
// 		{"бот по ключевому слову", "Googlebot/2.1 (+http://www.google.com/bot.html)", 20, "bot keyword"},
// 		{"не-браузерный префикс", "SomeRandomClient/1.0 with enough length", 15, "non-browser UA prefix"},
// 		{"нормальный браузер", "Mozilla/5.0 (X11; Linux x86_64) Gecko/20100101 Firefox/120.0", 0, ""},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			fp := newFP()
// 			fp.UserAgent = tt.ua
// 			s := &Score{}

// 			checkUserAgent(fp, s)

// 			if s.Total != tt.wantScore {
// 				t.Errorf("Total = %d, want %d (reasons: %v)", s.Total, tt.wantScore, s.Reasons)
// 			}
// 			if tt.wantReason == "" {
// 				if len(s.Reasons) != 0 {
// 					t.Errorf("ожидалось 0 причин, получили: %v", s.Reasons)
// 				}
// 				return
// 			}
// 			if !containsSubstring(s.Reasons, tt.wantReason) {
// 				t.Errorf("причины %v не содержат %q", s.Reasons, tt.wantReason)
// 			}
// 		})
// 	}
// }

// func TestCheckHeaders(t *testing.T) {
// 	tests := []struct {
// 		name      string
// 		headers   http.Header
// 		wantScore int
// 		wantCount int // сколько причин ожидаем
// 	}{
// 		{
// 			name:      "все заголовки на месте",
// 			headers:   http.Header{"Accept": {"*/*"}, "Accept-Language": {"en"}, "Accept-Encoding": {"gzip"}},
// 			wantScore: 0,
// 			wantCount: 0,
// 		},
// 		{
// 			name:      "пустые заголовки",
// 			headers:   http.Header{},
// 			wantScore: 15 + 15 + 10,
// 			wantCount: 3,
// 		},
// 		{
// 			name:      "нет только Accept-Encoding",
// 			headers:   http.Header{"Accept": {"*/*"}, "Accept-Language": {"en"}},
// 			wantScore: 10,
// 			wantCount: 1,
// 		},
// 		{
// 			name:      "нет Accept и Accept-Language",
// 			headers:   http.Header{"Accept-Encoding": {"gzip"}},
// 			wantScore: 15 + 15,
// 			wantCount: 2,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			fp := newFP()
// 			fp.Headers = tt.headers
// 			s := &Score{}

// 			checkHeaders(fp, s)

// 			if s.Total != tt.wantScore {
// 				t.Errorf("Total = %d, want %d", s.Total, tt.wantScore)
// 			}
// 			if len(s.Reasons) != tt.wantCount {
// 				t.Errorf("причин = %d, want %d (%v)", len(s.Reasons), tt.wantCount, s.Reasons)
// 			}
// 		})
// 	}
// }

// func TestCheckPath(t *testing.T) {
// 	tests := []struct {
// 		name       string
// 		path       string
// 		wantScore  int
// 		wantReason string
// 	}{
// 		{"обычный путь", "/api/v1/users", 0, ""},
// 		{"корень", "/", 0, ""},
// 		{"подозрительный .env", "/.env", 60, "unusual path"},
// 		{"wp-admin", "/wp-admin/index.php", 60, "unusual path"},
// 		{"в верхнем регистре (lower-case)", "/.ENV", 60, "unusual path"}, // checkPath применяет ToLower
// 		{"path traversal", "/files/../../etc/passwd", 40, "path traversal"},
// 		{"очень длинный путь", "/" + strings.Repeat("a", 250), 20, "very long path"},
// 		{
// 			name:      "длинный путь + traversal",
// 			path:      "/" + strings.Repeat("a", 250) + "/..",
// 			wantScore: 20 + 40,
// 			// здесь две причины — проверим обе ниже
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			fp := newFP()
// 			fp.Path = tt.path
// 			s := &Score{}

// 			checkPath(fp, s)

// 			if s.Total != tt.wantScore {
// 				t.Errorf("Total = %d, want %d (reasons: %v)", s.Total, tt.wantScore, s.Reasons)
// 			}
// 			if tt.wantReason != "" && !containsSubstring(s.Reasons, tt.wantReason) {
// 				t.Errorf("причины %v не содержат %q", s.Reasons, tt.wantReason)
// 			}
// 		})
// 	}
// }

// func TestCheckProto(t *testing.T) {
// 	tests := []struct {
// 		name      string
// 		proto     string
// 		wantScore int
// 	}{
// 		{"HTTP/1.1", "HTTP/1.1", 0},
// 		{"HTTP/2", "HTTP/2.0", 0},
// 		{"HTTP/1.0", "HTTP/1.0", 15},
// 		{"пустой", "", 0},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			fp := newFP()
// 			fp.Proto = tt.proto
// 			s := &Score{}

// 			checkProto(fp, s)

// 			if s.Total != tt.wantScore {
// 				t.Errorf("Total = %d, want %d", s.Total, tt.wantScore)
// 			}
// 		})
// 	}
// }

// func TestCheckMethod(t *testing.T) {
// 	tests := []struct {
// 		name      string
// 		method    string
// 		wantScore int
// 	}{
// 		{"GET", "GET", 0},
// 		{"POST", "POST", 0},
// 		{"PUT", "PUT", 0},
// 		{"TRACE", "TRACE", 40},
// 		{"CONNECT", "CONNECT", 40},
// 		{"DEBUG", "DEBUG", 40},
// 		// важный кейс: switch чувствителен к регистру, "trace" не сработает
// 		{"trace в нижнем регистре не триггерит", "trace", 0},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			fp := newFP()
// 			fp.Method = tt.method
// 			s := &Score{}

// 			checkMethod(fp, s)

// 			if s.Total != tt.wantScore {
// 				t.Errorf("Total = %d, want %d", s.Total, tt.wantScore)
// 			}
// 		})
// 	}
// }

// // --- утилиты ---

// func containsSubstring(reasons []string, sub string) bool {
// 	for _, r := range reasons {
// 		if strings.Contains(r, sub) {
// 			return true
// 		}
// 	}
// 	return false
// }

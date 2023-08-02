package httputil

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"
)

var testString = "huoys http client test"

func TestRequest(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(serveHTTP))
	wg := sync.WaitGroup{}
	wg.Add(10)
	for range [10]struct{}{} {
		go func() {
			code, res, err := Request(http.MethodPost, svr.URL, "", nil)
			if err != nil {
				t.Error("Request test error:", err)
			}
			if string(res) != testString || code != http.StatusOK {
				t.Errorf("Request test error.code=%d;res=%s", code, string(res))
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func serveHTTP(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second * 3)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(testString))
}

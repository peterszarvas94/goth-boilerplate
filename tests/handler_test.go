package tests

import (
	"goth/internal/server"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	s := &server.Server{}
	server := httptest.NewServer(http.HandlerFunc(s.HelloWorldHandler))
	defer server.Close()
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("e    rror making request to server. Err: %v", err)
	}
	defer resp.Body.Close()
	// Assertions
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", resp.Status)
	}
	expected := "<html lang=\"en\"><head><title>Index</title><script src=\"https://unpkg.com/htmx.org@1.9.10\"></script><meta name=\"description\" content=\"Welcome to the index page!\"></head><body><div><h1 class=\"text-4xl\">Index</h1><p>Welcome to the index page!</p></div></body></html>"
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("error reading response body. Err: %v", err)
	}
	if expected != string(body) {
		t.Errorf("expected response body to be %v; got %v", expected, string(body))
	}
}

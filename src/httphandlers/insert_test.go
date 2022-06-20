package httphandler

import (
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInsertRow(t *testing.T) {
	tests := []struct {
		name  string
		url   string
		query string
	}{
		{name: "Insert Data 1", url: "http://localhost:9000/insertrow", query: "title=Test 1"},
		{name: "Insert Data 2", url: "http://localhost:9000/insertrow", query: "title=Test 2"},
		{name: "Insert Data 3", url: "http://localhost:9000/insertrow", query: "title=Test 3"},
		{name: "Insert Data 4", url: "http://localhost:9000/insertrow", query: "title=Test 4"},
	}

	l, err := net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	server := httptest.NewUnstartedServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				InsertRow(w, r)
			},
		),
	)
	defer server.Close()
	server.Listener.Close()
	server.Listener = l
	server.Start()

	for _, d := range tests {
		t.Run(d.name, func(t *testing.T) {
			resp, err := http.Get(d.url + "?" + d.query)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			defer resp.Body.Close()
		})
	}
}

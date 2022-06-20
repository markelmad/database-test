package httphandler

import (
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpdateReadCount(t *testing.T) {
	tests := []struct {
		name  string
		url   string
		query string
	}{
		{name: "Update Data 1", url: "http://localhost:9000/updatereadcount", query: "id=1&readcount=4"},
		{name: "Update Data 2", url: "http://localhost:9000/updatereadcount", query: "id=2&readcount=5"},
		{name: "Update Data 3", url: "http://localhost:9000/updatereadcount", query: "id=3&readcount=20"},
		{name: "Update Data 4", url: "http://localhost:9000/updatereadcount", query: "id=4&readcount=60"},
	}

	l, err := net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	server := httptest.NewUnstartedServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				UpdateReadCount(w, r)
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

package httphandler

import (
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDeleteRowByID(t *testing.T) {

	tests := []struct {
		name  string
		url   string
		query string
	}{
		{name: "Delete ID 1", url: "http://localhost:9000/updatereadcount", query: "id=1"},
		{name: "Delete ID 100", url: "http://localhost:9000/updatereadcount", query: "id=100"},
		{name: "Delete ID A", url: "http://localhost:9000/updatereadcount", query: "id=A"},
		{name: "Delete ID 002", url: "http://localhost:9000/updatereadcount", query: "id=002"},
	}

	l, err := net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	server := httptest.NewUnstartedServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				DeleteRowByID(w, r)
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

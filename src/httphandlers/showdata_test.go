package httphandler

import (
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestShowAll(t *testing.T) {

	l, err := net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	server := httptest.NewUnstartedServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				ShowAll(w, r)
			},
		),
	)
	defer server.Close()
	server.Listener.Close()
	server.Listener = l
	server.Start()

	t.Run("Show All Data", func(t *testing.T) {
		resp, err := http.Get("http://localhost:9000/showdb")
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
		defer resp.Body.Close()
	})
}

func TestShowDataByCategory(t *testing.T) {

	tests := []struct {
		name     string
		url      string
		cateogry string
	}{
		{name: "Show by category comedy", url: "http://localhost:9000/showtitlesbycategory/", cateogry: "comedy"},
		{name: "Show by category tragedy", url: "http://localhost:9000/showtitlesbycategory/", cateogry: "tragedy"},
		{name: "Show by category history", url: "http://localhost:9000/showtitlesbycategory/", cateogry: "history"},
		{name: "Show by category poetry", url: "http://localhost:9000/showtitlesbycategory/", cateogry: "poetry"},
	}

	l, err := net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	server := httptest.NewUnstartedServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				ShowDataByCategory(w, r)
			},
		),
	)
	defer server.Close()
	server.Listener.Close()
	server.Listener = l
	server.Start()

	for _, d := range tests {
		t.Run(d.name, func(t *testing.T) {
			resp, err := http.Get(d.url + d.cateogry)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			defer resp.Body.Close()
		})
	}
}

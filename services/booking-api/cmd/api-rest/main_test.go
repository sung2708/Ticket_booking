package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootRouteReturnsHelloWorld(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	newRouter().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	if rec.Body.String() != "Hello World!" {
		t.Fatalf("expected body %q, got %q", "Hello World!", rec.Body.String())
	}
}

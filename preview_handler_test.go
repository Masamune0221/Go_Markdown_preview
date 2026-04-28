package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPreviewHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/", nil)
	previewHandler(w, r)
	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}
	// goldmarkの仕様で末尾に改行が入るため、\n を追加して比較する
	expected := "<h1>Hello World!!</h1>\n"
	if w.Body.String() != expected {
		t.Errorf("Expected body %q, got %q", expected, w.Body.String())
	}
}

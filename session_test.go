package gflights_test

import (
	"io"
	"net/http"
	"os"
	"testing"
)

func getReadCloser(t *testing.T, filePath string) io.ReadCloser {
	t.Helper()
	f, err := os.Open(filePath)
	if err != nil {
		t.Fatalf("failed to open test file %s: %v", filePath, err)
	}
	return io.NopCloser(f)
}

func getSetUpResponse() *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Header: http.Header{
			"Set-Cookie": []string{},
		},
	}
}

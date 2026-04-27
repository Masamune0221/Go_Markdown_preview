package main

import (
	"errors"
	"testing"
)

func TestErrorHandler(t *testing.T) {
	err := errors.New("エラーが発生しました")
	handler := NewErrorHandler(err, "テストです")
	handler.Handle(nil, nil)
}

package main

import (
	"fmt"
	"testing"
)

func TestStep1(t *testing.T) {
	// Markdownフォーマットの文字列
	input := []byte(`# TEST`)

	// MarkdownからHTMLに変換
	output, err := ConvertHTML(input)
	if err != nil {
		errorHandler := NewErrorHandler(err, "Failed to convert markdown to html")
		errorHandler.Handle()
	}

	// HTML文字列を表示
	fmt.Println(output.String())

}

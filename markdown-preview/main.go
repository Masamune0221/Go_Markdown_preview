package main

import (
	"bytes"
	"fmt"

	"github.com/yuin/goldmark"
)

func main() {
	// Markdownフォーマットの文字列
	input := []byte(`# Hello World!!`)

	// 結果を書き込むバッファ
	var output bytes.Buffer

	// Goldmarkのパーサーを初期化
	md := goldmark.New()

	// MarkdownからHTMLに変換
	if err := md.Convert(input, &output); err != nil {
		errorHandler := NewErrorHandler(err, "MarkdownからHTMLへの変換に失敗しました")
		errorHandler.Handle()
	}
	// HTML文字列を表示
	fmt.Println(output.String())

}

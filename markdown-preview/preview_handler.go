package main

import (
	"fmt"
	"net/http"
)

/**
 * プレビューハンドラー
 * @param w http.ResponseWriter: レスポンスライター
 * @param r *http.Request: リクエスト
 **/
func previewHandler(w http.ResponseWriter, r *http.Request) {
	// Markdownフォーマットの文字列
	input := []byte(`### Hello World!!
### Hello World!!
### Hello World!!`)
	// MarkdownをHTMLに変換
	output, err := ConvertHTML(input)
	// エラーハンドリング
	if err != nil {
		errHandler := NewErrorHandler(err, "Failed to convert markdown to html")
		errHandler.Handle(w, r)
		return
	}
	// レスポンスヘッダーを設定
	w.Header().Set(
		"Content-Type", "text/html; charset=utf-8",
	)
	// %s を output.String() で置き換えて、HTMLを出力
	fmt.Fprintf(w, htmlTemplate, output.String())
}

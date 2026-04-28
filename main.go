package main

import (
	"fmt"
	"net/http"
)

func main() {

	// ファイルが変更された時の処理（コールバック）を渡す
	go watchFile("target/sample.md", func() {
		broadcast()
	})
	// HTTPサーバーを設定
	// 静的ファイルの配信を設定
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	// メインハンドラーを設定
	http.HandleFunc("/", previewHandler)
	http.HandleFunc("/ws", wsHandler)
	fmt.Println("Server is running on http://localhost:8080")
	// HTTPサーバーを起動
	err := http.ListenAndServe(":8080", nil)
	// エラーハンドリング
	if err != nil {
		errHandler := NewErrorHandler(err, "Failed to start server")
		errHandler.Handle(nil, nil)
		return
	}
}

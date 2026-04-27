package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	// WebSocket接続を確立
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		errHandler := NewErrorHandler(err, "Failed to upgrade connection")
		errHandler.Handle(w, r)
		return
	}
	// 終了時に接続を閉じる
	defer conn.Close()
	// クライアントを登録
	clients[conn] = true

	// メッセージ受信ループ
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			delete(clients, conn)
			break
		}
		// 受信したメッセージを全クライアントにブロードキャスト（ここでは単純にループ）
		for client := range clients {
			if err := client.WriteMessage(websocket.TextMessage, message); err != nil {
				delete(clients, client)
			}
		}
	}
}

func broadcast() {
	for client := range clients {
		err := client.WriteMessage(websocket.TextMessage, []byte("update"))
		if err != nil {
			errHandler := NewErrorHandler(err, "Failed to broadcast message")
			errHandler.Handle(nil, nil)
			delete(clients, client)
		}
	}
}

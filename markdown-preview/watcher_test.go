package main

import (
	"os"
	"testing"
	"time"
)

func TestWatchFile(t *testing.T) {
	// 1. テスト用の一時ファイルを作成する
	tmpFile, err := os.CreateTemp("", "test_watch_*.md")
	if err != nil {
		t.Fatalf("一時ファイルの作成に失敗しました: %v", err)
	}
	// テストが終わったらファイルを削除して閉じる
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	// 2. 変更を検知したかを記録するチャネル（通信の筒）を作る
	changed := make(chan bool)

	// 3. バックグラウンドで監視をスタート
	// 変更があったら changed チャネルに true を送信する
	go watchFile(tmpFile.Name(), func() {
		changed <- true
	})

	// Watcherが確実に起動するまで少しだけ待つ
	time.Sleep(100 * time.Millisecond)

	// 4. ファイルを更新（書き込み）する
	if _, err := tmpFile.WriteString("hello test"); err != nil {
		t.Fatalf("一時ファイルへの書き込みに失敗しました: %v", err)
	}

	// 5. 変更イベントがちゃんと来るか（またはタイムアウトするか）をテストする
	select {
	case <-changed:
		// 無事に検知されたのでテストPASS！
		return
	case <-time.After(2 * time.Second):
		// 2秒待っても検知されなかったらFAIL
		t.Errorf("ファイルの変更が検知されませんでした（タイムアウト）")
	}
}

package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

// onChange はファイルが保存(更新)された時に呼ばれる関数です
func watchFile(targetPath string, onChange func()) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		errHandler := NewErrorHandler(err, "Failed to create file watcher")
		errHandler.Handle(nil, nil)
		return
	}
	defer watcher.Close()

	//
	err = watcher.Add(targetPath)
	if err != nil {
		errHandler := NewErrorHandler(err, "Failed to add target path to watcher")
		errHandler.Handle(nil, nil)
		return
	}

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			log.Println("event:", event)
			// ファイルが書き込まれた（保存された）ときのイベントを検知
			if event.Has(fsnotify.Write) {
				onChange()
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("error:", err)
		}
	}
}

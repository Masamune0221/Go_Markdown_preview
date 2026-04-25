package main

import (
	"log"
)

// エラーハンドラーの構造体
type ErrorHandler struct {
	errMessage string
	err        error
}

/**
 * エラーハンドリング関数
 * エラーが発生した場合、指定されたメッセージと共にログを出力し、プログラムを終了する
 **/
func (e ErrorHandler) Handle() {
	log.Fatalf("ERROR at %s: %v\n", e.errMessage, e.err)
}

/**
 * エラーハンドリング関数を初期化する
 * @param err error: エラー情報
 * @param message string: エラーメッセージ
 * @return ErrorHandler: エラーハンドラーのポインター
 **/
func NewErrorHandler(err error, message string) ErrorHandler {
	return ErrorHandler{
		err:        err,
		errMessage: message,
	}
}

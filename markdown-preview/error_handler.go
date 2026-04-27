package main

import (
	"log"
	"net/http"
)

// エラーハンドラーの構造体
type ErrorHandler struct {
	errMessage string
	err        error
}

/**
 * エラーハンドリング関数
 * エラーが発生した場合、指定されたメッセージと共にログを出力する
 * @param w http.ResponseWriter: レスポンスライター
 * @param r *http.Request: リクエスト
 **/
func (e ErrorHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if w != nil && r != nil {
		http.Error(w, e.errMessage, http.StatusInternalServerError)
	}
	log.Printf("ERROR at %s: %v\n", e.errMessage, e.err)
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

package main

import (
	"bytes"

	"github.com/yuin/goldmark"
)

/**
 * MarkdownをHTMLに変換する
 * @param input []byte: Markdownフォーマットの文字列
 * @return bytes.Buffer: HTMLフォーマットの文字列
 */
func ConvertHTML(input []byte) (bytes.Buffer, error) {
	var output bytes.Buffer
	md := goldmark.New()

	// MarkdownからHTMLへの変換
	if err := md.Convert(input, &output); err != nil {
		return output, err
	}
	return output, nil
}

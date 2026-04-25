package main

import "testing"

func TestConvertHTML(t *testing.T) {
	input := []byte(`# Hello World!!`)
	output, err := ConvertHTML(input)

	// エラーがないことを確認
	if err != nil {
		t.Fatalf("予期せぬエラーが発生しました: %v", err)
	}

	// goldmarkはデフォルトで末尾に改行を入れるので、改行を含めて比較する
	expected := "<h1>Hello World!!</h1>\n"
	if output.String() != expected {
		t.Errorf("ConvertHTML() = %q, want %q", output.String(), expected)
	}
}

# 📄 Markdown Preview Tool

Go言語で作ったローカルのMarkdownファイルをブラウザでリアルタイムにプレビューできるツールです。
ファイルを保存するたびに、ブラウザが自動でリロードします（ホットリロード機能）。

## 📸 機能

- **Markdownのリアルタイムプレビュー** — `target/` フォルダ内の `.md` ファイルをブラウザに表示
- **ホットリロード** — ファイルを保存した瞬間、ブラウザが自動で更新される（WebSocket使用）
- **ダークモード対応** — TailwindCSS v4 でスタイリング
- **シンプルなHTTPサーバー** — Go標準ライブラリの `net/http` のみでサーバーを構築

## 🛠️ 使用技術

| カテゴリ | 技術 |
|---|---|
| 言語 | Go |
| Markdownパーサー | [goldmark](https://github.com/yuin/goldmark) |
| ファイル監視 | [fsnotify](https://github.com/fsnotify/fsnotify) |
| WebSocket | [gorilla/websocket](https://github.com/gorilla/websocket) |
| スタイリング | [Tailwind CSS v4](https://tailwindcss.com/) |

## 📂 ファイル構成

```
markdown-preview/
├── main.go                  # エントリーポイント。サーバー起動・ファイル監視の起動
├── preview_handler.go       # HTTPハンドラー。Markdownを読んでHTMLを返す
├── convert_html.go          # MarkdownをHTMLに変換する関数
├── watcher.go               # fsnotify を使ったファイル監視
├── ws_handler.go            # WebSocketの接続管理・ブラウザへのリロード通知
├── html_template.go         # ブラウザに返すHTMLテンプレート
├── error_handler.go         # 共通エラーハンドラー
├── input.css                # TailwindCSSのエントリーポイント
├── static/css/style.css     # Tailwindがビルドした完成品CSS
└── target/
    └── sample.md            # プレビューするMarkdownファイル（ここを編集する）
```

## 🚀 セットアップと起動方法

### 必要なもの
- Go 1.20 以上
- Node.js / npm

### 1. 依存パッケージのインストール

```bash
# Go のライブラリをインストール
go mod tidy

# Node.js のパッケージをインストール（Tailwind CSS）
npm install
```

### 2. TailwindCSS のビルド（別ターミナルで起動）

```bash
npx tailwindcss -i input.css -o static/css/style.css --watch
```

### 3. Goサーバーの起動

```bash
go run .
```

サーバーが起動したら、ブラウザで以下のURLにアクセスしてください。

```
http://localhost:8080
```

### 4. プレビューしてみよう！

`target/sample.md` をエディタで編集して保存すると、ブラウザが自動でリロードされます！✨

## 🧪 テスト

```bash
go test -v ./...
```

## 📖 開発ステップ（学習ロードマップ）

このプロジェクトは以下の4ステップで作成しました。

1. **Step 1: HTML変換機能** — `goldmark` でMarkdown→HTML変換関数を実装
2. **Step 2: HTTPサーバー構築** — `net/http` でWebサーバーを立ち上げ、HTML配信
3. **Step 3: ファイル監視** — `fsnotify` で `.md` ファイルの保存イベントを検知
4. **Step 4: WebSocketによるライブリロード** — ファイル変更をWebSocketでブラウザに通知し、自動リロードを実現

## 💡 アピールポイント

- 外部フレームワークに依存せず、Goの標準ライブラリ（`net/http`）を使いこなしている
- `goroutine` を活用し、サーバー待機とファイル監視を並行して処理している
- 役割ごとにファイルを分割（関心の分離）し、保守しやすい構成になっている
- テストファイルを各機能に対して作成し、動作を担保している

package main

// HTMLのテンプレートを文字列で定義する
const htmlTemplate = `
<!DOCTYPE html>
<html lang="ja" class="dark">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Markdown Preview</title>
	<!-- さっき作ったstaticフォルダのCSSを読み込む -->
	<link rel="stylesheet" href="/static/css/style.css">
</head>
<!-- 背景と文字色をダークモードに対応、ふわっと色が変わるように transition を追加 -->
<body class="bg-gray-100 dark:bg-slate-900 text-gray-900 dark:text-gray-100 p-8 transition-colors duration-300">
	<div class="max-w-3xl mx-auto bg-white dark:bg-slate-800 p-10 rounded-xl shadow-lg prose transition-colors duration-300">
		%s
	</div>
</body>
</html>
`

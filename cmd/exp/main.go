package main

import (
	"fmt"
	"os"

	"github.com/russross/blackfriday/v2"
)

func main() {
	// 读取 Markdown 文件内容
	mdFilePath := "./notes/url_query.md"
	mdContent, err := os.ReadFile(mdFilePath)
	if err != nil {
		fmt.Printf("Failed to read file: %v\n", err)
		return
	}

	// 调用 mdToHTML 函数将 Markdown 转换为 HTML
	htmlContent := mdToHTML(mdContent)

	// 输出 HTML 内容
	fmt.Println(string(htmlContent))
}

func mdToHTML(md []byte) []byte {
	// 使用 blackfriday 进行 Markdown 到 HTML 的转换
	htmlContent := blackfriday.Run(md)

	return htmlContent
}

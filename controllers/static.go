package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/russross/blackfriday/v2"
)

// 闭包
func StaticHandler(tpl Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}

}

func FQA(tpl Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   template.HTML
	}{
		{
			Question: "Is there a free version",
			Answer:   "Yes!",
		},
		{
			Question: "Is there a free version",
			Answer:   "No!",
		},
		{
			Question: "baidu url",
			Answer:   `<a href="https://www.w3schools.com">Visit W3Schools.com!</a>`,
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}

// documentation md to html
func mdToHTML(md []byte) []byte {
	// 使用 blackfriday 进行 Markdown 到 HTML 的转换
	htmlContent := blackfriday.Run(md)

	return htmlContent
}
func htmlContent(path string) ([]byte, error) {
	// 读取 Markdown 文件内容
	mdFilePath := "./notes/url_query.md"
	mdContent, err := os.ReadFile(mdFilePath)
	if err != nil {
		fmt.Printf("Failed to read file: %v\n", err)
		return nil, err
	}

	// 调用 mdToHTML 函数将 Markdown 转换为 HTML
	htmlContent := mdToHTML(mdContent)

	// 输出 HTML 内容
	return htmlContent, nil
}

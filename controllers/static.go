package controllers

import (
	"html/template"
	"net/http"

	"github.com/baimiyishu13/lenslocked/views"
)

// 闭包
func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}

}

func FQA(tpl views.Template) http.HandlerFunc {
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
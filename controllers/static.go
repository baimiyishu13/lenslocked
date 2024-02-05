package controllers

import (
	"net/http"

	"github.com/baimiyishu13/lenslocked/views"
)

// 闭包
func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}

}

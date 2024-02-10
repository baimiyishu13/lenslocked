package controllers

import (
	"net/http"

	"github.com/baimiyishu13/lenslocked/views"
)

type Users struct {
	Templates struct {
		New views.Template
	}
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	//需要一个视图渲染
	u.Templates.New.Execute(w, nil)
}

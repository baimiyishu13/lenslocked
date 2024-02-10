package controllers

import (
	"net/http"
)

type Users struct {
	Templates struct {
		New Template
	}
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	//需要一个视图渲染
	u.Templates.New.Execute(w, nil)
}

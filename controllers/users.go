package controllers

import (
	"fmt"
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

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w, "Email: ", r.FormValue("email"))
	fmt.Fprintln(w, "Password: ", r.FormValue("password"))
	fmt.Fprintln(w, "Repeat-password: ", r.FormValue("repeat-password"))
}

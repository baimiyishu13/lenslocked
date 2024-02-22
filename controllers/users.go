package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/baimiyishu13/lenslocked/models"
)

type Users struct {
	Templates struct {
		New    Template
		SignIN Template
	}
	UserService *models.UserService
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	// 实现预填充
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	//需要一个视图渲染
	u.Templates.New.Execute(w, data)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	email := r.FormValue("email")
	password := r.FormValue("password")
	user, err := u.UserService.Create(email, password)
	if err != nil {
		log.Fatal("Error creating user")
	}
	fmt.Fprintf(w, "User created: %v\n", user)
}

func (u Users) SignIN(w http.ResponseWriter, r *http.Request) {
	// 实现预填充
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	//需要一个视图渲染
	u.Templates.SignIN.Execute(w, data)
}

func (u Users) ProccesSignIN(w http.ResponseWriter, r *http.Request) {
	// 实现预填充
	var data struct {
		Email    string
		Password string
	}
	data.Email = r.FormValue("email")
	data.Password = r.FormValue("password")
	user, err := u.UserService.Authenticate(data.Email, data.Password)
	if err != nil {
		fmt.Print(err)
		http.Error(w, "Something went wrong", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "User authenticated: %v\n", user)
}

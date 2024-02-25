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

	cookie := http.Cookie{
		Name:     "email",
		Value:    user.Email,
		Path:     "/",
		HttpOnly: true, // 防止 XSS跨站脚本，只允许Http浏览器请求
	}
	http.SetCookie(w, &cookie)
	fmt.Fprintf(w, "User authenticated: %v\n", user)
}

// 接受网络请求并且打印数据
func (u Users) CurrentUsers(w http.ResponseWriter, r *http.Request) {
	email, err := r.Cookie("email")
	if err != nil {
		fmt.Fprintf(w, "The email cookie could not be read")
		return
	}
	fmt.Fprintf(w, "email cookie: %s\n", email.Value)
	fmt.Fprintf(w, "Headers: %+v\n", r.Header)
}

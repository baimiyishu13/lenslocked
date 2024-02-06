// 定义代码所属的包
package main

//
import (
	"fmt"
	"net/http"

	"github.com/baimiyishu13/lenslocked/controllers"
	"github.com/baimiyishu13/lenslocked/templates"
	"github.com/baimiyishu13/lenslocked/views"
	"github.com/go-chi/chi/v5"
)

// chi
func main() {
	r := chi.NewRouter()
	// parse tpl

	r.Get("/", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "home.gohtml", "layout-parts.gohtml"))))

	r.Get("/contact", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "contact.gohtml", "layout-parts.gohtml"))))

	r.Get("/fqa", controllers.FQA(
		views.Must(views.ParseFS(templates.FS, "fqa.gohtml"))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on 3000 ...")
	http.ListenAndServe(":3000", r)
}

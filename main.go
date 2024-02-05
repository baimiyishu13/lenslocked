// 定义代码所属的包
package main

//
import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/baimiyishu13/lenslocked/controllers"
	"github.com/baimiyishu13/lenslocked/views"
	"github.com/go-chi/chi/v5"
)

// chi
func main() {
	r := chi.NewRouter()
	// parse tpl
	tpl, err := views.Parse(filepath.Join("templates", "home.gohtml"))
	if err != nil {
		panic(err)
	}

	r.Get("/", controllers.StaticHandler(tpl))

	tpl, err = views.Parse(filepath.Join("templates", "contact.gohtml"))
	if err != nil {
		panic(err)
	}

	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl, err = views.Parse(filepath.Join("templates", "fqa.gohtml"))
	if err != nil {
		panic(err)
	}

	r.Get("/fqa", controllers.StaticHandler(tpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on 3000 ...")
	http.ListenAndServe(":3000", r)
}

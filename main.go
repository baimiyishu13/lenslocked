// 定义代码所属的包
package main

//
import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/baimiyishu13/lenslocked/views"
	"github.com/go-chi/chi/v5"
)

// 解析&执行 gohtml文件
func executeTemplate(w http.ResponseWriter, filepath string) {
	t, err := views.Parse(filepath)
	if err != nil {
		log.Printf("Parsing template: %v", err)
		http.Error(w, "The wars an error parsing the template", http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)

}

// 基本的 web 应用程序
// 主 界面
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "gohtml")
	executeTemplate(w, tplPath)
}

// 联系 界面
func contactHandler(w http.ResponseWriter, r *http.Request) {
	tplpath := filepath.Join("templates", "contact.gohtml")
	executeTemplate(w, tplpath)
}

// FQA 界面
func fqaHandler(w http.ResponseWriter, r *http.Request) {
	tplpath := filepath.Join("templates", "fqa.gohtml")
	executeTemplate(w, tplpath)
}

// chi
func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/fqa", fqaHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		// w.Write([]byte("404"))
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on 3000 ...")
	http.ListenAndServe(":3000", r)
}

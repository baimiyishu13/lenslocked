// 定义代码所属的包
package main

//
import (
	"fmt"
	"log"
	"net/http"

	"github.com/baimiyishu13/lenslocked/controllers"
	"github.com/baimiyishu13/lenslocked/models"
	"github.com/baimiyishu13/lenslocked/templates"
	"github.com/baimiyishu13/lenslocked/views"
	"github.com/go-chi/chi/v5"
)

// chi
func main() {
	r := chi.NewRouter()
	// parse tpl

	r.Get("/", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS,
			"home.gohtml", "tailwind.gohtml",
		))))

	r.Get("/contact", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS,
			"contact.gohtml", "tailwind.gohtml",
		))))

	r.Get("/documentation", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS,
			"documentation.gohtml", "tailwind.gohtml",
		))))

	// User
	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}
	fmt.Println("open database connection")

	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal("failed to ping database: ", err)
	}

	userService := models.UserService{
		DB: db,
	}

	_, err = userService.DB.Exec(`
	   Create TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
        email TEXT UNIQUE NOT NULL,
        passwordHash TEXT UNIQUE NOT NULL
	   );
	`)
	if err != nil {
		log.Fatal("failed to create users table: ", err)
	}

	userC := controllers.Users{
		UserService: &userService, //TODO: set this
	}
	userC.Templates.New = views.Must(views.ParseFS(
		templates.FS,
		"signup.gohtml", "tailwind.gohtml",
	))
	userC.Templates.SignIN = views.Must(views.ParseFS(
		templates.FS,
		"signin.gohtml", "tailwind.gohtml",
	))

	r.Get("/signup", userC.New)
	r.Post("/users", userC.Create)
	r.Get("/signin", userC.SignIN)
	r.Post("/signin", userC.ProccesSignIN)

	// FQA
	r.Get("/fqa", controllers.FQA(
		views.Must(views.ParseFS(templates.FS,
			"fqa.gohtml", "tailwind.gohtml",
		))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on 3000 ...")
	http.ListenAndServe(":3000", r)

}

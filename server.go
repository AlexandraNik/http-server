package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	e := echo.New()

	e.GET("/hello", hello)
	e.GET("/users/:id", getUser)
	e.POST("/users", saveUser)

	s := http.Server{
		Addr:    ":8080",
		Handler: e,
		//ReadTimeout: 30 * time.Second, // customize http.Server timeouts
	}

	// http.HandleFunc("/hello", Tmp)

	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

var db *sql.DB

func GetDB() *sql.DB {
	var err error

	if db == nil {
		connStr := "user=postgres dbname=postgres"
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			panic(err)
		}
	}

	return db
}

// without echo
// func Tmp(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case http.MethodGet:
// 		TmpGet(w, r)
// 	case http.MethodPost:
// 		TmpPost(w, r)
// 	default:
// 		w.WriteHeader(http.StatusMethodNotAllowed)
// 	}
// }

// func TmpGet(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Hello, World!")
// }

// func TmpPost(w http.ResponseWriter, r *http.Request) {
// 	output, err := io.ReadAll(r.Body)
// 	if err != nil || len(output) == 0 {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	fmt.Fprintf(w, "Pulki + %v", string(output))
// }

//for echo
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

//for echo
func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

//for echo
func saveUser(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	db := GetDB()
	_, err := db.Exec("INSERT INTO postgres (username, email) VALUES($1, $2)", name, email)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "name:"+name+", email:"+email)
}
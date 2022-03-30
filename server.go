package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())

	e.GET("/hello", Hello)
	e.GET("/users/:id", GetUser)
	e.POST("/users", SaveUser)
	e.PUT("/users", EditUser)
	e.DELETE("/users", DeleteUser)

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

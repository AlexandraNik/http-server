package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1"
	dbname   = "postgres"
)

var db *sql.DB

func GetDB() *sql.DB {
	var err error

	if db == nil {
		connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			panic(err)
		}
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	return db
}

//for echo
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

//for echo
func SaveUser(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	db := GetDB()
	_, err := db.Exec("INSERT INTO accounts (username, email) VALUES($1, $2)", name, email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, "User name:"+name+", email:"+email+" is created")
}

//for echo
func GetUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")

	db := GetDB()
	rows, err := db.Query("SELECT username FROM accounts WHERE user_id = $1", id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()
	names := make([]string, 0)
	for rows.Next() {
		var username string
		err = rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		names = append(names, username)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return c.String(http.StatusOK, names[0])
}

func EditUser(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	db := GetDB()

	sqlUpdate := `
	UPDATE accounts
	SET email = $1
	WHERE username = $2;`
	_, err := db.Exec(sqlUpdate, email, name)
	if err != nil {
		panic(err)
	}
	return c.String(http.StatusOK, "User "+name+", email:"+email+" is updated")
}

func DeleteUser(c echo.Context) error {
	name := c.FormValue("name")
	db := GetDB()

	sqlDelete := `
	DELETE FROM accounts WHERE username = $1;`
	_, err := db.Exec(sqlDelete, name)
	if err != nil {
		panic(err)
	}

	return c.String(http.StatusOK, "User "+name+" is deleted")
}

package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

/* Lorama */
type App struct {
	folderpath string
	db         *sql.DB
	err        error
}

/* Lorama */
type User struct {
	Uid                  int
	Username, Departname string
	Created              string
}

func (App *App) connect() {
	App.db, App.err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golearn1?charset=utf8")
	checkErr(App.err)
}

func (App *App) indexHander(w http.ResponseWriter, r *http.Request) {
	rows, err := App.db.Query("SELECT * FROM userinfo")
	checkErr(err)
	defer rows.Close()

	t, _ := template.ParseFiles(App.folderpath + "/list.gtpl")

	users := make([]User, 0) // define empty collection of users

	for rows.Next() {
		var uid int
		var username string
		var departname string
		var created string
		err := rows.Scan(&uid, &username, &departname, &created)
		checkErr(err)
		users = append(users, User{uid, username, departname, created})
	}

	fmt.Println(users)
	t.Execute(w, users)

}

func (App *App) viewHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	created := ""
	user := new(User)
	err := App.db.QueryRow("SELECT * FROM userinfo WHERE uid=?", id).Scan(&user.Uid, &user.Username, &user.Departname, &created)
	checkErr(err)
	fmt.Println(user)

	t, _ := template.ParseFiles(App.folderpath + "/view.gtpl")
	t.Execute(w, user)
}

func (App *App) createHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles(App.folderpath + "/form.gtpl")
		t.Execute(w, nil)
	} else {
		// POST goes here
		r.ParseForm()
		stmt, err := App.db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
		checkErr(err)

		res, err := stmt.Exec(r.Form["username"][0], r.Form["departname"][0], time.Now())
		checkErr(err)

		id, err := res.LastInsertId()
		checkErr(err)

		fmt.Println(id)

		http.Redirect(w, r, "/view?id="+fmt.Sprintf("%d", id), 301)
	}

}

func (App *App) deleteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	stmt, err := App.db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err := stmt.Exec(id)
	checkErr(err)
	fmt.Println(res)
	http.Redirect(w, r, "/", 301)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	app := new(App)
	app.folderpath, _ = os.Getwd()
	app.connect()

	http.HandleFunc("/", app.indexHander) // setting router rule
	http.HandleFunc("/create", app.createHandler)
	http.HandleFunc("/view", app.viewHandler)
	http.HandleFunc("/delete", app.deleteHandler)

	err2 := http.ListenAndServe(":9090", nil) // setting listening port
	checkErr(err2)
}

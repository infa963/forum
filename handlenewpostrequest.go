package forum

import (
	"database/sql"
	"log"
	"net/http"
	"strings"
	"text/template"

	_ "github.com/mattn/go-sqlite3"
)

func HandleNewPostRequest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//Store specific information from that post that was received
	newPostTitle := r.Form["title"][0]
	newPostContent := r.Form["content"][0]
	newPostCategories := r.Form["category"][0]

	forumDatabase, err := sql.Open("sqlite3", "./forum-database.db")
	if err != nil {
		log.Println(err.Error())
	}
	defer forumDatabase.Close()
	createPostTable(forumDatabase)
	//I changed the err into '_' to remove the orange error as it states that it isn't being used
	tpl, _ := template.ParseFiles("../static/templates/index.gohtml")
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			tpl.Execute(w, nil)
			return
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	username := strings.SplitN(c.Value, "-", 2)[0]

	//Adds relevant information of the new post into the database
	insertNewPost(forumDatabase, username, newPostTitle, newPostContent, newPostCategories)
}
func insertNewPost(db *sql.DB, username string, title string, content string, categories string) {
	insertNewPostSQL := `INSERT INTO post (username, title, created_date, content, categories) VALUES (?, ?, CURRENT_TIMESTAMP, ?, ?)`
	statement, err := db.Prepare(insertNewPostSQL)

	if err != nil {
		log.Println(err.Error())
	}

	_, err = statement.Exec(username, title, content, categories)
	if err != nil {
		log.Println(err.Error())
	}
}

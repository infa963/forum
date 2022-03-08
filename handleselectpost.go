package forum

import (
	"database/sql"
	"log"
	"net/http"
	"strings"
	"text/template"
)

func HandleSelectPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	date := r.Form["date"][0]
	username := r.Form["username"][0]
	date = strings.Join(strings.Split(strings.Join(strings.Split(date, "T"), " "), "Z"), "")
	post := GetPost(date, username)
	likes := GetLikes(post.PostId)
	comments := GetComments(post.PostId)
	UserSession.Comments = nil
	UserSession.Likes = likes
	for _, comment := range comments {
		UserSession.Comments = append(UserSession.Comments, comment)
	}
	log.Println("COMMENTS:", UserSession.Comments)
	UserSession.SelectedPost = post
	tpl, err := template.ParseFiles("../static/templates/index.gohtml")
	if err != nil {
		log.Println(err.Error())
	}
	tpl.Execute(w, UserSession)
}
func GetLikes(postID string) string {
	forumDatabase, err := sql.Open("sqlite3", "./forum-database.db")
	if err != nil {
		log.Println(err.Error())
	}
	defer forumDatabase.Close()
	getLikesSQL := `SELECT (SELECT COUNT(like) from like WHERE post_id = ? and like = 1) - (SELECT COUNT(like) from like WHERE post_id = ? AND like = 0);`
	statement, err := forumDatabase.Prepare(getLikesSQL)
	if err != nil {
		log.Println(err.Error())
	}
	var likes string
	err = statement.QueryRow(postID, postID).Scan(&likes)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Likes: ", likes)
	return likes
}

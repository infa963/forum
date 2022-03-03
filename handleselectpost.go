package forum

import (
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
	UserSession.SelectedPost = post
	tpl, err := template.ParseFiles("../static/templates/index.gohtml")
	if err != nil {
		log.Println(err.Error())
	}
	tpl.Execute(w, UserSession)
}

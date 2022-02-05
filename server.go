package forum

import "net/http"

func Server() {
	http.Handle("/index.css", http.FileServer(http.Dir("../static")))
	fileServer := http.FileServer(http.Dir("../static/img"))
	http.Handle("/img/", http.StripPrefix("/img", fileServer))
	http.HandleFunc("/", HandleDefault)
	http.HandleFunc("/signin", HandlePostRequest)
	http.ListenAndServe(":8080", nil)
}

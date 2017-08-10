package web

import (
	chandler "common/handler"
	"html/template"
	"net/http"
	"web/handler"

	"github.com/gorilla/mux"
)

func init() {
	r := mux.NewRouter()

	r.HandleFunc("/", chandler.DefaultWrapper(index)).Methods("GET")
	r.HandleFunc("/comment", chandler.DefaultWrapper(handler.GetCommentList)).Methods("GET")
	r.HandleFunc("/comment", chandler.DefaultWrapper(handler.PostComment)).Methods("POST")

	http.Handle("/", r)
}

func index(w *chandler.RespReqWrapper) {
	t, err := template.ParseFiles("static/index.html")
	if err != nil {
		w.ServerError(err)
		return
	}
	if err := t.Execute(w.W, nil); err != nil {
		w.ServerError(err)
		return
	}
}

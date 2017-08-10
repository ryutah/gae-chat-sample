package web

import (
	"html/template"
	"net/http"
	chandler "server/common/handler"
	"server/web/handler"

	"github.com/gorilla/mux"
)

func Init() {
	r := mux.NewRouter()

	r.HandleFunc("/", chandler.DefaultWrapper(index)).Methods("GET")
	r.HandleFunc("/comment", chandler.DefaultWrapper(handler.GetCommentList)).Methods("GET")
	r.HandleFunc("/comment", chandler.DefaultWrapper(handler.PostComment)).Methods("POST")

	http.Handle("/", r)
}

func index(w *chandler.RespReqWrapper) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		w.ServerError(err)
		return
	}
	if err := t.Execute(w.W, nil); err != nil {
		w.ServerError(err)
		return
	}
}

package web

import (
	"net/http"
	chandler "server/common/handler"
	"server/web/handler"

	"github.com/gorilla/mux"
)

func Init() {
	r := mux.NewRouter()

	r.HandleFunc("/comment", chandler.DefaultWrapper(handler.GetCommentList)).Methods(http.MethodGet)

	http.Handle("/", r)
}

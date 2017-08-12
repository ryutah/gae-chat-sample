package handler

import (
	"encoding/json"
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

type RespReqWrapper struct {
	W       http.ResponseWriter
	R       *http.Request
	Context context.Context
}

func (w RespReqWrapper) ParseJSON(v interface{}) error {
	return json.NewDecoder(w.R.Body).Decode(v)
}

func (w RespReqWrapper) JSON(code int, v interface{}) {
	for k, v := range w.R.Header {
		log.Infof(w.Context, "%v : %v", k, v)
	}
	w.W.Header().Set("Content-Type", "application/json")
	w.W.WriteHeader(code)
	if err := json.NewEncoder(w.W).Encode(v); err != nil {
		w.ServerError(err)
	}
}

func (w RespReqWrapper) BadRequest(msg string) {
	w.JSON(http.StatusBadRequest, Value{
		"msg": msg,
	})
}

func (w RespReqWrapper) ServerError(err error) {
	log.Errorf(w.Context, "%v", err)
	http.Error(w.W, err.Error(), http.StatusInternalServerError)
}

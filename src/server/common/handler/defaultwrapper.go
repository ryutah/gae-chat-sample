package handler

import (
	"net/http"

	"google.golang.org/appengine"
)

type DefaultWrapperHandler func(*RespReqWrapper)

func DefaultWrapper(h DefaultWrapperHandler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)

		wrapper := &RespReqWrapper{
			W:       w,
			R:       r,
			Context: ctx,
		}

		h(wrapper)
	}
}

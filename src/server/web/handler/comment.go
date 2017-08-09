package handler

import (
	"net/http"
	"server/common/db"
	"server/common/handler"
)

type respGetCommentListInner struct {
	ID       int64  `json:"id"`
	Register string `json:"register"`
	Text     string `json:"text"`
}

type respGetCommentList struct {
	Comments []respGetCommentListInner `json:"comments"`
}

func createRespGetCommentList(comments []db.Comment) respGetCommentList {
	list := make([]respGetCommentListInner, len(comments))
	for i, comment := range comments {
		list[i] = respGetCommentListInner{
			ID:       comment.ID,
			Register: comment.CreatedBy,
			Text:     comment.Text,
		}
	}
	return respGetCommentList{
		Comments: list,
	}
}

func GetCommentList(w *handler.RespReqWrapper) {
	comments, err := db.GetAllComment(w.Context)
	if err != nil {
		w.ServerError(err)
		return
	}

	resp := createRespGetCommentList(comments)
	w.JSON(http.StatusOK, resp)
}

type reqPostComment struct {
	Register string `json:"register"`
	Text     string `json:"text"`
}

func PostComment(w *handler.RespReqWrapper) {
	payload := new(reqPostComment)
	if err := w.ParseJSON(payload); err != nil {
		w.BadRequest("can not parse request json")
		return
	}

	comment := &db.Comment{
		Text:      payload.Text,
		CreatedBy: payload.Register,
	}
	if err := db.SaveComment(w.Context, comment); err != nil {
		w.ServerError(err)
		return
	}
	w.JSON(http.StatusCreated, handler.Value{
		"id": comment.ID,
	})
}

package db

import (
	"testing"

	"google.golang.org/appengine/aetest"
)

func Test_CommentValidate(t *testing.T) {
	c := Comment{
		Text:      "hogehoge",
		CreatedBy: "fugafuga",
	}
	if err := c.Validate(); err != nil {
		t.Errorf("should not be return error, got %v", err)
	}
}

func Test_CommentValidateEmptyText(t *testing.T) {
	c := Comment{
		Text:      "",
		CreatedBy: "fugafuga",
	}
	if err := c.Validate(); err == nil {
		t.Error("should be return error")
	}
}

func Test_CommentValidateEmptyCreatedBy(t *testing.T) {
	c := Comment{
		Text:      "hogehgoe",
		CreatedBy: "",
	}
	if err := c.Validate(); err == nil {
		t.Error("should be return error")
	}
}

func Test_SaveComment(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	comment := &Comment{
		Text:      "hogehoge",
		CreatedBy: "fugafuga",
	}
	if err := SaveComment(ctx, comment); err != nil {
		t.Fatal(err)
	}

	if comment.ID == 0 {
		t.Error("id should be set")
	}
	if comment.CreatedAt.IsZero() {
		t.Error("CreatedAt should be set")
	}
	if comment.UpdatedAt.IsZero() {
		t.Error("UpdatedAt should be set")
	}
}

func Test_SaveCommentInvalidComment(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	comment := &Comment{
		Text:      "",
		CreatedBy: "",
	}
	if err := SaveComment(ctx, comment); err == nil {
		t.Error("should be return error")
	}
}

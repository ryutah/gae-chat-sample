package db

import (
	ustr "common/util/strings"
	"time"

	"github.com/pkg/errors"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

const KindComment = "Comment"

type Comment struct {
	ID        int64 `datastore:"-"`
	Text      string
	CreatedBy string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (c Comment) Validate() error {
	if ustr.IsBlank(c.Text) {
		return errors.New("Text should not be empty")
	}
	if ustr.IsBlank(c.CreatedBy) {
		return errors.New("CreatedBy should not be empty")
	}
	return nil
}

func SaveComment(ctx context.Context, comment *Comment) error {
	if err := comment.Validate(); err != nil {
		return errors.WithMessage(err, "SaveComment")
	}

	key := datastore.NewIncompleteKey(ctx, KindComment, nil)

	now := time.Now()
	comment.CreatedAt, comment.UpdatedAt = now, now
	newKey, err := datastore.Put(ctx, key, comment)
	if err != nil {
		return errors.WithMessage(err, "SaveComment")
	}

	comment.ID = newKey.IntID()
	return nil
}

func GetAllComment(ctx context.Context) ([]Comment, error) {
	var comments []Comment

	keys, err := datastore.NewQuery(KindComment).Order("-CreatedAt").Limit(1000).GetAll(ctx, &comments)
	if err != nil {
		return nil, errors.WithMessage(err, "GetAllComment")
	}

	for i, key := range keys {
		comments[i].ID = key.IntID()
	}

	return comments, nil
}

package handler

import (
	"common/db"
	"common/handler"
	"common/util/tests"
	"fmt"
	"testing"

	"golang.org/x/net/context"
)

func Test_GetCommentList(t *testing.T) {
	env, err := tests.CreateAppengineTestEnv(&tests.TestEnvOption{
		StronglyConsistentDatastore: true,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer env.Close()

	if _, err := createTestComments(env.Context, 100); err != nil {
		t.Fatal(err)
	}

	GetCommentList(&handler.RespReqWrapper{
		Context: env.Context,
		R:       env.Request,
		W:       env.Response,
	})

	if env.Response.Code != 200 {
		t.Error("should be response 200")
	}
}

func Test_PostComment(t *testing.T) {
	body, err := tests.CreateJSONRequestBody(map[string]string{
		"register": "SampleUser",
		"text":     "SampleText",
	})
	if err != nil {
		t.Fatal(err)
	}

	env, err := tests.CreateAppengineTestEnv(&tests.TestEnvOption{
		RequestBody: body,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer env.Close()

	PostComment(&handler.RespReqWrapper{
		Context: env.Context,
		R:       env.Request,
		W:       env.Response,
	})

	if env.Response.Code != 201 {
		t.Error("should be response 201")
	}
	tests.AssertResponseJSONHasKeys(t, env.Response, "id")
}

func createTestComments(ctx context.Context, count int) ([]db.Comment, error) {
	comments := make([]db.Comment, count)
	for i := 1; i <= count; i++ {
		comment := db.Comment{
			CreatedBy: fmt.Sprintf("User%v", i),
			Text:      fmt.Sprintf("Text%v", i),
		}
		if err := db.SaveComment(ctx, &comment); err != nil {
			return nil, err
		}
		comments[i-1] = comment
	}
	return comments, nil
}

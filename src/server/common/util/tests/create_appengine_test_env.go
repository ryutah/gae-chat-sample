package tests

import (
	"io"
	"net/http"
	"net/http/httptest"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/aetest"
)

type TestEnvOption struct {
	Method                      string
	URL                         string
	RequestBody                 io.Reader
	TestOption                  *aetest.Options
	AppID                       string
	StronglyConsistentDatastore bool
}

type TestEnv struct {
	Instance aetest.Instance
	Context  context.Context
	Request  *http.Request
	Response *httptest.ResponseRecorder
}

func (t TestEnv) Close() error {
	if t.Instance != nil {
		return t.Instance.Close()
	}
	return nil
}

func CreateAppengineTestEnv(option *TestEnvOption) (*TestEnv, error) {
	parsedOption := parseOption(option)

	inst, err := aetest.NewInstance(&aetest.Options{
		AppID: parsedOption.AppID,
		StronglyConsistentDatastore: parsedOption.StronglyConsistentDatastore,
	})
	if err != nil {
		return nil, err
	}

	r, err := inst.NewRequest(parsedOption.Method, parsedOption.URL, parsedOption.RequestBody)
	if err != nil {
		inst.Close()
		return nil, err
	}

	ctx := appengine.NewContext(r)

	env := &TestEnv{
		Instance: inst,
		Context:  ctx,
		Response: httptest.NewRecorder(),
		Request:  r,
	}
	return env, nil
}

func parseOption(option *TestEnvOption) *TestEnvOption {
	retVal := &TestEnvOption{
		Method: "GET",
		URL:    "/",
	}

	if option == nil {
		return retVal
	}
	if option.Method != "" {
		retVal.Method = option.Method
	}
	if option.URL != "" {
		retVal.URL = option.URL
	}
	retVal.RequestBody = option.RequestBody
	retVal.TestOption = option.TestOption

	return retVal
}

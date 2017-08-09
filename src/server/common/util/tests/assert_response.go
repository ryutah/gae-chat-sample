package tests

import (
	"encoding/json"
	"net/http/httptest"
	"reflect"
	"testing"
)

func AssertResponseJSONHasKeys(t *testing.T, r *httptest.ResponseRecorder, keys ...string) {
	v := make(map[string]interface{})
	if err := json.Unmarshal(r.Body.Bytes(), &v); err != nil {
		t.Fatal(err)
	}

	var lackKeys []string
	for _, key := range keys {
		if _, ok := v[key]; !ok {
			lackKeys = append(lackKeys, key)
		}
	}
	if len(lackKeys) != 0 {
		t.Errorf("response json does not have %v", lackKeys)
	}
}

func AssertResponseJSONEql(t *testing.T, r *httptest.ResponseRecorder, v interface{}) {
	vType := reflect.TypeOf(v)

	var val reflect.Value
	if vType.Kind() == reflect.Map {
		typ := reflect.MapOf(vType.Key(), vType.Elem())
		val = reflect.New(typ)
	} else if vType.Kind() == reflect.Ptr {
		typ := vType.Elem()
		val = reflect.New(typ)
	} else {
		val = reflect.New(vType)
	}

	if err := json.Unmarshal(r.Body.Bytes(), val.Interface()); err != nil {
		t.Fatal(err)
	}

	var got interface{}
	if vType.Kind() == reflect.Ptr || vType.Kind() == reflect.Map {
		got = val.Interface()
	} else {
		got = val.Elem().Interface()
	}
	if !reflect.DeepEqual(v, got) {
		t.Errorf("response body want %v, got %v", v, got)
	}
}

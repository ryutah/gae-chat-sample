package strings

import (
	"testing"
)

func Test_Include(t *testing.T) {
	tgt := []string{"a", "b", "c"}
	if !Include("a", tgt) {
		t.Error("should be return true")
	}
}

func Test_Include_NotIncludeString(t *testing.T) {
	tgt := []string{"a", "b", "c"}
	if Include("d", tgt) {
		t.Error("should be return false")
	}
}

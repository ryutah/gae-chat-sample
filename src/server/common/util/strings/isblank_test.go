package strings

import (
	"testing"
)

func Test_IsBlank(t *testing.T) {
	s := " 　\n\t"
	if !IsBlank(s) {
		t.Errorf("should be return true")
	}
}

func Test_IsBlankNotBlankText(t *testing.T) {
	s := " 　\nab\t"
	if IsBlank(s) {
		t.Errorf("should be return false")
	}
}

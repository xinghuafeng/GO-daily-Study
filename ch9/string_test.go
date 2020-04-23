package ch9

import (
	"strings"
	"testing"
)

func TestStringFn(t *testing.T) {
	m := "A,B,C"
	parts := strings.Split(m, ",")
	for _, v := range parts {
		t.Log(v)
	}
	t.Log(strings.Join(parts, "-"))
}

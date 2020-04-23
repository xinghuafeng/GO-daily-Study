package ch9

import (
	"strconv"
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
func TestConvert(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("str" + s)

}

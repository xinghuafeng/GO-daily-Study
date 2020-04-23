package ch8

import "testing"

// TODO go利用map 可以实现工厂模式
func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](2), m[2](2), m[3](2))

}

//go 中实现set 已实现
func TestMapForSet(t *testing.T) {
	myset := map[int]bool{}
	myset[1] = true
	n := 3
	if myset[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is  not existing", n)

	}
	myset[3] = true
	t.Log(len(myset))
	delete(myset, 1)
	t.Log(myset[1])

}

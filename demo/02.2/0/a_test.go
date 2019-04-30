package a

import (
	"testing"
)

func Test_A(t *testing.T)  {
	storage := &RedisStorage{}
	a := NewA(storage)
	a.Save("aa", "1")
	aa, _ := a.Get("aa")
	if aa != "1" {
		t.Errorf("aa expected 1, but get %v", aa)
	}
}

 
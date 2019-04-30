package a

import (
	"testing"
) 

type MockStorage struct {
	m map[string]string
}

func NewMockStorage() Storage  {
	return &MockStorage{
		m : make(map[string]string, 0),
	}
} 

func (s *MockStorage) Set(key, val string) error {
	s.m[key] = val
	return nil
}

func (s *MockStorage) Get(key string) (val string,  err error) {
	val = s.m[key]
	return val, nil
}

 
func Test_A(t *testing.T)  {
	a := NewA(NewMockStorage())
	a.Save("aa", "1")
	aa, _ := a.Get("aa")
	if aa != "1" {
		t.Errorf("aa expected 1, but get %v", aa)
	}
}



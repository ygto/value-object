package valobj_test

import (
	"testing"
	"github.com/ygto/value-object"
)

func TestString(t *testing.T) {
	username := "John Doe"
	data := valobj.Val(username)
	if data.String() != username {
		t.Fail()
	}
}
func TestInt(t *testing.T) {
	age := 50
	data := valobj.Val(age)
	if data.Int() != age {
		t.Fail()
	}
}
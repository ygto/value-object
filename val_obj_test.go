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
func TestInt8(t *testing.T) {
	var age int8
	age = 50
	data := valobj.Val(age)
	if data.Int8() != age {
		t.Fail()
	}
}
func TestInt16(t *testing.T) {
	var age int16
	age = 50
	data := valobj.Val(age)
	if data.Int16() != age {
		t.Fail()
	}
}
func TestInt32(t *testing.T) {
	var age int32
	age = 50
	data := valobj.Val(age)
	if data.Int32() != age {
		t.Fail()
	}
}
func TestInt64(t *testing.T) {
	var age int64
	age = 50
	data := valobj.Val(age)
	if data.Int64() != age {
		t.Fail()
	}
}

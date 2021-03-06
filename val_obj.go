package valobj

type Valobj struct {
	data interface{}
}

func Val(data interface{}) *Valobj {
	v := Valobj{}
	v.data = data
	return &v
}
func (v *Valobj) Set(data interface{}) {
	v.data = data
}
func (v *Valobj) Get() interface{} {
	return v.data
}
func (v *Valobj) String() string {
	return v.Get().(string)
}
func (v *Valobj) Int() int {
	return v.Get().(int)
}

func (v *Valobj) Int8() int8 {
	return v.Get().(int8)
}

func (v *Valobj) Int16() int16 {
	return v.Get().(int16)
}
func (v *Valobj) Int32() int32 {
	return v.Get().(int32)
}
func (v *Valobj) Int64() int64 {
	return v.Get().(int64)
}

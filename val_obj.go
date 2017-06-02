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

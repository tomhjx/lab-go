package object

import (
	"testing"
)

type IObj interface {
	GetName() string
}

type Obj1 struct {
	name string
}

func (o *Obj1) GetName() string {
	return o.name
}

type Obj2 struct {
	Obj1
}

func TestGetObjectInfo(t *testing.T) {
	obj1 := &Obj1{name: "object1"}
	r := GetObjectInfo(obj1)
	t.Logf("obj1.GetName:%s", obj1.GetName())
	t.Logf("%v", r)

	obj2 := &Obj2{Obj1{name: "object2"}}
	r = GetObjectInfo(obj2)
	t.Logf("obj2.GetName:%s", obj2.GetName())
	t.Logf("%v", r)

}

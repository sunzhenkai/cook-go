package main

import (
	"cooking-go/pkg/gen/pb/cook/pb"
	"google.golang.org/protobuf/proto"
)

//type Int64 int64

func SetFoo(foo *pb.Foo) {
	var id = int32(1)
	foo.Id = &id
}

type A struct {
	V int
}

func main() {
	var x = int64(5)
	println(x)

	var a = A{}
	println(a.V)

	v := pb.Foo{}
	SetFoo(&v)
	println(*v.Id)

	var vp *int
	vp = new(int)
	println(*vp)

	var i int = int(1)
	vp = &i
	vp = new(int)
	println(i)

	var ap *A
	ap = &A{}
	println(ap.V)

	foo := pb.Foo{}
	foo.Id = proto.Int32(23)
	println(*foo.Id)
}

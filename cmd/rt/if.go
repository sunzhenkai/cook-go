package main

import "reflect"

type Person struct {
	Age int
	Name string
}

func IF(condition bool, tv, fv interface{}) interface{} {
	if condition {
		return tv
	} else {
		return fv
	}
}

func main() {
	a := 1
	b := 2
	c := IF(a > b, a, b)
	println(reflect.TypeOf(c).Name())
	println(c.(int))

	pa := Person{Age: 18, Name: "wii"}
	pb := Person{Age: 16, Name: "bvz"}
	pc := IF(pa.Age > pb.Age, pa, pb)
	println(&pa, &pb, &pc)

	pd := IF(pa.Age > pb.Age, &pa, &pb).(*Person)
	println(&pa, &pb, pd)
}

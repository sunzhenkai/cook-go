package main

import "fmt"

func f(msg *string) {
	println("run f: " + *msg)
}

type S struct {
	B bool
}

func Sample() {
	//ss := "sa"
	s := S{}
	println(s.B)
	v := "va"
	defer f(&v)
	println("done ", v)
	v = "vb"
	defer f(&v)
	defer func() {
		println(s.B)
	}()
	defer println(s.B)
	//ss = "sb"
	s.B = true
	println("done ", v)

	v = "vc"
	defer func(v *string) {
		f(v)
	}(&v)
}

type SA struct {
	Name string
}

func SampleARef(sa *SA) {
	*sa = SA{"OK"}
}

func SampleAP(sa **SA) {
	*sa = &SA{"Point"}
}

var sap *SA

func SampleA() {
	SampleAP(&sap)
	fmt.Println(sap.Name)
	fmt.Println("HHH")
	sa := SA{"NO"}
	SampleARef(&sa)
	fmt.Println(sa.Name)
}

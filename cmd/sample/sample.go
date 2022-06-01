package main

func f(msg *string) {
	println("run f: " + *msg)
}

type S struct {
	B bool
}

func main() {
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

package main

type Dt struct {
	content string
}

func Ifc(ifc interface{})  {
	println("Ifc, &ifc:", &ifc)
}

func IfcPtr(ifc interface{})  {
	println("IfcPtr, ifc:", ifc.(*Dt))
}

func main() {
	dt := Dt{"wii"}
	println("main, &dt:", &dt)
	Ifc(dt)
	IfcPtr(&dt)
}
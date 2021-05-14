package main

type Data struct {
	name string
}

func (p *Data) Ptr() {
	println("ptr, p: ", p)
	println("ptr: ", p.name)
}

func (p Data) Cp() {
	println("cp, &p: ", &p)
	println("cp: ", p.name)
}

func Ptr(p *Data) {
	println("ptr, p: ", p)
	println("ptr: ", p.name)
}

func Cp(p Data) {
	println("cp, &p: ", &p)
	println("cp: ", p.name)
}

func main() {
	da := Data{name: "wii"}
	println("main, &da: ", &da)
	da.Ptr()
	(&da).Ptr()
	da.Cp()
	(&da).Cp()

	println("-----")

	Ptr(&da)
	// Ptr(da) // ERROR: Cannot use 'da' (type Data) as the type *Data
	Cp(da)
	// Cp(&da)    // ERROR: Cannot use '&da' (type *Data) as the type Data
}

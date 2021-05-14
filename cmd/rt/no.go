package main

func main() {
	i := 10
	f := 1.2
	// println(i * f)     ERROR: Invalid operation: i * f (mismatched types int and float64)
	println(float64(i) * f)
	println(i * int(f))
}

package main

import "fmt"

type Bird interface {
	Name() string
	Fly() bool
	Chirp() string
}

type Duck struct {
	nm     string
	CanFly bool
	Sound  string
}

func NewDuck() *Duck {
	return &Duck{CanFly: false, Sound: "quack", nm: "Duck"}
}

func (duck *Duck) Name() string {
	return duck.nm
}

func (duck *Duck) Fly() bool {
	return duck.CanFly
}

func (duck *Duck) Chirp() string {
	return duck.Sound
}

func IFC(condition bool, tv interface{}, fv interface{}) interface{} {
	if condition {
		return tv
	} else {
		return fv
	}
}

func main() {
	var bird Bird = NewDuck()
	fmt.Printf("%v\n", IFC(bird.Fly(), bird.Name()+" is flying!", bird.Name()+" cannot fly!"))
	println(bird.Name()+" is chirping:", bird.Chirp(), bird.Chirp(), "...")
}

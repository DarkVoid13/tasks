// task1.go
package main

import (
	"fmt"
)

type animals struct {
	age    int
	weight int
	color  string
}
type Duck struct {
	animals
	name   string
	height int
	price  float32
}

func (a animals) Inf1() (string, int, int) {
	return a.color, a.age, a.weight
}
func Inf2(a animals) (int, int) {
	return a.age, a.weight
}
func main() {
	a1 := animals{12, 14, "blue"}
	fmt.Println(a1.Inf1())
	fmt.Println(Inf2(a1))
}

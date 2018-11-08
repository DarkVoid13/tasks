// task1.go
package main

import (
	"fmt"
)

type Animals struct {
	age    int
	weight int
	color  string
}
type Duck struct {
	Animals
	name   string
	height int
	price  float32
}

func (a Animals) Inf1() (string, int, int) {
	return a.color, a.age, a.weight
}
func Inf2(a Animals) (int, int) {
	return a.age, a.weight
}
func main() {
	a1 := Animals{12, 14, "blue"}
	fmt.Println(a1.Inf1())
	fmt.Println(Inf2(a1))
}

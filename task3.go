// task2.go
package main

import (
	"fmt"
)

type animals struct {
	age    int
	weight int
	color  string
}

type inform interface {
	Inf() string
}

func (a animals) Inf() string {
	return a.color
}

func (a animals) String() string {
	return fmt.Sprintf("%d %d %s", a.age, a.weight, a.color)
}
func informFunc(i inform) string {
	return i.Inf()
}
func main() {
	array := [3]inform{
		animals{12, 11, "blue"},
		animals{11, 12, "blue"},
		animals{12, 13, "blue"}}

	for i := range array {
		fmt.Println(i, " | ", array[i].Inf())
	}

	for i := range array {
		fmt.Println(i, " | ", informFunc(array[i]))
	}
}

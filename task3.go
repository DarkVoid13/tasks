// task2.go
package main

import (
	"fmt"
)

type Animals struct {
	age    int
	weight int
	color  string
}

type inform interface {
	Inf() string
}

func (a Animals) Inf() string {
	return a.color
}

func (a Animals) String() string {
	return fmt.Sprintf("%d %d %s", a.age, a.weight, a.color)
}
func informFunc(i inform) string {
	return i.Inf()
}
func main() {
	array := [3]inform{
		Animals{12, 11, "blue"},
		Animals{11, 12, "blue"},
		Animals{12, 13, "blue"}}

	for i := range array {
		fmt.Println(i, " | ", array[i].Inf())
	}

	for i := range array {
		fmt.Println(i, " | ", informFunc(array[i]))
	}
}

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

func main() {
	array := []Animals{
		{12, 11, "blue"},
		{11, 12, "blue"},
		{12, 13, "blue"}}

	for i := range array {
		fmt.Println(i, "|", array[i])
	}
}

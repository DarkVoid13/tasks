// task4
package main

import (
	"fmt"
)

func Art() func() {
	return func() {
		defer fmt.Println("World")
		fmt.Println("Hello")
	}
}

func main() {
	a := Art()
	a()

}

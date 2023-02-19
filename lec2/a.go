package main

import (
	"fmt"
	"github.com/moonidelight/go_course/module"
)

func main() {
	fmt.Println("Hello")
	radius := module.Circle{5}
	fmt.Print("area of circle ")
	fmt.Println(module.Area(radius))
	fmt.Println(module.Length(radius))
	rec := module.Rectangle{16, 20}
	fmt.Println(module.Area(rec))
	fmt.Println(module.Length(rec))

}

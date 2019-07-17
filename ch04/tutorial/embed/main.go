package main

import "fmt"

type Point struct {
	X, Y int
}

type Coord struct {
	X, Y int
}

type Circle struct {
	Point
	//Coord  // add this field will cause ambiguity in w.X and w.Y, which won't pass compilation
	Radius float64
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	w := Wheel{Circle{Point{1, 2}, 3.14}, 20}
	fmt.Printf("%v\n", w)
	fmt.Println(w.X, w.Y)
}

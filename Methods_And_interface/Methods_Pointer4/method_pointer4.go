package main

import (
	f "fmt"
)

// ประกาศฟังก์ชัน struct ชื่อตัวแปร Vertex กำหนดฟิลด์สองตัว x และ y ชนิดข้อมูล float64
type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	f.Println(v, p)
}

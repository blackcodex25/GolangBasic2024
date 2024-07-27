package main

import (
	f "fmt"
)

/*
	การเปรียบเทียบการใช้งานฟังก์ชันที่รับ pointer เป็นพารามิเตอร์และเมธอดที่ใช้ pointer

receiver ในภาษา Go
*/
// ประกาศฟังก์ชัน struct ชื่อตัวแปร Vertex ที่มีสองฟิลด์คือ x และ y ชนิดข้อมูล float64
type Vertex struct {
	X, Y float64
}

// เมธอดชื่อตัวแปร Scale ที่ทำงานกับ pointer receiver ชื่อตัวแปร *Vertex
func (v *Vertex) Scale(f float64) {
	// เมธอดนี้ปรับขนาดของ x และ y โดยคูณด้วย f
	v.X = v.X * f
	v.Y = v.Y * f
}

// ฟังก์ชัน ScaleFunc ที่รับ pointer ไปยัง Vertex เป็นพารามิเตอร์
func ScaleFunc(v *Vertex, f float64) {
	// เมธอดนี้ปรับขนาดของ x และ y โดยคูณด้วย f
	v.X = v.X * f
	v.Y = v.Y * f
}
func main() {
	v := Vertex{3, 4}
	v.Scale(2)        // เรียกใช้เมธอด Scale ที่มี pointer receiver
	ScaleFunc(&v, 10) // เรียกใช้ฟังก์ชัน ScaleFunc โดยส่ง pointer

	p := &Vertex{4, 3}
	v.Scale(3)      // เรียกใช้เมธอด Scale โดย pointer
	ScaleFunc(p, 8) // เรียกใช้ฟังก์ชัน ScaleFunc โดย pointer

	f.Println(v, p) // พิมพ์ผลลัพธ์ของ v และ p ออกจอภาพ
}

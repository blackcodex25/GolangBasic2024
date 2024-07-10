package main

import (
	f "fmt"
)

// ประกาศตัวแปรนอกฟังก์ชัน
// การใช้ var หลายบรรทัดในรูปแบบ
// ของ block เพื่อประกาศตัวแปร d, e, และ f
var (
	a, b, c int = 1, 2, 3
	d       int = 4
	e       int = 5
	x       int = 6
)

// ประกาศตัวแปรภายในฟังก์ชัน
// ประกาศตัวแปรสามตัวพร้อมกันภายใน
// ฟังก์ชัน main และกำหนดค่าเริ่มต้นโดยใช้ :=
func main() {
	g, h, i := 7, 8, "Hello"
	f.Println(a, b, c)
	f.Println(d, e, x)
	f.Println(g, h, i)
}

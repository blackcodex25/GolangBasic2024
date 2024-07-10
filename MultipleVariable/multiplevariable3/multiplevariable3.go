package main

import (
	f "fmt"
)

// การประกาศตัวแปรหลายตัวในบล็อก
// ฟังก์ชัน var ใช้ได้ทั้งนอกและในฟังก์ชันหลัก
var (
	a int    //ค่าเริ่มต้นคือ 0
	b int    = 1
	c string = "hello"
)

func main() {
	//ปริ้นออกจอ
	f.Println(a)
	f.Println(b)
	f.Println(c)
	f.Println(a, b, c) //การใช้งานปริ้นหลายตัวแปรในบรรทัดเดียว
}

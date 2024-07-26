package main

import (
	f "fmt"
	"math"
)

/*
การประกาศและการใช้ Interface
กำหนด interface ชื่อ Abser
*/
type Abser interface {
	Abs() float64
}

/*
Abser คือ interface ที่กำหนด method ชื่อตัวแปร Abs ซึ่งต้องคืนค่าประเภท float64
การใช้ interface
*/

// ประกาศฟังก์ชัน struct ชื่อตัวแปร Vertex กำหนดค่าแกน x และ แกน y ชนิด float64
type Vertex struct {
	X, Y float64
}

/*
ตัวอย่างการประกาศและการใช้ Method
สร้าง method ชื่อตัวแปร Abs สำหรับ Vertex
- สร้าง Method ชื่อ Abs() ที่มี receiver ชื่อ v ซึ่งเป็นชนิด Vertex
- Abs จะคำนวณและคืนค่าระยะห่างจากจุดศูนย์กลางของ Vertex
*/
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// ฟังก์ชันที่ทำงานกับ interface
func Describe(a Abser) {
	f.Println(a.Abs())
}

func main() {
	v := Vertex{3, 4}  // ประกาศตัวแปร v เรียกใช้งาน struct กำหนดค่าแกน x และแกน y
	f.Println(v.Abs()) // เรียก method Abs โดยตรง

	var a Abser = v    // ใช้ interface รับค่าของ Vertex
	f.Println(a.Abs()) // เรียก method Abs ผ่าน interface
	//Describe(v)        // ใช้ interface ในฟังก์ชัน Describe

	/* var ตัวแปร a = v คือการใช้ interface ชื่อตัวแปร Abser รับ
	instance ของ Vertex
	a.Abs() เรียกใช้ method ผ่าน interface
	*/
}

/* ฟังก์ชัน Describe รับ interface ชื่อตัวแปร Abser
เป็นพารามิเตอร์และเรียก method ชื่อตัวแปร Abs บน interface

Describe(v) เรียกใช้ฟังก์ชันโดยส่ง instance ของ Vertex
ซึ่ง implement method ชื่อตัวแปร Abs
*/
/* สรุป
- Methods คือฟังก์ชันที่มี receiver เป็นชนิดข้อมูลที่กำหนดและใช้เพื่อดำเนินการกับ
ข้อมูลของชนิดนั้น
- interfaces คือชนิดข้อมูลที่กำหนดชุดของ method ที่ชนิดข้อมูลอื่นๆ
ต้อง implement โดยไม่ต้องรู้รายละเอียดการทำงานของชนิดข้อมูลเหล่านั้น

การใช้ methods และ interfaces ใน Go ช่วยให้โปรแกรมมีความยืดหยุ่นสูงและ
สามารถจัดการกับโค้ดที่ซับซ้อนได้ง่ายขึ้น
*/

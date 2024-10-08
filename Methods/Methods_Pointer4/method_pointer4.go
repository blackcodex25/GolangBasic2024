package main

import (
	f "fmt"
)

/* โค้ดนี้เป็นโปรแกรมภาษา Go ที่ใช้สำหรับการปรับขนาด (scaling) ของเวกเตอร์
ในระบบพิกัดคาร์ทีเซียน (Cartesian coordinates)
โดยมีการกำหนดชนิดข้อมูล (type) สำหรับเวกเตอร์ และมีการใช้งานทั้งเมธอด (method)
และฟังก์ชัน (function) ในการปรับขนาดของเวกเตอร์
*/
// ประกาศฟังก์ชัน struct ชื่อตัวแปร Vertex กำหนดฟิลด์สองตัว x และ y ชนิดข้อมูล float64
type Vertex struct {
	X, Y float64
}

// เมธอดนี้ปรับขนาดของฟิลด์ x และ y ของ Vertex โดยคูณด้วยค่าพารามิเตอร์ f
// ประกาศเมธอดชื่อ Scale() ชื่อตัวแปร f ชนิด float64 สำหรับ Vertex ที่ใช้ pointer receiver ชื่อตัวแปร v ชี้ค่าสตรัค *Vertex
func (v *Vertex) Scale(f float64) {
	// การเข้าถึงฟิลด์ของสตรัค Vertex โดยเข้าถึงผ่าน pointer receiver ชื่อตัวแปร v
	// ประกาศตัวแปร v.X กำหนดเข้าถึงฟิลด์ของสตรัค Vertex ชื่อฟิลด์ x และกำหนด v.X คูณ พารามิเตอร์ f
	// ประกาศตัวแปร v.Y กำหนดเข้าถึงฟิลด์ของสตรัค Vertex ชื่อฟิลด์ y และกำหนด x.Y คูณ พารามิเตอร์ f
	v.X = v.X * f
	v.Y = v.Y * f
}

// ประกาศฟังก์ชัน ScaleFunc ที่รับ Pointer ไปยัง Vertex (*Vertex) และค่าพารามิเตอร์ f ชนิด float64
// ฟังก์ชันนี้ปรับขนาดของฟิลด์ x และ y ของ Vertex โดยคูณด้วยค่าพารามิเตอร์ f
// ประกาศฟังก์ชัน ScaleFunc กำหนดตัวแปร v ชนิดพอยเตอร์ *Vertex และกำหนดพารามิเตอร์ f ชนิด float64
func ScaleFunc(v *Vertex, f float64) {
	// การเข้าถึงฟิลด์ของสตรัค Vertex โดยผ่าน pointer receiver ชื่อตัวแปร v
	// ประกาศตัวแปร v.X กำหนดเข้าถึงฟิลด์ของสตรัค Vertex ชื่อฟิลด์ x และกำหนด v.X คูณ พารามิเตอร์ f
	// ประกาศตัวแปร v.Y กำหนดเข้าถึงฟิลด์ของสตรัค Vertex ชื่อฟิลด์ y และกำหนด v.Y คูณ พารามิเตอร์ f
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	// สร้างตัวแปร v ซึ่งเป็นชนิดฟังก์ชันสตรัคชื่อ Vertex และกำหนดค่าเริ่มต้นอากิวเมนต์ของฟิลด์ x เป็น 3 และ y เป็น 4
	v := Vertex{3, 4}
	// เรียกใช้เมธอด Scale ของ v โดยคูณค่าของ x และ y ด้วย 2
	v.Scale(2)
	// เรียกใช้ฟังก์ชัน ScaleFunc() โดยส่ง pointer ของ v และค่าพารามิเตอร์ 10
	ScaleFunc(&v, 10)
	// สร้าง pointer ชื่อตัวแปร p ที่ชี้ไปยัง Vertex ที่มีค่าเริ่มต้น x ค่าพารามิเตอร์เป็น 4 และ y ค่าพารามิเตอร์เป็น 3
	p := &Vertex{4, 3}
	// เรียกใช้เมธอด Scale ของ pointer ตัวแปร p โดยคูณค่าของ x และ y ด้วยค่าพารามิเตอร์ 3
	p.Scale(3)
	// เรียกใช้ฟังก์ชัน ScaleFunc โดยส่ง pointer ชื่อตัวแปร p กำหนดค่าพารามิเตอร์ 8
	ScaleFunc(p, 8)
	// พิมพ์ค่าของ v และ p ออกทางหน้าจอ
	f.Println(v, p)
}

/* สรุป
โปรแกรมนี้แสดงการใช้งานทั้งเมธอดและฟังก์ชันในการปรับ
ขนาดของเวกเตอร์ในระบบพิกัดคาร์ทีเซียน (Cartesian coordinates) โดยมีการแสดง
ให้เห็นว่าการใช้ pointer receiver ในเมธอดช่วยให้สามารถเรียกใช้งานได้ทั้งจากตัวแปร
และ pointer
ในขณะที่การใช้ฟังก์ชันที่รับ pointer เป็นพารามิเตอร์จะต้องส่ง pointer
เข้าไปเท่านั้น
*/

package main

import (
	f "fmt"
)

// วิธีการใช้ Pointer เพื่อเข้าถึงและแก้ไขฟิลด์ของ struct โดยตรงในภาษา Go
// ประกาศ struct ชื่อ Vertex กำหนดฟิลด์ x และ y ชนิด int
type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2} // สร้างตัวแปร v ของประเภท Vertex และตั้งค่าเริ่มต้นให้กับฟิลด์ x และ y
	p := &v           // ประกาศ pointer ชื่อตัวแปร p ที่ชี้ไปยังตัวแปร v
	p.X = 1e9         // เปลี่ยนค่าของฟิลด์ x ผ่าน pointer ชื่อตัวแปร p
	f.Println(v)      // แสดงค่าของตัวแปร v ออกจอภาพ
}

// Note: 1e9 เป็นวิธีการเขียนเลขฐานสิบแแบบ exponential notation
// ในภาษา Go ซึ่งหมายถึง 1 x 10 ยกกำลัง 9 หรือ 1000000000 (หนึ่งพันล้าน)

/* สรุป
การประกาศ: struct ใช้ type ประกาศ struct ใหม่
การเข้าถึงฟิลด์ผ่าน: pointer สามารถเข้าถึงและแก้ไขฟิลด์ของ struct
ผ่าน pointer โดยไม่ต้องใช้การ dereferencing ชัดเจน
การแสดงผล: ใช้ฟังก์ชัน Println() เพื่อแสดงค่าของ struct
*/

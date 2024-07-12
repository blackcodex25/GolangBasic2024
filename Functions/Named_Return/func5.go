package main

import (
	f "fmt"
)

/*
โค้ดตัวอย่างนี้ เราตั้งชื่อค่าที่คืนว่า ressult (เป็นชนิด int) และคืน
ค่าด้วยการใช้ "naked return"(หมายความว่าเราใช้คำสั่ง return
โดยไม่ระบุชื่อตัวแปร)
*/
//ประกาศฟังก์ชันชื่อ myfunction ซึ่งมีพารามิเตอร์ 2 ตัว x และ y รับข้อมูลเป็น int
//ประกาศพารามิเตอร์ชื่อ result ชนิดข้อมูล int
func myFunction(x int, y int) (result int) {
	result = x + y //เรียกใช้พารามิเตอร์ result รับค่าผลรวม x + y
	return
}
func main() {
	//เรียกใช้งานฟังก์ชัน ชื่อ myfunction() กำหนดค่า 1 , 2 แทน x และ y
	f.Println(myFunction(1, 2)) //พิมพ์ผลลัพธ์ออกจอ
}

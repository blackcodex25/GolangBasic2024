package main

import (
	f "fmt"
)

// การใช้คำสั่ง return ระบุชื่อตัวแปรที่จะคืนค่า
// ประกาศฟังก์ชันชื่อ myFunction() กำหนดพารามิเตอร์ชื่อ x และ y
// สร้างพารามิเตอร์ที่จะคืนค่าชื่อ result
func myFunction(x int, y int) (result int) {
	result = x + y //เรียกใช้พารามิเตอร์ชื่อ result รับผลรวมค่า x + y
	return result  //ประกาศฟังก์ชัน return ระบุชื่อตัวแปรที่จะคืนค่า ชื่อ result
}
func main() {
	//เรียกใช้งานฟังก์ชันชื่อ myFunctions() กำหนดค่า 1, 2
	f.Println(myFunction(1, 2)) //พิมพ์ออกจอภาพ
} // ผลลัพธ์ 3

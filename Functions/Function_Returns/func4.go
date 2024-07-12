package main

import (
	f "fmt"
)

//การคืนค่าจากฟังก์ชันในภาษา Go
/*การคืนค่า (Return Values)
หากต้องการให้ฟังก์ชันคืนค่า คุณต้องกำหนดชนิดของข้อมูลที่ฟังก์ชันจะคืนค่า
(เช่น int, string,ฯลฯ)และใช้คำสั่ง return ภายในฟังก์ชัน
*/
/*
	โค้ดตัวอย่างการคืนค่าจากฟังก์ชัน myfunction() รับอินพุตเป็นจำนวนเต็ม
	สองจำนวน (x และ y) และคืนค่าผลรวมของทั้งสองเป็นจำนวนเต็ม (int)

			*Syntax
	*func FunctionName(param1 type, param2 type) type {
	*โค้ดที่จะทำงาน
	*return output
}
*/
//ประกาศฟังก์ชันชื่อ myfunction ที่มีพารามิเตอร์ 2 ตัว x และ y รับข้อมูลเป็น int
//ประกาศ type ชนิดข้อมูล int นอกวงเล็บ
func myFunction(x int, y int) int {
	return x + y //คืนค่าผลรวม x + y ชนิดข้อมูล int
}
func main() {
	//เรียกใช้งานฟังก์ชัน ชื่อ myfunction() กำหนดค่า 1 , 2 แทน x และ y
	f.Println(myFunction(1, 2)) //พิมผลลัพธ์ออกจอ
}

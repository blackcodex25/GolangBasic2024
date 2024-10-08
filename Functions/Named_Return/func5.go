package main

import (
	f "fmt"
)

/*
โค้ดตัวอย่างนี้ เราตั้งชื่อค่าที่คืนว่า result (เป็นชนิด int) และคืน
ค่าด้วยการใช้ "naked return"(หมายความว่าเราใช้คำสั่ง return
โดยไม่ระบุชื่อตัวแปร)
*/
/*ประกาศฟังก์ชันชื่อ myfunction ซึ่งมีพารามิเตอร์ 2 ตัว x และ y รับข้อมูลเป็น int
ประกาศพารามิเตอร์ชื่อ result ชนิดข้อมูล int */

// การใช้คำสั่ง return แบบไม่ระบุชื่อตัวแปร
func myFunction(x int, y int) (result int) {
	result = x + y //เรียกใช้พารามิเตอร์ result รับค่าผลรวม x + y
	return         //คืนค่าให้ตัวแปรอะไร ?
}
func main() {
	//เรียกใช้งานฟังก์ชัน ชื่อ myfunction() กำหนดค่า 1 , 2 แทน x และ y
	f.Println(myFunction(1, 2)) //พิมพ์ผลลัพธ์ออกจอ
}

/*ฟังก์ชันสามารถมีค่าที่คืนได้มากกว่าหนึ่งค่า
โดยการใช้วงเว็บล้อมรอบชนิดข้อมูลที่จะคืน
ค่าที่คืนจากฟังก์ชันสามารถใช้งานได้โดยตรงในโค้ดที่เรียกใช้ฟังก์ชัน*/
/* การใช้งานฟังก์ชันที่คืนค่าเป็นสิ่งที่มีประโยชน์มากในการเขียนโปรแกรม ทำให้โค้ด
มีความกระชับและสามารถแยกส่วนการทำงานได้ดีขึ้น
*/

package main

import (
	f "fmt"
)

//พารามิเตอร์และอาร์กิวเมนต์
/*สามารถส่งข้อมูลให้ฟังก์ชันได้ในรูปแบบของพารามิเตอร์ โดย
พารามิเตอร์จะทำหน้าที่เป็นตัวแปรภายในฟังก์ชัน

กำหนดพารามิเตอร์
พารามิเตอร์และชนิดของพารามิเตอร์จะถูกกำหนดหลังชื่อฟังก์ชัน
ภายในวงเล็บคุณสามารถเพิ่มพารามิเตอร์ได้มากเท่าที่ต้องการ
โดยแยกแต่ละพารามิเตอร์ด้วยเครื่องหมายจุลภาค(,)
*/
/* 			Syntax
*func functionName(param1 type, param2 type, param3 type){
 โค้ดที่จะถูกดำเนินการ
}
*/
/* ฟังก์ชันที่มีพารามิเตอร์หนึ่งตัว (dev) ที่มีชนิด
เป็น string เมื่อเรียกใช้ฟังก์ชัน programing() เราจะส่งชื่อไปด้วย
(เช่น Golang) และชื่อนั้นจะถูกใช้ภายในฟังก์ชันเพื่อแสดงผลชื่อเต็ม
*/
func programing(dev string) {
	f.Println("Hi", dev, "Backend")
}
func main() {
	programing("Golang")
	programing("C++")
	programing("Python")
}

/*อธิบายโค้ด
dev เป็นพารามิเตอร์ ส่วน Golang, C++, และ Python เป็นอาร์กิวเมนต์
*/
/* Note
เมื่อพารามิเตอร์ถูกส่งไปยังฟังก์ชัน จะเรียกว่าอาร์กิวเมนต์
(argument) ดังนั้นจากตัวอย่างข้างต้น dev คือพารามิเตอร์
ส่วน Golang, C++, และ Python คืออาร์กิวเมนต์
*/

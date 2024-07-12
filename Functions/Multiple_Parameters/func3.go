package main

import (
	f "fmt"
)

//การใช้พารามิเตอร์หลายตัว
/*การใช้พารามิเตอร์หลายตัวในฟังก์ชัน
ภายในฟังก์ชัน เราสามารถเพิ่มพารามิเตอร์ได้ตามต้องการ
*/
/* 			Syntax
*func functionName(param1 type, param2 type, param3 type){
 โค้ดที่จะถูกดำเนินการ
}
*/
// ประกาศฟังก์ชันชื่อ familyName() มีพารามิเตอร์ 2 ตัว ชื่อ fname ชนิดข้อมูล string
// และชื่อ age ชนิดข้อมูล int
func familyName(fname string, age int) {
	f.Println("Hi", age, "year", fname, "Refsnes") //พิมพ์ข้อความออกจอภาพ
}
func main() {
	//เรียกใช้งานฟังก์ชัน familyName() ด้วยชื่อ Liam และอายุ 3
	//เรียกใช้งานฟังก์ชัน familyName() ด้วยชื่อ jenny และอายุ 14
	//เรียกใช้งานฟังก์ชัน familyName() ด้วยชื่อ Anja และอายุ 30
	familyName("Liam", 3)   // ผลลัพธ์ Hi 3 year Liam Refsnes
	familyName("jenny", 14) // ผลลัพธ์ Hi 14 year jenny Refsnes
	familyName("Anja", 30)  // ผลลัพธ์ Hi 30 year Anja Refsnes
}

/*อธิบายโค้ด
ฟังก์ชัน familyName() มีพารามิเตอร์สองตัว fname ที่เป็น
string และ age ที่เป็น int เมื่อเรียกใช้ฟังก์ชันนี้
เราต้องส่งอากิวเมนต์สองตัวไปด้วยเช่น"Liam" และ 3,
"๋Jenny" และ14,"Anja" และ 30
*/
/*Note
เมื่อทำงานกับพารามิเตอร์หลายตัว การเรียกใช้ฟังก์ชันจะต้องมี
จำนวนอาร์กิวเมนต์เท่ากับจำนวนพารามิเตอร์ที่ฟังก์ชันรับ และต้อง
ส่งอาร์กิวเมนต์ในลำดับเดียวกับที่กำหนดพารามิเตอร์ในฟังก์ชัน
*/

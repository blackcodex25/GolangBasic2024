package main

import (
	f "fmt"
)

var a int     // ประกาศตัวแปร a ชนิด int โดยไม่กำหนดค่าเริ่มต้น
var b int = 2 // ประกาศตัวแปร b ชนิด int และกำหนดค่าเริ่มต้นเป็น 2
var c = 3     // ประกาศตัวแปร c โดยไม่ระบุชนิดตัวแปร แต่กำหนดค่าเริ่มต้นเป็น 3 (Go จะตัดสินใจชนิดตัวแปรให้โดยอัตโนมัติ)
// g := 1 // จะเกิดข้อผิดพลาดที่นี่
func main() {
	a = 1        // กำหนดค่าให้ตัวแปร a เป็น 1
	f.Println(a) // พิมพ์ค่าของ a
	f.Println(b) // พิมพ์ค่าของ b
	f.Println(c) // พิมพ์ค่าของ c
	//	f.Println(g) //syntax error: non-declaration statement outside function body
}

//var การระบุประเภทของตัวแปร การประกาศโดยไม่ต้องกำหนดค่าเริ่มต้น สามารถใช้ได้ทุกที่
//:= Operator ตัวดำเนินการ ใช้งานได้ภายในฟังก์ชันเท่านั้น ต้องกำหนดค่าเริ่มต้น การกำหนดประเภทอัตโนมัติ
/*สรุป
ใช้ var เมื่อคุณต้องการระบุประเภทของตัวแปรหรือเมื่อคุณต้องการประกาศตัวแปรโดยไม่กำหนดค่าเริ่มต้น และสามารถใช้ได้ทั้งในระดับ package และภายในฟังก์ชัน
ใช้ := เมื่อคุณต้องการกำหนดค่าตัวแปรพร้อมกับประกาศตัวแปร โดยไม่ต้องระบุประเภท และใช้ได้เฉพาะภายในฟังก์ชัน*/

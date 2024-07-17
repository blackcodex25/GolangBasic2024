package main

import (
	f "fmt"
)

/*การตรวจสอบว่ามีสมาชิกใน map หรือไม่ (Check For Specific Elements in a Map)
ในภาษา Go */
/*เราสามารถตรวจสอบว่ามีคีย์เฉพาะอยู่ใน map หรือไม่โดยใช้รูปแบบดังนี้
Syntax
val, ok := map_name[key]
ถ้าเราต้องการตรวจสอบแค่การมีอยู่ของคีย์เฉพาะ สามารถใช้ _ (blank identifier) แทน val ได้
*/
//  สร้าง maps ชื่อ a ที่เก็บข้อมูลประเภท string:string โดยมี key คือ "brand": "Ford", "model": "Mustang", "year": "1964", "day": ""
var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964", "day": ""}

func main() {
	// ตรวจสอบคีย์ "brand" คีย์ "brand" มีอยู่ในแผนที่ ดังนั้น val1 จะได้ค่า "Ford" และ ok1 จะได้ค่า true
	val1, ok1 := a["brand"] // ตรวจสอบคีย์ที่มีอยู่และค่าของมัน
	// ตรวจสอบคีย์ "color" คีย์ "color" ไม่มีอยู่ใน map ดังนั้น val2 จะได้ค่า "" (ค่าเริ่มต้นสำหรับ string) และ ok2 จะได้ค่า false
	val2, ok2 := a["color"] // ตรวจสอบคีย์ที่ไม่มีอยู่และค่าของมัน
	// ตรวจสอบคีย์ "day" คีย์ "day" มีอยู่ใน map แต่ค่าของมันเป็น "" ดังนั้น val3 จะได้ค่า "" และ ok3 จะได้ค่า true
	val3, ok3 := a["day"] // ตรวจสอบคีย์ที่มีอยู่และค่าของมัน
	//
	_, ok4 := a["model"] // ตรวจสอบแค่คีย์ที่มีอยู่โดยไม่สนใจค่าของมัน

	f.Println(val1, ok1)
	f.Println(val2, ok2)
	f.Print(val3, ok3)
	f.Println(ok4)
}

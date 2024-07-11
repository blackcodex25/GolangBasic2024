package main

import (
	f "fmt"
)

//การใช้คำสั่ง switch ในภาษา Go
//ใช้คำสั่ง switch เพื่อเลือกหนึ่งในหลายๆ บล็อกโค้ดที่จะถูกดำเนินการ
/* คำสั่ง switch ในภาษา Go คล้ายกับในภาษา C,C++,java,javascript,
และ PHP ความแตกต่างคือมันจะรันเฉพาะเคสที่ตรงกันเท่านั้น ไม่จำเป็นต้องใช้คำสั่ง
break
*/
/*Syntax ของ switch แบบ Single-Case
switch expression {
case x:
	บล็อกโค้ด
case y:
	บล็อกโค้ด
case z:
	บล็อกโค้ด
default:
	บล็อกโค้ด
}
*/
/*วิธีการทำงานของ Switch
1.ประเมินค่า expression หนึ่งครั้ง
2.ค่าของ switch expression จะถูกเปรียบเทียบกับค่าของแต่ละ case
3.หากมีการตรงกัน บล็อกโค้ดที่เกี่ยวข้องจะถูกดำเนินการ
4.คำสั่ง default เป็นตัวเลือก มันจะระบุโค้ดที่จะรันหากไม่มีเคสใดตรงกัน
*/
//ตัวอย่าง switch แบบ Single-Case
//โค้ดตัวอย่างด้านล่างใช้หมายเลขของวันในสัปดาห์เพื่อคำนวณชื่อวัน
func main() {
	day := 4

	switch day {
	case 1:
		f.Println("Monday")
	case 2:
		f.Println("Tuesday")
	case 3:
		f.Println("Wednesday")
	case 4:
		f.Println("Thursday")
	case 5:
		f.Println("Friday")
	case 6:
		f.Println("Saturday")
	case 7:
		f.Println("Sunday")
	default:
		f.Println("วันที่ไม่ถูกต้อง")
	}
}

/*อธิบายโค้ด
1.ประกาศตัวแปร day และกำหนดค่าเป็น 4
2.ใช้คำสั่ง switch เพื่อตรวจสอบค่า day
3.ตรวจสอบเคส
หาก day เท่ากับ 1 จะแสดงผล "Monday"
หาก day เท่ากับ 2 จะแสดงผล "Tuesday"
หาก day เท่ากับ 3 จะแสดงผล "Wednesday"
หาก day เท่ากับ 4 จะแสดงผล "Thursday"
หาก day เท่ากับ 5 จะแสดงผล "Friday"
หาก day เท่ากับ 6 จะแสดงผล "Saturday"
หาก day เท่ากับ 7 จะแสดงผล "Sunday"
หากไม่มีเคสใดตรงกัน (กรณี default) จะแสดงผล "วันที่ไม่ถูกต้อง"

ในโค้ดตัวอย่างนี้ ค่า day เท่ากับ 4 ดังนั้นโปรแกรมจะแสดงผล "Thursday"
*/
/*สรุป
คำสั่ง switch ในภาษา Go ช่วยให้สามารถจัดการกับหลายๆ
เงื่อนไขได้อย่างเป็นระเบียบและเข้าใจง่าย โดยไม่ต้องใช้คำสั่ง
break เหมือนในภาษาการเขียนโปรแกรมอื่นๆ
*/

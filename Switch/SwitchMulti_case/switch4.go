package main

import (
	f "fmt"
)

//Go Multi-case switch Statement
/*ในภาษา Go เราสามารถกำหนดค่าหลายค่าในแต่ละเคสของคำสั่ง switch ได้
ซึ่งช่วยให้การเขียนโปรแกรมมีความยืดหยุ่นและสะดวกมากขึ้น
*/
/*Syntax switch expression {
case x, y:
	โค้ดที่จะถูกเรียกใช้ถ้า expression เท่ากับ x หรือ y
case v, w:
	โค้ดจะถูกเรียกใช้ถ้า expression	เท่ากับ v หรือ w
case z:
	โค้ดที่จะถูกเรียกใช้ถ้า expression เท่ากับ z
default:
	โค้ดที่จะถูกเรียกใช้ถ้า expression ไม่ตรงกับเคสใด ๆ
}
*/
//โค้ดตัวอย่างการใช้ Multi-case switch
//เราใช้ตัวเลขของวันในสัปดาห์เพื่อคืน ข้อความที่แตกต่างกันตามวันนั้นๆ

func main() {
	day := 5
	switch day {
	case 1, 3, 5: // หาก day เท่ากับ 1, 3, หรือ 5 จะแสดงผล "Odd Weekday"
		f.Println("Odd weekday")
	case 2, 4: // หาก day เท่ากับ 2 หรือ 4 จะแสดงผล "Even weekday"
		f.Println("Even weekday")
	case 6, 7: // หาก day เท่ากับ 6 หรือ 7 จะแสดงผล "Weekend"
		f.Println("Weekend")
	default: // หากไม่มีเคสใดตรงกัน (กรณี default) จะแสดงผล "วันในสัปดาห์ไม่ถูกต้อง"
		f.Println("วันในสัปดาห์ไม่ถูกต้อง")
	}
}

/*อธิบายโค้ด
1.ประกาศตัวแปร day และกำหนดค่าเป็น 5
2.ใช้คำสั่ง switch เพื่อตรวจสอบค่า day
3.ตรวจสอบเคส
หาก day เท่ากับ 1, 3, หรือ 5 จะแสดงผล "Odd Weekday"
หาก day เท่ากับ 2 หรือ 4 จะแสดงผล "Even weekday"
หาก day เท่ากับ 6 หรือ 7 จะแสดงผล "Weekend"
หากไม่มีเคสใดตรงกัน (กรณี default) จะแสดงผล "วันในสัปดาห์ไม่ถูกต้อง"

โค้ดตัวอย่างนี้ ค่า day เท่ากับ 5 ซึ่งตรงงกับเคส 1, 3, หรือ 5 ดังนั้น
โปรแกรมจะแสดงผล "Odd weekday"
*/
/*สรุป
ในแต่ละเคสสามารถมีค่าหลายค่าได้
ซึ่งช่วยให้สามารถจัดการกับหลายเงื่อนไขได้ในเคสเดียวกัน
default เคสจะถูกเรียกใช้เมื่อไม่มีเคสใดตรงกับ expression
*/

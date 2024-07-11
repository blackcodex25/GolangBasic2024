package main

import (
	f "fmt"
)

// คำสำคัญ default ใช้เพื่อระบุโค้ดที่จะรันหาก
// ไม่มีเคสใดตรงกับ expression ที่ประเมินผล
func main() {
	day := 8

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
		f.Println("ไม่ใช่วันธรรมดา")

	}
}

/*อธิบายโค้ด
1.ประกาศตัวแปร day และกำหนดค่าเป็น 8
2.ใช้คำสั่ง switch เพื่อตรวจสอบค่า day
3.ตรวจสอบเคส
หาก day เท่ากับ 1 จะแสดงผล "Monday"
หาก day เท่ากับ 2 จะแสดงผล "Tuesday"
หาก day เท่ากับ 3 จะแสดงผล "Wednesday"
หาก day เท่ากับ 4 จะแสดงผล "Thursday"
หาก day เท่ากับ 5 จะแสดงผล "Friday"
หาก day เท่ากับ 6 จะแสดงผล "Saturday"
หาก day เท่ากับ 7 จะแสดงผล "Sunday"
หากไม่มีเคสใดตรงกัน (กรณี default) จะแสดงผล "ไม่ใช่วันธรรมดา"

โค้ดตัวอย่างนี้ ค่า day เท่ากับ 8 ซึ่งไม่มีเคสที่ตรงกัน ดังนั้น
โปรแกรมจะแสดงผล "ไม่ใช่วันธรรมดา"
*/

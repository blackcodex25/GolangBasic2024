package main

import (
	f "fmt"
	"time"
)

/*
	โค้ดตัวอย่างนี้ใช้ในการบอกว่าจากวันนี้วันไหนจะถึงวันเสาร์
	โดยใช้คำสั่ง switch และนำเข้าแพ็กเกจ time มาใช้งาน

switch เพื่อประเมินว่าวันเสาร์จะถึงในอีกกีี่วัน
*/
func main() {
	f.Println("เมื่อไหร่จะถึงวันเสาร์?") // พิมพ์ข้อความออกจอ
	today := time.Now().Weekday()        // ประกาศตัวแปร today และกำหนดค่าเป็นวันในสัปดาห์ปัจจุบัน โดยใช้ฟังก์ชัน time.Now().Weekday()
	// เงื่อนไขตรวจสอบวันปัจจุบัน
	switch time.Saturday { // ใช้คำสั่ง switch เพื่อเปรียบเทียบว่า time.Saturday (วันเสาร์) ตรงกับ case ใด
	case today + 0: // ถ้าวันนี้เป็นวันเสาร์ (today เท่ากับ time.Saturday) จะทำงานคำสั่งในบล็อกนี้
		f.Println("วันนี้") // พิมพ์ข้อความออกจอภาพ
	case today + 1: // ถ้าวันนี้เป็นวันก่อนวันเสาร์ 1 วัน (today + 1 เท่ากับ time.Saturday) จะทำงานคำสั่งในบล็อกนี้
		f.Println("พรุ่งนี้") // พิมพ์ข้อความออกจอภาพ
	case today + 2: // ถ้าวันนี้เป็นวันก่อนวันเสาร์ 2 วัน (today + 2 เท่ากับ time.Saturday)
		f.Println("ในสองวัน") // พิมพ์ข้อความออกจอภาพ
	default: // ถ้าวันเสาร์อยู่ห่างจากวันนี้มากกว่า 2 วัน จะทำงานคำสั่งในบล็อกนี้
		f.Println("อีกหลายวัน") // พิมพ์ข้อความออกจอภาพ
	}
} // Note: เลขหนึ่ง คือแทน วัน

/* การทำงานของโปรแกรม
1.โปรแกรมจะเริ่มต้นด้วยการพิมพ์ข้อความ "When's Saturday?"
2.จากนั้นตรวจสอบวันปัจจุบันด้วยฟังก์ชัน time.Now().Weekday()
3.ใช้คำสั่ง switch เพื่อเปรียบเทียบวันปัจจุบันกับวันเสาร์ (time.Saturday)
4.โปรแกรมจะพิมพ์ข้อความตามวันปัจจุบัน
- ถ้าวันนี้เป็นวันเสาร์ จะพิมพ์ข้อความ วันนี้
- ถ้าวันนี้เป็นวันศุกร์ จะพิมพ์ข้อความ พรุ่งนี้
- ถ้าวันนี้เป็นวันพฤหัสบดี จะพิมพ์ข้อความ ในสองวัน
- ถ้าวันปัจจุบันเป็นวันอื่นๆ จะพิมพ์ข้อความ อีกหลายวัน
*/

/*Note:
1.การประเมินค่าในคำสั่ง switch คำสั่ง switch ใน Go จะทำการประเมินค่าเงื่อนไขจากบนลงล่าง
และหยุดที่ case แรกที่ตรงเงื่อนไข โดยไม่จำเป็นต้องใช้คำสั่ง break
เหมือนในภาษาอื่นๆ เช่น C, C++, Java
2.การจัดการเวลาของ Go
เวลาใน Go จะช่วยให้การทำงานที่เกี่ยวกับวันและเวลาเป็นเรื่องง่ายขึ้น
ด้วยการใช้งานฟังก์ชันจากแพ็กเกจ time สามารถใช้ในการทดสอบโค้ดที่เกี่ยวข้องกับวันและเวลาได้
*/

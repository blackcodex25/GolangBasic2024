package main

import (
	f "fmt"
	"time"
)

/* โค้ดนี้เป็นตัวอย่างการใช้งาน select พร้อม default case
ในการทำงานกับ time channels เพื่อทำงานบางอย่างเป็นระยะๆ
และจะหยุดการทำงานเมื่อเวลาผ่านไปถึงจุดที่กำหนด
*/
// นำเข้าแพ็กเกจ time เพื่อใช้สำหรับการทำงานกับเวลา
func main() {
	// สร้างช่อง tick และ boom
	/* tick เป็นช่องที่ส่งค่าทุกๆ 100 มิลลิวินาที
	   boom เป็นช่องที่จะส่งค่าเพียงครั้งเดียวหลังจาก 500 มิลลิ
	   วินาที
	*/
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	// ใช้ลูป for เพื่อวนลูปการทำงานอย่างต่อเนื่อง
	for {
		// ใช้ select ภายในลูปเพื่อเลือกการดำเนินการ
		// ตามสถานะของ channels ตัวแปร tick และ boom
		select {

		/* กรณี tick ของ channels เมื่อ channels tick
		ส่งค่า (ทุกๆ 100 มิลลิวินาที), ข้อความ "tick"
		จะถูกพิมพ์ออกจอ
		*/
		case <-tick:
			f.Println("tick.")

		/* กรณี boom channel
		เมื่อ channels ตัวแปร boom ส่งค่า
		(หลังจาก 500 มิลลิวินาที), ข้อความ "BOOM!"
		จะถูกพิมพ์ออกทางหน้าจอและฟังก์ชัน main
		จะหยุดการทำงานด้วย return
		*/
		case <-boom:
			f.Println("BOOM!")
			return

		/* กรณี default
		หากไม่มี channels ใดส่งค่า (ไม่มี case ไหนพร้อมทำงาน)
		ข้อความจุด . จะถูกพิมพ์ออกจอและโปรแกรมจะหยุดพัก
		50 มิลลิวินาที
		*/
		default:
			f.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

/* การทำงานของโปรแกรม
1.สร้าง time channels
tick := time.Tick(100 * time.Millisecond)
boom := time.After(500 * time.Millisecond)
tick เป็น channels ที่จะส่งค่าเป็นระยะๆ ทุกๆ 100 มิลลิวินาที
boom เป็น channels ที่จะส่งค่าเพียงครั้งเดียวหลังจาก 500 มิลลิวินาที

2.ลูป for และ select
for {
	select {
case <-tick:
	fmt.Println("tick.")

case <-boom:
	fmt.Println("BOOM!")
	return
default:
	fmt.Println("    .")
	time.Sleep(50 * time.Millisecond)
	}
}
ใช้ลูป for เพื่อวนลูปการทำงานอย่างต่อเนื่อง
ใช้ select เพื่อเลือกการดำเนินการตามสถานะของ channels
tick และ boom

หาก tick ส่งค่า
ข้อความ "tick." จะถูกพิมพ์

หาก boom ส่งค่า
ข้อความ "BOOM!" จะถูกพิมพ์และฟังก์ชันจะหยุดการทำงาน

หากไม่มี channels ใดส่งค่า
ข้อความ " ." จะถูกพิมพ์และโปรแกรมจะหยุดพัก 50 มิลลิวินาที
*/
/* ผลลัพธ์ที่เป็นไปได้
ข้อความ " ." จะถูกพิมพ์ทุกๆ 50 มิลลิวินาทีเมื่อไม่มีการส่งค่าจาก
channels tick และ boom
ข้อความ "tick." จะถูกพิมพ์ทุกๆ 100 มิลลิวินาที
เมื่อ tick ส่งค่า
ข้อความ "BOOM!" จะถูกพิมพ์หลังจาก 500 มิลลิวินาที
เมื่อ boom ส่งค่า และโปรแกรมจะหยุดทำงาน
    .
    .
tick.
    .
    .
tick.
    .
    .
tick.
    .
    .
tick.
    .
    .
BOOM!
*/
/* สรุป
โค้ดนี้แสดงการใช้ select พร้อม default case เพื่อจัดการ
การส่งและรับค่าจาก time channels ในภาษา Go การใช้
default case ช่วยให้สามารถทำงานได้โดยไม่บล็อกการทำงาน
ของโปรแกรมเมื่อไม่มี channels ใดที่พร้อมทำงาน
การสร้าง time channels ทำให้สามารถกำหนด
การทำงานเป็นระยะๆ และการทำงานเมื่อเวลาผ่านไปถึงจุด
ที่กำหนดได้
*/

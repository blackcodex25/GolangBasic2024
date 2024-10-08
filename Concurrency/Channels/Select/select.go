package main

import (
	f "fmt"
)

/*
	คำสั่ง select ช่วยให้ Goroutine

รอการดำเนินการสื่อสารหลายรายการได้
คำสั่ง select จะบล็อกจนกว่าจะมีกรณีหนึ่งที่สามารถทำงานได้ จากนั้นจะ
ดำเนินการตามกรณีนั้น มันจะสุ่มเลือกหนึ่งกรณี ถ้ามีหลายกรณีที่พร้อมทำงาน
*/
/* โค้ดนี้แสดงการใช้ Goroutine และ select statement เพื่อจัดการกับ Channels
หลายตัวพร้อมกันในการสร้างลำดับฟีโบนัชชี (Fibonacci sequence) และตรวจสอบเงื่อนไข
เพื่อออกจากการทำงานเมื่อได้รับ channels ตัวแปร quit
*/
/* ประกาศฟังก์ชัน fibonacci กำหนดพารามิเตอร์สองตัว c และ quit ชนิด channels
ซึ่งมีชนิดข้อมูล int */
func fibonacci(c, quit chan int) {
	// ประกาศตัวแปรสองตัว x และ y กำหนดค่า 0 และ 1 ตามลำดับ
	x, y := 0, 1
	// ลูป for ทำงานตลอดไปจนกว่าจะได้รับ channels จากตัวแปร quit
	for {
		// ประกาศฟังก์ชัน select กำหนดเลือก case ที่พร้อมทำงาน
		select {
		// ถ้าพารามิเตอร์ c พร้อมรับค่า x จะถูกส่งไปยัง channel พารามิเตอร์ c
		case c <- x:
			// ทำการคำนวณค่าถัดไปของลำดับ fibonacci
			x, y = y, x+y
		// ถ้าได้รับสัญญาณจากช่อง quit (case <-quit:)
		case <-quit:
			// จะพิมพ์ "quit" และออกจากฟังก์ชัน
			f.Println("quit")
			return // หยุดการทำงาน โดยใช้ฟังก์ชัน return
		}
	}
}
func main() {
	// สร้าง Channels ตัวแปร c สำหรับส่งค่าลำดับฟิโบนัชชี
	c := make(chan int)
	// สร้าง Channels ตัวแปร quit สำหรับส่งสัญญาณหยุดการทำงาน
	quit := make(chan int)
	/* สร้าง Goroutine ที่จะรับค่าจาก channels ตัวแปร c
	และพิมพ์ค่าที่ได้รับออกจอ 10 ครั้ง
	หลังจากพิมพ์ครบ 10 ครั้ง จะส่งค่า 0 ไปยัง channels
	ตัวแปร quit เพื่อส่งสัญญาณหยุดการทำงาน
	*/
	// การประกาศ Anonymous Function
	go func() {
		/* ประกาศตัวแปร i กำหนดค่า 0
		ตรวจสอบเงื่อนไขตัวแปร i น้อยกว่า 10 หรือไม่
		เพิ่มค่า i++ หลังจบลูปแต่ละครั้ง */
		for i := 0; i < 10; i++ {
			/*ในแต่ละรอบของลูป Goroutine จะรับค่าจาก Channels ตัวแปร c (ด้วย <-c)
			และพิมพ์ค่าที่ได้รับออกจอ
			*/
			f.Println(<-c)
		}
		/* หลังจากพิมพ์ครบ 10 ครั้ง
		จะส่งค่า 0 ไปยัง channels ตัวแปร quit
		เพื่อส่งสัญญาณหยุดการทำงาน
		*/
		quit <- 0 // ส่งค่า 0 ไปยัง channels ตัวแปร quit เพื่อหยุดการทำงาน
	}()
	/*เรียกใช้ฟังก์ชัน fibonacci() เพื่อเริ่มการคำนวณและส่งค่าลำดับฟีโบนัชชี
	ไปยัง channels ตัวแปร c*/
	fibonacci(c, quit)

} /* Note: คำสั่ง go func() { ... }()
ใช้เพื่อสร้าง Goroutine ใหม่ที่
รันฟังก์ชัน anonymous (ฟังก์ชันที่ไม่มีชื่อ)
โค้ดภายในฟังก์ชัน anonymous นี้จะรันใน Goroutine แยกจาก
Goroutine หลัก
*/
/* การทำงานของฟังก์ชัน fibonacci และ select
ในฟังก์ชัน fibonacci
ใช้ select เพื่อเลือกการดำเนินการตามสถานะของ channels
		case c <- x: ส่งค่าฟีโบนัชชีปัจจุบันไปยัง channels ตัวแปร c
		case <-quit: รับสัญญาณหยุดการทำงานจาก channels
		quit, พิมพ์ข้อความ "quit" และหยุดการทำงาน
*/
/* สรุป
โค้ดนี้แสดงการใช้ Goroutine และ channels ในภาษา Go เพื่อสร้าง
และส่งค่าลำดับฟีโบนัชชี และการใช้ select เพื่อเลือกการดำเนิน
การตามสถานะของช่องสื่อสาร Goroutine หลักจะรันฟังก์ชัน
fibonacci และส่งค่าฟีโบนัชชีไปยัง channels ตัวแปร c
ในขณะที่ Goroutine ที่สร้างขึ้นจะรับค่าจาก channels ตัวแปร c
และพิมพ์ค่าที่ได้รับ หลังจากค่า 10 ครั้งจะส่งสัญญาณให้ Goroutine
หลักหยุดการทำงานโดยการส่งค่าไปยัง channels ตัวแปร quit
*/

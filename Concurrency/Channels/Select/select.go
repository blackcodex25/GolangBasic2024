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
เพื่อออกจากการทำงานเมื่อได้รับสัญญาณ quit
*/
/* ประกาศฟังก์ชัน fibonacci กำหนดพารามิเตอร์สองตัว c และ quit ชนิด channels
ซึ่งมีชนิดข้อมูล int */
func fibonacci(c, quit chan int) {
	// ประกาศตัวแปรสองตัว x และ y กำหนดค่า 0 และ 1 ตามลำดับ
	x, y := 0, 1
	// ลูป for ทำงานตลอดไปจนกว่าจะได้รับสัญญาณจากตัวแปร quit
	for {
		// ประกาศฟังก์ชัน select กำหนดเลือก case ที่พร้อมทำงาน
		select {
		// ถ้าพารามิเตอร์ c พร้อมรับค่า x จะถูกส่งไปยัง channel พารามิเตอร์ c
		case c <- x:
			// ทำการคำนวณค่าถัดไปของลำดับ fibonacci
			x, y = y, x+y
		// ถ้า quit ได้รับสัญญาณ
		case <-quit:
			// จะพิมพ์ "quit" และออกจากฟังก์ชัน
			f.Println("quit")
			return // คืนค่าให้กับ พารามิเตอร์ quit
		}
	}
}
func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			f.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)

}

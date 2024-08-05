package main

import (
	f "fmt"
	"time"
)

// นำเข้าแพ็กเกจ time เพื่อใช้ในการจัดการกับเวลา
/* ประกาศ struct ชื่อ MyError
ที่มีฟิลด์ When ชนิด time.Time และฟิลด์ What ชนิด string*/
type MyError struct {
	When time.Time
	What string
}

/*
	เมธอด Error ถูกประกาศสำหรับ pointer

ไปยัง MyError (*MyError)
*/
func (e *MyError) Error() string {
	/* เมธอดนี้จะคืนค่าสตริงที่บรรยายถึงข้อผิดพลาด โดยใช้
	ฟิลด์ When และ What ของ MyError
	*/
	return f.Sprintf("ที่ %v, %s", e.When, e.What)
}

// ประกาศฟังก์ชัน run() คืนค่า error
func run() error {
	/*ประกาศฟังก์ชัน return คืนค่า pointer
	ไปยัง MyError (*MyError) */
	return &MyError{
		// ฟิลด์ When ถูกกำหนดค่าเป็นเวลาปัจจุบัน (time.Now())
		time.Now(),
		// ฟิลด์ What ถูกกำหนดค่าเป็นสตริง "มันใช้งานไม่ได้"
		"มันใช้งานไม่ได้",
	}
}
func main() {
	// เรีียกใช้ฟังก์ชัน run() และตรวจสอบค่าคืน error
	if err := run(); err != nil {
		/* ถ้า error ไม่เป็น nil จะพิมพ์ค่าข้อผิดพลาดโดยใช้
		Println(err)
		*/
		f.Println(err)
	}
}

/* สรุป
- โปรแกรมนี้แสดงการสร้างและใช้งานค่า error
แบบกำหนดเองในภาษา Go โดยใช้ struct ชื่อ MyError
- ฟังก์ชัน run() คืนค่า error และในฟังก์ชันหลัก (main)
จะตรวจสอบและจัดการกับ error นั้น
- การใช้งาน type assertion และ type switch
ในภาษา Go ช่วยให้โปรแกรมสามารถจัดการกับค่าที่เป็นชนิด
interface{} ได้อย่างยืดหยุ่นและปลอดภัย
*/

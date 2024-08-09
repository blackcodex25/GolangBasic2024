package main

import (
	f "fmt"

	"golang.org/x/tour/tree"
)

/*
1.ฟังก์ชัน Walk
ฟังก์ชัน Walk ทำหน้าที่เด
*/
/* ประกาศฟังก์ชัน Walk() กำหนดพารามิเตอร์ t ชนิดแพ็กเกจ tree.Tree
ที่เรียกใช้สตรัค Tree และกำหนดพารามิเตอร์ของ channels ชื่อ ch */
func Walk(t *tree.Tree, ch chan int) {
	// ใช้ keyword defer เพื่อรอการทำงานในฟังก์ชันทั้งหมดเสร็จก่อนถึงจะเรียกใช้ ch
	defer close(ch) // ปิด channels หลังจาก traverse เสร็จสิ้น
	// นิยามฟังก์ชันช่วยเดินทางผ่านต้นไม้แบบ recursive ฟังก์ชันแบบเรียกซ้ำ
	var walkTree func(*tree.Tree)
	// เรียกใช้ฟังก์ชัน walkTree(t *tree.Tree) กำหนดพารามิเตอร์ t
	walkTree = func(t *tree.Tree) {
		if t == nil { // ตัวแปร t เท่ากับ nil ที่มีค่า 0
			return // ถ้าใช่ คืนค่า
		}
		walkTree(t.Left)  // เดินทางผ่านโหนดลูกทางซ้าย
		ch <- t.Value     // ส่งค่าจากโหนดปัจจุบันไปยังช่อง
		walkTree(t.Right) // เดินทางผ่านโหนดลูกทางขวา
	}
	walkTree(t) // เริ่มต้นการเดินทางจากรากต้นไม้
}

// ฟังก์ชัน Same() กำหนดพารามิเตอร์สองตัว t1 และ t2
// ชนิดข้อมูลต้นไม้ *tree.Tree คืนค่าชนิด bool
func Same(t1, t2 *tree.Tree) bool {
	// สร้าง channels สำหรับต้นไม้ t1
	// สร้าง channels สำหรับต้นไม้ t2
	ch1 := make(chan int)
	ch2 := make(chan int)
	/* เรียกใช้ฟังก์ชัน Walk() ใช้ goroutine
	เพื่อเดินทางผ่านต้นไม้ทั้งคู่พร้อมกัน */
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	// การเปรียบเทียบค่าในลูปจาก Channels ทั้งสอง
	// ตรวจสอบว่าค่าที่ได้รับจากทั้งสอง channels เท่ากันหรือไม่
	for v1 := range ch1 {
		v2, ok := <-ch2
		// ถ้าไม่เท่ากัน หรือ channels ชื่อ ch2 ปิดก่อนครบการเดินทาง
		if !ok || v1 != v2 {
			return false // ถ้าค่าหรือสถานะไม่ตรงกัน return คืนค่า false
		}
	}
	// ตรวจสอบว่า channels ชื่อ ch2 ถูกปิดแล้ว
	/* หลังจากเปรียบเทียบทุกค่าแล้ว ฟังก์ชันจะตรวจสอบว่า channels ชื่อ ch2
	ถูกปิดหรือไม่ ถ้าถูกปิดแล้ว คืนค่า true */
	_, ok := <-ch2
	return !ok
}

func main() {
	// สร้าง channels ชื่อ ch
	ch := make(chan int)
	// ทดสอบฟังก์ชัน Walk()
	// สร้างต้นไม้ใหม่และพิมพ์ค่า 10 ค่าแรกที่ได้รับจาก channels
	go Walk(tree.New(1), ch)
	// การใช้ for short statement ใช้ได้เพียงแค่ในบล็อก for เท่านั้น
	// ประกาศตัวแปร i กำหนดค่า 1 เงื่อนไขตรวจสอบ i น้อยกว่าเท่ากับ 10 หรือไม่ เพิ่มค่า i++ หลังจบลูปแต่ละครั้ง
	for i := 1; i <= 10; i++ {
		f.Println(<-ch)
	}
	// ทดสอบฟังก์ชัน Same()
	// โดยการเปรียบเทียบต้นไม้สองต้นที่มีค่าเหมือนกันและต่างกัน และพิมพ์ผลลัพธ์ออกมา
	f.Println(Same(tree.New(1), tree.New(1)))
	f.Println(Same(tree.New(1), tree.New(2)))
}

/* สรุป
ฟังก์ชัน Walk() เดินทางผ่านต้นไม้แบบ in-order และส่งค่าผ่าน channels
ฟังก์ชัน Same() ใช้ Walk() เพื่อเปรียบเทียบว่าต้นไม้สองต้นมีค่าที่เก็บอยู่เหมือนกันหรือไม่
ฟังก์ชัน main ใช้ทดสอบฟังก์ชันทั้งสองโดยการสร้างต้นไม้และตรวจสอบผลลัพธ์
*/

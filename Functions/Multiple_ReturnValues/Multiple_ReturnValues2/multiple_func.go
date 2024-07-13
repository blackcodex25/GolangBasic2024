package main

import (
	f "fmt"
)

//การจัดเก็บค่าที่คืนหลายค่าในตัวแปร
/*โค้ดตัวอย่างนี้ ฟังก์ชัน myFunction คืนค่าหลายค่า ซึ่งเราสามารถ
จัดเก็บค่าเหลานั้นลงในตัวแปรสองตัว a และ b
*/

// การสร้างฟังก์ชัน (Function Creation)
// ประกาศฟังก์ชันชื่อ myFunction() รับค่าอากิวเมนต์ x ชนิดข้อมูล int และ y ชนิดข้อมูล string
// ประกาศฟังก์ชันคืนค่าชื่อ result ชนิดข้อมูล int และตัวที่สองชื่อ txt1 ชนิดข้อมูล string
func myFunction(x int, y string) (result int, txt1 string) {
	result = x + x      // เรียกใช้งานฟังก์ชันคืนค่าชื่อ result รับค่าผลรวมจำนวนเต็ม x + x
	txt1 = y + "World!" // เรียกใช้งานฟังก์ชันคืนค่าชื่อ txt1 รับค่าผลรวมข้อความ y + "World!"
	return              // ประกาศฟังก์ชัน return ไม่ระบุค่าที่จะคืน จะคืนค่าผลรวมทั้ง result และ txt1 แล้วส่งออกไป
}

// การเรียกใช้ฟังก์ชันและจัดเก็บค่าที่คืน (Function Call and Storing Returned Values)
func main() {
	a, b := myFunction(5, "Hello ") // เก็บค่าที่ส่งคืนจาก myFunction ไว้ในตัวแปร a และ b
	f.Println(a, b)                 // พิมพ์ค่าของ a และ b
	//ผลลัพธ์ 10 Hello World!
}

/*อธิบายโค้ด
- ฟังก์ชัน myFunction รับค่าพารามิเตอร์สองค่า คือ x (ชนิด int) และ y (ชนิด string)
- ฟังก์ชันนี้คืนค่าผลลัพธ์สองค่าคือ result (ชนิด int) และ txt1 (ชนิด string)
- การใช้ return แบบไม่ระบุค่าค่าที่จะคืน จะคืนค่าทั้ง result และ txt1 ออกไป

- ในฟังก์ชัน main เราเรียกใช้ myFunction โดยส่งค่า 5 และ Hello เป็นอาร์กิวเมนต์
- ผลลัพธ์ที่คืนจาก myFunction จะถูกจัดเก็บในตัวแปร a และ b
- a จะได้รับค่า 10 (เนื่องจาก 5 + 5)
- b จะได้รับค่า "Hello world!"
- ผลลัพธ์ 10 Hello World!
*/

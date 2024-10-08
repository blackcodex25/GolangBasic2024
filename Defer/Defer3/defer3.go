package main

import (
	f "fmt"
) // นำเข้าแพ็กเกจ fmt

func main() {
	f.Println("Counting")     // พิมพ์ข้อความ counting
	for i := 0; i < 10; i++ { // ลูป for ที่ทำงาน 10 ครั้ง (0 ถึง 9)
		defer f.Println(i) // คำสั่ง defer จะทำให้การเรียกใช้ฟังก์ชัน Println() ถูกเลื่อนออกไปจนกว่าฟังก์ชัน main จะสิ้นสุดการทำงาน
		// การเรียกใช้ฟังก์ชัน defer จะถูกบันทึกไว้ใน stack ของการเรียกใช้ฟังก์ชัน
	}
	f.Println("done") // พิมพ์ข้อความ done ออกจอภาพ
}

/* ฟังก์ชัน main สิ้นสุดการทำงาน
- เมื่อฟังก์ชัน main สิ้นสุดการทำงาน คำสั่งที่ถูก defer
ไว้จะถูกเรียกใช้ตามลำดับย้อนกลับ
- เนื่องจากมีการ defer การเรียกใช้ Println(i) ในลูป for จำนวน 10 ครั้ง
คำสั่ง Println(i) จะถูกเรียกใช้ในลำดับจากมากไปน้อย (9,8,7,...,0)
*/
/* สรุป
- คำสั่ง defer จะบันทึกการเรียกใช้ฟังก์ชันไว้ใน stack
- การเรียกใช้ฟังก์ชันที่ถูก defer ไว้จะถูกดำเนินการในลำดับย้อนกลับ
เมื่อฟังก์ชันรอบๆ ส่งคืนการทำงาน
- โค้ดตัวอย่างนี้แสดงการใช้งาน defer ภายในลูป for และแสดงให้เห็น
ว่าคำสั่งที่ถูก defer ไว้จะถูกเรียกใช้ในลำดับ
ย้อนกลับเมื่อฟังก์ชัน main สิ้นสุดการทำงาน
*/

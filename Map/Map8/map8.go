package main

import (
	f "fmt"
)

// map เป็นการอ้างอิง (Maps Are References) ในภาษา Go
/* map ในภาษา Go เป็นการอ้างอิงถึงตารางแฮช (hash tables)
ซึ่งหมายความว่าถ้าตัวแปร map สองตัวอ้างอิงถึงตารางแฮชเดียวกัน
การเปลี่ยนแปลงเนื้อหาของตัวแปรหนึ่งจะส่งผลต่อเนื้อหา
ของอีกตัวแปรด้วย
*/
//ประกาศฟังก์ชัน map ชื่อ a กำหนดชนิดข้อมูล string:string และกำหนดค่าของ map ซึ่งคือ "brand": "Ford", "model": "Mustang", "year": "1964"
var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}

func main() {
	b := a                                   // ประกาศตัวแปร b อ้างอิงตารางแฮชของตัวแปร a
	f.Println(a)                             // พิมพ์ค่าเริ่มต้นข้อมูลตัวแปร a ออกจอภาพ
	f.Println(b)                             // พิมพ์ค่าเริ่มต้นข้อมูลตัวแปร b ที่อ้างอิงตารางแฮชของตัวแปร a ออกจอภาพ
	b["year"] = "1970"                       // เพิ่มค่าคีย์เข้าไปใน map
	f.Println("หลังจากเปลี่ยนแปลงค่าของ b:") // พิมพ์ผลลัพธ์ออกจอภาพ
	//พิมพ์ค่าของ map หลังการเปลี่ยนแปลงค่า
	f.Println(a) // พิมพ์ค่าข้อมูลตัวแปร a
	f.Println(b) // พิมพ์ค่าข้อมูลตัวแปร b ที่อ้างอิงตารางแฮชของตัวแปร a ออกจอภาพ
}

/* เนื่องจาก map ใน go เป็นการอ้างอิงถึงตารางแฮช (hash table)
การเปลี่ยนแปลงค่าของ map หนึ่งจะส่งผลต่ออีก map ที่อ้างอิงถึงตารางแฮช
ดังนั้น การเข้าใจลักษณะนี้จะช่วยให้เรา จัดการกับ map ใน Go ได้อย่างมีประสิทธิภาพ
และหลีกเลี่ยงข้อผิดพลาดที่อาจเกิดขึ้นจากการอ้างอิงแบบไม่ตั้งใจ
*/

package main

import (
	f "fmt"
)

/*การสร้างค่าว่างเปล่าของ Map Empty
มีสองวิธีในการสร้างค่าว่างเปล่าของ Map Empty ในภาษา Go วิธีแรกคือการใช้ฟังก์ชัน make()
และอีกวิธีคือการประกาศโดยตรง
*/
/* ไวยากรณ์คำสั่ง
var a map[Keytype]Valuetype
*/
/* Note
การใช้ฟังก์ชัน make() เป็นวิธีที่ถูกต้องในการสร้างค่าว่างเปล่าของ Map Empty
ถ้าเราสร้าง Map Empty โดยวิธีอื่นและเขียนค่าเข้าไป มันจะทำให้เกิด runtime panic
*/
/*โค้ดตัวอย่างนี้แสดงถึงความแตกต่างระหว่างการประกาศ map empty
โดยใช้ฟังก์ชัน make() และการประกาศโดยไม่ใช้ฟังก์ชัน make()
*/
// สร้าง map ว่างเปล่าโดยใช้ฟังก์ชัน make()

// ประกาศตัวแปร a และใช้ฟังก์ชัน make() กำหนดชนิดข้อมูลเป็น string:string
var a = make(map[string]string)

/* เราใช้ฟังก์ชัน make() เพื่อสร้าง map ของตัวแปร a ที่ว่างเปล่า
และพร้อมที่จะใช้งาน การสร้าง map โดยใช้ make() จะทำให้ map มีค่าเป็น non-nil */

// ประกาศฟังก์ชัน map ว่างเปล่าโดยไม่ใช้ฟังก์ชัน make()
var b map[string]string

/*
ประกาศคีย์เวิร์ด map ของตัวแปร b ที่ว่างเปล่าโดยไม่ใช้ฟังก์ชัน make()
การประกาศ map แบบนี้จะทำให้ map ของตัวแปร b มีค่าเป็น nil และยังไม่พร้อมใช้งาน
*/
func main() {
	// ตรวจสอบว่า map empty ว่างเปล่าหรือไม่
	// ผลลัพธ์จะเป็น false เพราะ a ถูกสร้างขึ้นด้วย make() และไม่ใช่ nil
	f.Println(a == nil)
	// ผลลัพธ์จะเป็น true เพราะ b ยังไม่ถูกสร้างด้วย make() และเป็น nil
	f.Println(b == nil)
	// *Note: nil ค่าที่ยังไม่ได้กำหนดให้ตัวแปรอื่นหรือยังไม่เปลี่ยนแปลงเป็น type อื่น มันจะไม่มีtype
}

/* Note:
การสร้าง map ด้วย make() เป็นวิธีที่ถูกต้องและปลอดภัยเพื่อหลีกเลี่ยงปัญหา runtime panic เมื่อคุณพยายามเขียนค่าเข้าไป
เข้าไปใน map ที่ยังไม่ได้สร้างขึ้นอย่างถูกต้อง
*/

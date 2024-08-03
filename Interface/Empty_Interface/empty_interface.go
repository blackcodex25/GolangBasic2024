package main

import (
	f "fmt"
)

/* ประเภทอินเทอร์เฟซที่ไม่ได้ระบุเมธอดใดๆ
เรียกว่าอินเทอร์เฟซว่างเปล่า
interface{}
อินเตอร์เฟซว่างสามารถถือค่าของประเภทใดก็ได้
(ทุกประเภทจะมีอย่างน้อยเมธอดศูนย์ตัว)
อินเทอร์เฟซว่างเปล่าใช้ในโค้ดที่จัดการกับค่าที่ไม่ทราบประเภทล่วง
หน้า ตัวอย่างเช่น fmt.Print รับพารามิเตอร์ได้หลายตัว
ที่มีประเภทเป็น interface{}
*/
// *อินเตอร์เฟซว่าง interface{} สามารเก็บค่าของชนิดข้อมูลใดก็ได้
// ประกาศตัวแปรชื่อ i ชนิดอินเตอร์เฟซว่าง interface{} ยังไม่ได้กำหนดค่า
var i interface{} // ค่าเริ่มต้นคือ nil
type A struct {
	S int
}

func main() {
	// ประกาศฟังก์ชัน describe() ค่าเริ่มต้นของ i
	describe(i) // ผลลัพธ์พิมพ์ค่าเริ่มต้นของ i และชนิดข้อมูลของ i
	i = 42      // เรียกใช้ตัวแปร i กำหนดค่า 42 ชนิด int
	// ประกาศฟังก์ชัน describe() กำหนดค่า i
	describe(i) // พิมพ์ค่าของ i และชนิดข้อมูลของ i หลังถูกกำหนดค่า

	// *การใช้งานอินเตอร์เฟซว่าง interface{} กำหนดค่าและชนิดข้อมูลต่างกัน

	// เรียกใช้ตัวแปร i กำหนดค่าเป็น hello ซึ่งเป็นชนิด string
	i = "hello"
	describe(i) //พิมพ์ค่าของ i และชนิดข้อมูลของ i หลังถูกกำหนดค่า
	// เรียกใช้ตัวแปร i กำหนดค่า 3.14 ซึ่งเป็นชนิด float64
	i = 3.14
	describe(i) // พิมพ์ค่าของ i และชนิดข้อมูลของ i หลังถูกกำหนดค่า

}

// ประกาศฟังก์ชันชื่อ describe() กำหนดพารามิเตอร์ชื่อ i ชนิดข้อมูลอินเตอร์เฟซว่าง interface{}
func describe(i interface{}) {
	// ประกาศฟังก์ชัน Printf() ไม่ขึ้นบรรทัดใหม่ ใช้สำหรับแสดงผลค่าและชนิดของตัวแปร
	f.Printf("(%v, %T)\n", i, i) // พิมพ์ค่าของ i และชนิดข้อมูลของ i
} /* %v แสดงค่าของตัวแปร
%T แสดงชนิดข้อมูลของตัวแปร
*/
/* สรุป
โปรแกรมนี้แสดงการใช้งานชนิดข้อมูลอินเตอร์เฟซว่าง interface{}
ในภาษา Go ซึ่งสามารถเก็บค่าของชนิดข้อมูลใดก็ได้
ฟังก์ชัน describe()
ใช้เพื่อพิมพ์ค่าของตัวแปร i และชนิดข้อมูลของ i การใช้งาน
อินเตอร์เฟซ interface{}
ทำให้โปรแกรมมีความยืดหยุ่นในการจัดการกับข้อมูล
ที่มีชนิดแตกต่างกันในตัวแปรเดียว
*/

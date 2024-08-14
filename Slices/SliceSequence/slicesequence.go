package main

import (
	"fmt"
	"slices"

	slicesequence2 "github.com/blackcodex25/GolangBasic2024/Slices/SliceSequence/SliceSequence2"
)

/*
	เริ่มจาก package iter หรือ iterators

พวก type ของ Seq และ Seq2

ซึ่งสามารถสร้างมาจากข้อมูลประเภท slice และ map
Seq คือ sequence ของ value เท่านั้น
Seq2 คือ sequence ของ Key, Value
*/

// นำเข้าแพ็กเกจ slices เพื่อเรียกใช้งานฟังก์ชันที่เกี่ยวข้องกับการจัดการ slice
// นำเข้าแพ็กเกจ slicesequence2 ซึ่งเป็นแพ็กเกจที่เราสร้างไว้เอง เพื่อเรียกใช้ฟังก์ชัน Sort()
func main() {
	// ประกาศตัวแปรของสไลซ์ชื่อ words ที่ประกอบด้วย string หลายตัว
	words := []string{"Hello", "Go", "1.23"}

	// การใช้ slices.All
	/* ฟังก์ชัน slices.All คืนค่าตัววนซ้ำ (iterator) ของ slice ที่มีทั้ง index (i)
	และค่า (v) ของแต่ละ element ใน slice
	วนลูปและพิมพ์ index และ value ของแต่ละ element ใน slice ชื่อ words
	*/
	for i, v := range slices.All(words) {
		fmt.Println(i, v)
	}

	// การใช้ slice.Values
	/* ฟังก์ชัน slices.Values คืนค่าตัววนซ้ำที่ประกอบด้วยค่า (value)
	ของแต่ละ element ใน slice
	วนลูปและพิมพ์ค่า (value) ของแต่ละ element ใน slice ชื่อตัวแปร words
	*/
	for v := range slices.Values(words) {
		fmt.Println(v)
	}

	/* การใช้ slice.Backward
	ฟังก์ชัน slice.Backward คืนค่าตัววนซ้ำที่วนลูปจากท้ายไปหน้า
	(reverse order) ใน slice
	วนลูปและพิมพ์ index และ value ของแต่ละ element ใน slice ชื่อ words
	โดยเริ่มจากท้ายสุดไปยังตัวแรกสุด
	*/
	for i, v := range slices.Backward(words) {
		fmt.Println(i, v)
	}

	// Values() ส่งคืนตัววนซ้ำของแต่ละ elements ใน slice
	/* ฟังก์ชัน slices.AppendSeq ใช้เพื่อเพิ่มค่าจากตัววนซ้ำ iterator (extends)
	ลงใน slice ที่มีอยู่แล้ว (words)
	สร้างตัววนซ้ำ extends จาก slice ที่มี string New และ World
	รวม extends เข้ากับ words และจัดเก็บผลลัพธ์ไว้ใน s
	วนลูปและพิมพ์ index และ value ของ slice ใหม่ s
	*/
	extends := slices.Values([]string{"New", "World"})
	// AppendSeq() เพิ่มค่าจากตัววนซ้ำ (iterator) ลงใน slice ที่มีอยู่แล้ว
	s := slices.AppendSeq(words, extends)
	for i, v := range s {
		fmt.Println(i, v)
	}

	// เรียกใช้ฟังก์ชัน Sort() จากโฟล์เดอร์ SliceSequence2
	slicesequence2.Sort() // พิมพ์ผลลัพธ์ออกจอ

}

/* สรุป
โค้ดนี้แสดงการใช้งานแพ็กเกจ slices ใน Go ซึ่งเป็นชุดเครื่องมือที่ใช้ในการจัดการ slice
ต่างๆ รวมถึงการสร้าง iterator สำหรับ slice การวนซ้ำไปข้างหน้าและย้อนกลับ และ
การเพิ่มสมาชิกใน slice
นอกจากนี้ยังมีการนำเข้าและเรียกใช้ฟังก์ชันจากแพ็กเกจที่สร้างขึ้น (slicesequence2.Sort())
*/

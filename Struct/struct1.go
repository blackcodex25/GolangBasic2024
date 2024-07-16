package main

import (
	f "fmt"
)

/* โครงสร้างข้อมูลในภาษา Go (Go Structures)
โครงสร้างข้อมูล (strcut) ในภาษา Go ใช้สำหรับการสร้างชุดของสมาชิกที่มี
ชนิดข้อมูลต่างกัน ให้เป็นตัวแปรเดียว โดย struct สามารถมีประโยชน์ในการรวม
ข้อมูลเข้าด้วยกันเพื่อสร้างบันทึกข้อมูล
*/
/* ประกาศโครงสร้าง (Declare a Struct)
ในการประกาศโครงสร้างใในภาษา Go ใช้คีย์เวิร์ด type และ struct ตามไวยากรณ์
*/
/*โครงสร้างคำสั่ง struct
type struct_name struct {
	member1 datatype
	member2 datatype
	member3 datatype
}
*/
/*การประกาศ Struct
เราประกาศชนิดข้อมูล struct ชื่อ person ซึ่งมีสมาชิกดังนี้
name age job และ salary
*/
/* การประกาศสมาชิกของ Struct:
สมาชิก name และ job เป็นชนิดข้อมูล string
สมาชิก age และ salary เป็นชนิดข้อมูล int
*/
//ประกาศ struct ชื่อ Person
type Person struct {
	name   string
	age    int
	job    string
	salary int
} // สร้างชนิดข้อมูล struct ที่ชื่อ Person ที่มีสมาชิก 4 ตัว: name, age, job และ salary

func main() {
	// สร้างตัวแปร p1 ของชนิด Person และกำหนดค่าให้กับสมาชิก
	p1 := Person{
		name:   "John Doe",
		age:    24,
		job:    "Software Engineer",
		salary: 50000,
	} // สร้างตัวแปร p1 และกำหนดค่าให้กับสมาชิกแต่ละตัวของ Person

	// เข้าถึงและพิมพ์ค่าสมาชิกของ p1
	f.Println("Name:", p1.name)
	f.Println("Age:", p1.age)
	f.Println("Job:", p1.job)
	f.Println("Salary:", p1.salary)
	// เข้าถึงและพิมพ์ค่าสมาชิก name, age, job และ salary ของ p1
}

/* ประโยชน์ของ struct
- ใช้สำหรับการรวมข้อมูลหลายๆ ชนิดเข้าด้วยกัน
- สะดวกในการจัดการและเข้าถึงข้อมูลที่เกี่ยวข้องกัน
Struct ในภาษา Go จึงเป็นเครื่องมือที่มีประโยชน์มากสำหรับ
การสร้างบันทึกข้อมูลที่ซับซ้อนและเป็นระเบียบ
*/

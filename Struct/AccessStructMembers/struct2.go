package main

import (
	f "fmt"
)

// การเข้าถึงสมาชิกของ Struct
/*ในการเข้าถึงสมาชิกของโครงสร้าง (struct) ใช้เครื่องหมายจุด (.)
ระหว่างชื่อตัวแปรโครงสร้างและสมาชิกของโครงสร้าง
*/
// ประกาศ struct ชื่อ Person
type Person struct {
	name   string
	age    int
	job    string
	salary int
}                // สร้างชนิดข้อมูล struct ที่ชื่อ Person ที่มีสมาชิก 4 ตัว: name age job และ salary
var pers1 Person // ประกาศตัวแปรชื่อ pers1 ชนิดข้อมูลของ struct ชื่อ Person
var pers2 Person // ประกาศตัวแปรชื่อ pers2 ชนิดข้อมูลของ struct ชื่อ Person
func main() {

	// การกำหนดค่าสำหรับ pers1
	pers1.name = "Hege"
	pers1.age = 45
	pers1.job = "Teacher"
	pers1.salary = 6000
	// กำหนดค่าของสมาชิก name, age, job และ salary ให้กับ pers1

	// การกำหนดค่าสำหรับ pers2
	pers2.name = "Cecilie"
	pers2.age = 24
	pers2.job = "Marketing"
	pers2.salary = 4500
	// กำหนดค่าของสมาชิก name, age, job และ salary ให้กับ pers2

	// การเข้าถึงและพิมพ์ข้อมูลของ pers1
	f.Println("Name: ", pers1.name)
	f.Println("Age: ", pers1.age)
	f.Println("Job: ", pers1.job)
	f.Println("salary: ", pers1.salary)
	// เข้าถึงและพิมพ์ค่าสมาชิก name, age, job และ salary ของ pers1

	// การเข้าถึงและพิมพ์ข้อมูลของ pers2
	f.Println("Name: ", pers2.name)
	f.Println("Age: ", pers2.age)
	f.Println("Job: ", pers2.job)
	f.Println("salary: ", pers2.salary)
	// เข้าถึงและพิมพ์ค่าสมาชิก name, age, job และ salary ของ pers2

}

/* สรุป
เราได้ประกาศ struct ชื่อ Person และสร้างตัวแปรสองตัวของชนิด Person จากนั้น
เรากำหนดค่าให้กับสมาชิกของ แต่ละตัวแปรและพิมพ์ข้อมูลของพวกเข้าออกมา การเข้าถึงสมาชิก
ของ struct ทำได้โดยการใช้เครื่องหมายจุด (.) ระหว่างชื่อตัวแปรและชื่อสมาชิกของ struct
*/

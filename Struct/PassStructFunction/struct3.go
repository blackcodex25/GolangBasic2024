package main

import (
	f "fmt"
)

// ประกาศโครงสร้าง struct ขื่อ Person ที่มีสมาชิก 4 ตัวคือ name age job และ salary
type Person struct {
	name   string
	age    int
	job    string
	salary int
}

// ประกาศฟังก์ชัน printPerson
func printPerson(pers Person) {
	f.Println("Name:", pers.name)
	f.Println("Age: ", pers.age)
	f.Println("Job: ", pers.job)
	f.Println("Salary: ", pers.salary)
} /* ฟังก์ชัน PrintPerson รับพารามิเตอร์ชนิด Person และ
พิมพ์ข้อมูลของสมาชิก name, age, job และ salary ของ struct Person*/

var pers1 Person // ประกาศตัวแปรชื่อ pers1 ชนิดข้อมูล Person
var pers2 Person // ประกาศตัวแปรชื่อ pers2 ชนิดข้อมูล Person

func main() {

	// กำหนดค่าสำหรับ pers1
	pers1.name = "Hege"
	pers1.age = 45
	pers1.job = "Teacher"
	pers1.salary = 6000
	// กำหนดค่าสมาชิก name, age, job และ salary ให้กับ pers1

	//การกำหนดค่าสำหรับ pers2
	pers2.name = "Cecillie"
	pers2.age = 24
	pers2.job = "Marketing"
	pers2.salary = 4500
	// กำหนดค่าสมาชิก name, age, job และ salary ให้กับ pers2

	// เรียกใช้ฟังก์ชัน printPerson เพื่อพิมพ์ข้อมูลของ pers1
	printPerson(pers1) // ส่งค่า pers1 เป็นพารามิเตอร์ให้กับฟังก์ชัน printPerson

	// เรียกใช้ฟังก์ชัน printPerson เมื่อพิมข้อมูลของ pers2
	printPerson(pers2) // ส่งค่า pers2 เป็นพารามิเตอร์ให้กับฟังก์ชัน printPerson

}

/*สรุป
เราได้ประกาศ struct ชื่อ Person และสร้างตัวแปรสองตัวของชนิด Person จากนั้น
เรากำหนดค่าให้กับสมาชิกของแต่ละตัวแปรและพิมพ์ข้อมูลของพวกเขาออกมาโดย
การส่งค่าตัวแปร struct เป็นพารามิเตอร์ให้กับฟังก์ชัน printPerson
*/

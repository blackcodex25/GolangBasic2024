package main

import (
	f "fmt"
)

/*
ตัวดำเนินการทางคณิตศาสตร์
ตัวดำเนินการทางคณิตศาสตร์ใช้สำหรับการปฏิบัติการทางคณิตศาสตร์ทั่วไป
ในภาษา Go
*/
var a = 10
var b = 5

func main() {
	// การบวก
	f.Println("a + b =", a+b) // 15

	// การลบ
	f.Println("a - b =", a-b) // 5

	// การคูณ
	f.Println("a * b =", a*b) // 50

	// การหาร
	f.Println("a / b =", a/b) // 2

	// การหารเอาเศษ
	f.Println("a % b =", a%b) // 0

	// การเพิ่มค่า 1
	a++
	f.Println("a++ =", a) // 11

	// การลดค่า 1
	b--
	f.Println("a-- =", b) // 4

}

/*รายการตัวดำเนินการทางคณิตศาสตร์
+ : บวก (Addition)
- : ลบ (Subtraction)
* : คูณ (Multiplication)
/ : หาร (Division)
% : หารเอาเศษ (Modulus)
++ : เพิ่มค่า 1 (Increment)
-- : ลดค่า 1 (Decrement)
*/

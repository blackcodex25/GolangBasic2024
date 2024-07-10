// Go กฎในการตั้งชื่อตัวแปร
package main

import (
	f "fmt"
)

//var ใช้ได้ทั้งนอกและในฟังก์ชัน
/*age, carName, _underscoreVariable, userID123, myFloat32Value
เป็นชื่อตัวแปรที่ถูกต้องตามกฎของ Go*/
var age int
var carName string
var _underscoreVariable int
var userID123 int
var myFloat32Value float32

func main() {
	//แต่ละตัวแปรถูกใช้เพื่อเก็บข้อมูลประเภทที่ต่างกันและถูกพิมพ์ออกจอภาพ
	age = 26
	carName = "Toyota"
	_underscoreVariable = 100
	userID123 = 123
	myFloat32Value = 3.14
	f.Println(age, carName, _underscoreVariable, userID123, myFloat32Value)
}

//การตั้งชื่อตัวแปรที่เหมาะสมช่วยให้โค้ดมีความอ่านง่ายและเข้าใจได้ดีขึ้น
//เพราะฉะนั้นควรเลือกใช้ชื่อที่สื่อความหมายได้ชัดเจน ตามประเภทข้อมูลที่ต้องการเก็บในตัวแปรนั้น ๆ

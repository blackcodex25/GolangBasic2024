package main

import (
	f "fmt"
)

// Syntax Variable Type
// var variablename type = value
// var ชื่อตัวแปร ประเภทของตัวแปร = ค่า
// variablename := value
// ชื่อตัวแปร := ค่า
func main() {
	// ไม่สามารถประกาศตัวแปรโดยใช้ := โดยไม่กำหนดค่าให้กับตัวแปรได้
	/*dev := "Programing"
	age := 1337
	f.Println("ตำแหน่ง", dev, "อายุ", age)*/
	var student1 string = "John" // ประเภทของตัวแปรเป็น string
	var student2 = "Jane"        // = Go จะตั้งประเภทของตัวแปรโดยอัตโนมัติตามค่าที่กำหนด
	x := 2                       // เพื่อประกาศตัวแปรและกำหนดค่าให้กับมันพร้อมๆ กันโดยไม่ต้องระบุประเภทของตัวแปร

	f.Println(student1)
	f.Println(student2)
	f.Println(x)
}

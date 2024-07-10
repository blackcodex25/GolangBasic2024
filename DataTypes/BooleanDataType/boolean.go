package main

import (
	f "fmt"
)

// ประเภทข้อมูล boolean ใน Go
// สามารถเก็บค่าได้เพียงแค่ true หรือ false เท่านั้น
var b1 bool = true // การประกาศแบบกำหนดชนิดข้อมูลและค่าเริ่มต้น
var b2 = true      // การประกาศแบบไม่กำหนดชนิดข้อมูลแต่กำหนดค่าเริ่มต้น
var b3 bool        // การประกาศแบบไม่กำหนดชนิดข้อมูลและกำหนดค่าเริ่มต้น

func main() {
	b4 := true    //การประกาศแบบไม่กำหนดชนิดข้อมูลและกำหนดค่าเริ่มต้น
	f.Println(b1) // ผลลัพธ์ true
	f.Println(b2) // ผลลัพธ์ true
	f.Println(b3) // ผลลัพธ์ false (ค่าเริ่มต้นของ boolean คือ false)
	f.Println(b4) // ผลลัพธ์ true
}

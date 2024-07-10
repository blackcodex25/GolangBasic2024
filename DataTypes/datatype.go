package main

import (
	f "fmt"
)

//การใช้งานและแสดงรายละเอียดของ Go Data Types (ประเภทข้อมูลใน Go)
//Go Data Types
/*
bool: แทนค่าบูลีน (boolean) ซึ่งมีค่าเป็น true หรือ false
Numeric: แทนค่าทางตัวเลขทั้งแบบจำนวนเต็ม (integer), ทศนิยม (floating point), และค่าซับซ้อน (complex)
string: แทนค่าข้อความ (string)
*/
var a bool = true    //ตัวแปรชนิด boolean
var b int = 5        // ตัวแปรชนิด integer
var c float32 = 3.14 // floating point number
var d string = "Hi!" // ตัวแปรชนิด string
func main() {
	f.Println("Boolean: ", a)
	f.Println("Integer: ", b)
	f.Println("float:   ", c)
	f.Println("String:  ", d)
}

/*
การใช้งานประเภทข้อมูลใน Go ช่วยให้เราสามารถจัดการข้อมูลต่าง ๆ
ได้อย่างมีประสิทธิภาพและแม่นยำตามความต้องการของแต่ละประเภทข้อมูล
*/

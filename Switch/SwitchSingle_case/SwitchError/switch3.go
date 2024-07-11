package main

import (
	f "fmt"
)

// ข้อกำหนดของชนิดข้อมูล
/*ค่าของแต่ละเคสควรมีชนิดข้อมูลเหมือนกับ expression ของ switch
หากไม่ตรงกันจะทำให้เกิดข้อผิดพลาดจากคอมไพเลอร์
*/
//โค้ดตัวอย่างข้อผิดพลาด
func main() {
	a := 3
	switch a {
	case 1:
		f.Println("a คือ 1")
	case "b":
		f.Println("a คือ b")
	} //./switch3.go:17:7: cannot convert "b" (untyped string constant) to type int

}

/*โค้ดตัวอย่างนี้ ตัวแปร a มีชนิดข้อมูลเป็น int แต่เคส b มีชนิดข้อมูลเป็น string
ซึ่งทำให้เกิดข้อผิดพลาดจากคอมไพเลอร์
*/

package main

import (
	f "fmt"
)

//ฟังก์ชันที่ใช้กับ Slices ใน Go

/* 1. ฟังก์ชัน len()
	ใช้เพื่อคืนค่าความยาวของ slice (จำนวนสมาชิกใน slice)
   2. ฟังก์ชัน cap()
	ใช้เพื่อคืนค่าความจุของ slice (จำนวนสมาชิกที่ slice สามารถขยายหรือหดได้)
*/
//ตัวอย่างการใช้งาน []datatype{values} และการใช้ฟังก์ชัน len() และ cap()

func main() {
	myslice := []int{}      //ประกาศ slice ว่างเปล่า
	f.Println(len(myslice)) //พิมพ์ค่าความยาวของ myslice1
	f.Println(cap(myslice)) //พิมพ์ความจุของ myslice1
	f.Println(myslice)      //พิมพ์ค่าใน myslice1

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"} //ประกาศและกำหนดค่าเริ่มต้นให้กับ slice
	f.Println(len(myslice2))                                //พิมพ์ความยาวของ myslice2
	f.Println(cap(myslice2))                                //พิมพ์ความจุของ myslice2
	f.Println(myslice2)                                     //พิมพ์ค่าใน myslice2
}

/* ประกาศและกำหนดค่าเริ่มต้นให้กับ slice
myslice2 :=[]string{"Go","slices","Are","Powerful"}
  - ประกาศ slice ที่ชื่อ myslice2 เป็น string และมีค่าเริ่มต้นคือ
    "Go","Slices","Are",และ "Powerful"
f.Println(len(myslice2))
  - พิมพ์ความยาวของ myslice2 ซึ่งคือ 4
f.Println(cap(myslice2))
  - พิมพ์ความจุของ myslice2 ซึ่งคือ 4
f.Println(myslice2)
  - พิมพ์ค่าใน myslice2 ซึ่งจะเป็น [Go Slices Are Powerful]
*/
/*สรุป
ในตัวอย่างนี้เราได้เห็น
ใน slice แรก (myslice1) ไม่มีการระบุสมาชิกจริง
ดังนั้นความยาวและความจุของ slice จะเป็น 0
ใน slice ที่สอง (myslice2) มีการระบุสมาชิกจริง ดังนั้ัน
ทั้งความยาวและความจุของ slice จะเท่ากับจำนวนสมาชิกที่ระบุ

การใช้ฟังก์ชัน len() และ cap() ช่วยเราสามารถตรวจสอบและจัดการกับขนาด
และความจุของ slice ได้อย่างมีประสิทธิภาพในภาษา Go
*/

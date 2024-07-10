package main

import (
	f "fmt"
)

// การเพิ่ม Slice หนึ่งลงใน Slice อื่น
/*เราสามารถเพิ่มสมาชิกทั้งหมดของ Slice หนึ่งลงใน Slice อื่น
โดยใช้ฟังก์ชัน append()
*/
/* Note: จุดสามจุด (...) หลัง slice2 จำเป็นเมื่อ
เราต้องการเพิ่มสมาชิกทั้งหมดของ slice2 ลงใน slice1
*/
// Syntax slice3 = append(slice1, slice2...)
func main() {
	myslice1 := []int{1, 2, 3} //ประกาศ slice ชื่อ myslice1 ที่มี 3 สมาชิก
	myslice2 := []int{4, 5, 6} // ประกาศ slice ชื่อ myslice2 ที่มี 3 สมาชิก

	myslice3 := append(myslice1, myslice2...) //เพิ่มสมาชิกของ myslice2 ลงใน myslice1 และเก็บผลลัพธ์ใน myslice3

	f.Printf("myslice3 = %v\n", myslice3)      //พิมพ์ค่าใหม่ของ myslice3
	f.Printf("ความยาว =  %d\n", len(myslice3)) //พิมพ์ความยาวใหม่ของ myslice3
	f.Printf("ความจุ = %d\n", cap(myslice3))   //พิมพ์ความจุใหม่ของ myslice3

}

/*จดตามความเข้าใจ
ประกาศ slice ชื่อ myslice1 ที่มี 3 สมาชิก
และประกาศ slice อีกรอบ ชื่อ myslice2 ที่มี 3 สมาชิก
เพิ่มสมาชิกของ myslice2 ลงใน myslice1 และเก็บผลลัพธ์ลงใน myslice3
พิมพ์ค่าใหม่ของ myslice3 ซึ่งคือ [1, 2, 3, 4, 5, 6]
พิมพ์ความยาวใหม่ของ myslice3 ซึ่งคือ 6
พิมพ์ความจุใหม่ของ myslice3 ซึ่งคือ 6
*/
/* 1.การประกาศ slice
- myslice1 := []int{1, 2, 3}
ประกาศ slice ชื่อ myslice1 ที่มีประเภทข้อมูล
เป็น int และมีค่าเริ่มต้นเป็น 1 , 2 , และ 3
- myslice2 := []int{4, 5, 6}
ประกาศ slice ชื่อ myslice2 ที่มีประเภทข้อมููลเป็น
int และมีค่าเริ่มต้นเป็น 4, 5, และ 6

	2.การเพิ่มสมาชิกของ slice หนึ่งลงใน slice อื่น
myslice3 := append(myslice1, myslice2...)
ใช้ฟังก์ชัน append() เพื่อเพิ่มสมาชิกทั้งหมดของ myslice2 ลงใน myslice1
และเก็บผลลัพธ์ใน myslice3
จุดสามจุด(...)หลัง myslice2 บอกให้ฟังก์ชัน append() เพิ่มสมาชิกทั้งหมดของ myslice2
ลงใน myslice1

	3.การพิมพ์ค่าใหม่ของ slice
f.Printf("myslice3 = %v\n", myslice3)
พิมพ์ค่าใหม่ของ myslice3 ซึ่งตอนนี้คือ [1, 2, 3, 4, 5, 6]
f.Printf("ความยาว = %d\n", len(myslice3))
พิมพ์ความยาวใหม่ของ myslice3 ซึ่งคือ 6
f.Printf("ความจุ = %d\n",cap(myslice3))
พิมพ์ความจุใหม่ของ myslice3 ซึ่งคือ 6
*/
/* สรุป
- ฟังก์ชัน append() สามารถใช้ในการเพิ่มสมาชิกของ slice หนึ่งลงใน slice
อื่นได้โดยการระบุ slice ที่ต้องการเพิ่มและใช้จุดสามจุด (...)
- ความยาวและความจุของ slice ที่เกิดขึ้นจะเป็นผลรวมของสมาชิกทั้งหมดใน slice ใหม่
- โค้ดตัวอย่างนี้ทำให้เห็นว่าการเพิ่ม slice หนึ่งลงใน slice อื่นใน Go ทำได้โดยง่ายและมีประสิทธิภาพ
*/

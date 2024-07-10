package main

import (
	f "fmt"
)

// การเพิ่มสมาชิกให้กับ slice
// เราสามารถเพิ่มสมาชิกใหม่ลงใน slice ได้โดยใช้ฟังก์ชัน append()
// Syntax slice_name = append(slice_name, element1, element2, ...)
func main() {
	//วิธีการเพิ่มสมาชิกใหม่ลงใน slice ชื่อ myslice1
	myslice1 := []int{1, 2, 3, 4, 5, 6}       //ประกาศ slice ชื่อ myslice1 ที่มี 6 สมาชิก
	f.Printf("myslice1 = %v\n", myslice1)     // พิมพ์ค่าเริ่มต้นของ myslice1
	f.Printf("ความยาว = %d\n", len(myslice1)) // พิมพ์ความยาวเริ่มต้นของ myslice1
	f.Printf("ความจุ = %d\n", cap(myslice1))  // พิมพ์ความจุเริ่มต้นของ myslice1

	myslice1 = append(myslice1, 20, 21)       // เพิ่มสมาชิกใหม่ 20 และ 21 ลงใน myslice1
	f.Printf("myslice1 = %v\n", myslice1)     // พิมพ์ค่าใหม่ของ myslice1
	f.Printf("ความยาว = %d\n", len(myslice1)) // พิมพ์ความยาวใหม่ของ myslice1
	f.Printf("ความจุ = %d\n", cap(myslice1))  // พิมพ์ความจุใหม่ของ myslice1
}

/* 1.การประกาศ slice
- myslice1 := []int{1, 2, 3, 4, 5, 6}
ประกาศ slice ชื่อ myslice1 ที่มีประเภทข้อมูลเป็น int และมีค่าเริ่มต้นเป็น
1, 2, 3, 4, 5, และ 6
   2.การพิมพ์ค่าเริ่มต้นของ slice
- f.Printf("myslice1 = %v\n", myslice1)
พิมพ์ค่าเริ่มต้นของ myslice1 ซึ่งคือ [1, 2, 3, 4, 5, 6]
- f.Printf("ความยาว = %d\n", len(myslice1))
พิมพ์ความยาวเริ่มต้นของ myslice1 ซึ่งคือ 6
- f.Printf("ความจุ = %d\n", cap(myslice1))
พิมพ์ความจุเริ่มต้นของ myslice1 ซึ่งคือ 6
   3.การเพิ่มสมาชิกใหม่ใน slice1
- myslice1 = append(myslice1, 20, 21)
เพิ่มสมาชิกใหม่ 20 และ 21 ลงใน myslice1
- f.Printf("myslice1 = %v\n", myslice1)
พิมพ์ค่าใหม่ของ myslice1 ซึ่งตอนนี้คือ [1, 2, 3, 4, 5, 6, 20, 21]
- f.Printf("ความยาว = %d\n", len(myslice1))
พิมพ์ความยาวใหม่ของ myslice1 ซึ่งคือ 8
- f.Printf("ความจุ = %d\n", cap(myslice1))
พิมพ์ความจุใหม่ของ myslice1 ซึ่งตอนนี้คือ 12
*/
/* สรุป
- ฟังก์ชัน append() ใช้ในการเพิ่มสมาชิกใหม่ลงใน slice
- เมื่อเพิ่มสมาชิกใหม่ลงใน slice ความยาวของ slice
จะเพิ่มขึ้น
- หากความจุปัจจุบันของ slice ไม่เพียงพอสำหรับสมาชิกใหม่
Go จะจัดสรรความจุใหม่ที่ใหญ่กว่าเดิมเพื่อรองรับการเติบโตของ slice
*/

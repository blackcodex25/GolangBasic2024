package main

import (
	f "fmt"
)

// การเปลี่ยนความยาวของ slice
// ไม่เหมือนกับ array เราสามารถเปลี่ยนความยาวของ slice ได้ใน Go
// วิธีการเปลี่ยนความยาวของ Slice
// Syntax index:value
func main() {
	arr1 := [6]int{9, 10, 11, 12, 13, 14}     //ประกาศ array ชื่อ arr1 ที่มี 6 สมาชิก
	myslice1 := arr1[1:5]                     // ประกาศ slice ชื่อ myslice1 ที่เริ่มต้นจาก index 1 ถึง 4 ของ arr1
	f.Printf("myslice1 = %v\n", myslice1)     //พิมพ์ค่าเริ่มต้นของ myslice1
	f.Printf("ความยาว = %d\n", len(myslice1)) //พิมพ์ความยาวเริ่มต้น myslice1
	f.Printf("ความจุ = %d\n", cap(myslice1))  //พิมพ์ความจุเริ่มต้น myslice1

	myslice1 = arr1[1:3]                      //เปลี่ยนความยาวโดยการสร้าง slice ใหม่จาก arr1
	f.Printf("myslice1 = %v\n", myslice1)     //พิมพ์ค่าเริ่มต้นของ slice1
	f.Printf("ความยาว = %d\n", len(myslice1)) // พิมพ์ความยาวใหม่ของ slice1
	f.Printf("ความจุ = %d\n", cap(myslice1))  // พิมพ์ความจุใหม่ของ slice1

	myslice1 = append(myslice1, 20, 21, 22, 23) //เปลี่ยนความยาวโดยการเพิ่มสมาชิกใหม่
	f.Printf("myslice1 = %v\n", myslice1)       //พิมพ์ค่าใหม่ของ myslice1 หลังจากเพิ่มสมาชิกใหม่
	f.Printf("ความยาว = %d\n", len(myslice1))   // พิมพ์ความยาวใหม่ของ myslice1 หลังจากเพิ่มสมาชิกใหม่
	f.Printf("ความจุ = %d\n", cap(myslice1))    // พิมพ์ความจุใหม่ของ myslice1 หลังจากเพิ่มสมาชิกใหม่
}

/* สรุป
ใน Go เราสามารถเปลี่ยนความยาวของ slice ได้โดยการ re-slice
หรือการเพิ่มสมาชิกใหม่ด้วยฟังก์ชัน append()
การ re-slice สามารถทำได้โดยการระบุช่วง index ใหม่
ฟังก์ชัน append() ใช้ในการเพิ่มสมาชิกใหม่ลงใน slice
และจะทำให้ความยาวของ slice เพิ่มขึ้นตามจำนวนสมาชิกที่เพิ่ม
*/
/* 1.การประกาศ array และ slice
- arr1 :=[6]int{9, 10, 11, 12, 13, 14}
ประกาศ array ชื่อ arr1 ที่มีประเภทข้อมูลเป็น
int และมีความเริ่มต้นเป็น 9, 10, 11, 12, 13, และ 14
- myslice1 := arr1[1:5]
ประกาศ slice ชื่อ myslice1 ที่เริ่มต้นจาก index 1 ถึง 4 ของ arr1
myslice1 จะมีค่า [10, 11, 12, 13]

   2.การพิมพ์ค่าเริ่มต้นของ slice
- f.Printf("myslice1 = %v\n", myslice1)
พิมพ์ค่าเริ่มต้นของ myslice1 ซึ่งคือ [10, 11, 12, 13]
- f.Printf("ความยาว = %d\n", len(myslice1))
พิมพ์ความยาวเริ่มต้นของ myslice1 ซึ่งคือ 4
- f.Printf("ความจุ = %d\n", cap(myslice1))
พิมพ์ความจุเริ่มต้นของ myslice1 ซึ่งคือ 5

  3.การเปลี่ยนความยาวของ slice โดยการ re-slice
- myslice1 = arr1[1:3]
เปลี่ยนความยาวของ myslice1 โดยการสร้าง slice ใหม่จาก arr1 ที่เริ่มต้นจาก
index 1 ถึง 2
- myslice1 จะมีค่าใหม่เป็น [10, 11]
- f.Printf("myslice1 = %v\n", myslice1)
พิมพ์ค่าใหม่ของ myslice1 ซึ่งคือ [10, 11]
- f.Printf("ความยาว = %d\n", len(myslice1))
พิมพ์ความยาวใหม่ของ myslice1 ซึ่งคือ 2
- f.Printf("ความจุ = %d\n", cap(myslice1))
พิมพ์ความจุใหม่ของ myslice1 ซึ่งคือ 5

   4.การเปลี่ยนความยาวของ slice โดยการเพิ่มสมาชิกใหม่
- myslice1 =append(myslice1, 20, 21, 22, 23)
เพิ่มสมาชิกใหม่ 20, 21, 22, และ 23 ลงใน myslice1
myslice1 จะมีค่าใหม่เป็น [10, 11, 20, 21, 22, 23]
- f.Printf("myslice1 = %v\n", myslice1)
พิมพ์ความใหม่ของ myslice1 หลังจากเพิ่มสมาชิกใหม่
ซึ่งคือ [10, 11, 20, 21, 22, 23]
- f.Printf("ความยาว = %d\n", len(myslice1))
พิมพ์ความยาวใหม่ของ myslice1 หลังจากเพิ่มสมาชิกใหม่ ซึ่งคือ 6
- f.Printf("ความจุ = %d\n", cap(myslice1))
พิมพ์ความจุใหม่ของ myslice1 หลังจากเพิ่มสมาชิกใหม่ ซึ่งคือ 10
*/

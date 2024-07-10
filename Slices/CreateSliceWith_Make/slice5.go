package main

import (
	f "fmt"
)

// การสร้าง Slice ด้วยฟังก์ชัน make()
// ฟังก์ชัน make() สามารถใช้ในการสร้าง Slice ได้
// Note: หากไม่ระบุพารามิเตอร์ capacity จะถือว่า capacity เท่ากับ length
func main() {
	// Syntax slice_name := make([]type, lenght, capacity)
	// กรณีที่ระบุความจุ
	myslice1 := make([]int, 5, 10)            // สร้าง slice ที่มีความยาว 5 และความจุ 10
	f.Printf("myslice1 = %v\n", myslice1)     //พิมพ์ค่าที่กำหนดเริ่มต้นใน myslice1
	f.Printf("ความยาว = %d\n", len(myslice1)) //พิมพ์ค่าความยาวของ myslice1
	f.Printf("ความจุ = %d\n", cap(myslice1))  //พิมพ์ค่าความจุของ myslice1

	// กรณีที่ไม่ระบุความจุ
	// หากไม่ระบุ ความจุ (capacity), ความจุจะเท่ากับความยาว
	myslice2 := make([]int, 5)
	f.Printf("myslice2 = %v\n", myslice2)     //พิมพ์ค่าที่กำหนดเริ่มต้นใน myslice2
	f.Printf("ความยาว = %d\n", len(myslice2)) //พิมพ์ค่าความยาวของ myslice2
	f.Printf("ความจุ = %d\n", cap(myslice2))  //พิมพ์ค่าความจุของ myslice2

	//ฟังก์ชัน Printf ใช้ในการจัดรูปแบบ (Format)
	//ตามการกำหนดรูปแบบ(formatting verb)ที่กำหนดใว้
}

/* 1.สร้าง slice ด้วยความยาวและความจุที่กำหนด
 - myslice1 := make([]int,5,10)
สร้าง slice ชื่อ myslice1 ที่มีประเภทข้อมูลเป็น int ความยาวเป็น 5 และ
ความจุเป็น 10
- f.Printf("myslice1 = %v\n",myslice1)
พิมพ์ค่าใน myslice1 ซึ่งเป็น [0 0 0 0 0 0] (ค่าเริ่มต้นของ int คือ 0)
- f.Printf("ความยาว = %d\n", len(myslice1))
พิมพ์ความยาวของ myslice1 ซึ่งคือ 5
- f.Printf("ความจุ = %d\n", cap(myslice1))
พิมพ์ความจุของ myslice1 ซึ่งคือ 10
*/
/* 2.สร้าง slice โดยไม่ระบุความจุ
- myslice2 := make([]int, 5)
สร้าง slice ชื่อ myslice2 ที่มีประเภทข้อมูลเป็น int
ความยาวและความจุเป็น 5
- f.Printf("myslice2 = %v\n", myslice2)
พิมพ์ค่าใน myslice2 ซึ่งเป็น [0 0 0 0 0]
- f.Printf("ความยาว = %d\n", len(myslice2))
พิมพ์ความยาวของ myslice2 ซึ่งคือ 5
- f.Printf("ความจุ = %d\n", cap(myslice2))
พิมพ์ความจุของ myslice2 ซึ่งคือ 5
*/
/*สรุป
-ฟังก์ชัน make() ใช้ในการสร้าง slice โดยระบุประเภท ข้อมูล,
ความยาว และความจุ
-หากไม่ระบุ capacity ความจุจะเท่ากับความยาว
-ฟังก์ชัน len() ใช้ในการคืนค่าความยาวของ slice
-ฟังก์ชัน cap() ใช้ในการคืนค่าความจุของ slice
การใช้ฟังก์ชัน make() ช่วยทำให้คุณสามารถสร้าง slice ที่มีความยืดหยุ่น
ในการจัดการความยาวและความจุได้ตามต้องการ

*/

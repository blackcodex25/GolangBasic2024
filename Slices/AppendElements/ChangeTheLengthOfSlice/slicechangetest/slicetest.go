package main

import (
	f "fmt"
)

// การเปลี่ยนความยาวของ slice
func main() {
	//ประกาศ array ชื่อ arr1 ที่มี 5 สมาชิก
	//	index		0      1    2     3      4
	arr1 := [5]int{1999, 2000, 2001, 2002, 2003} //ค่าข้อมูล 5
	// ประกาศ slice ชื่อ myslice1 ที่เริ่มต้นจาก index 1 ถึง 3 ของ arr1
	myslice1 := arr1[1:5] //index:value
	f.Printf("myslice1 = %v\n", myslice1)
	f.Printf("ความยาว = %d\n", len(myslice1))
	f.Printf("ความจุ = %d\n", cap(myslice1))

	//เปลี่ยนความยาวโดยการสร้าง Slice ใหม่จาก arr1

	myslice1 = arr1[1:4]
	f.Printf("myslice1 = %v\n", myslice1)
	f.Printf("ความยาว = %d\n", len(myslice1))
	f.Printf("ความจุ = %d\n", cap(myslice1))

	//เปลี่ยนความยาวโดยการเพิ่มสมาชิกใหม่
	myslice1 = append(myslice1, 2004, 2005, 2006, 2007, 2008)
	f.Printf("myslice1 = %v\n", myslice1)
	f.Printf("ความยาว = %d\n", len(myslice1))
	f.Printf("ความจุ = %d\n", cap(myslice1))

}

/* ประกาศ array ชื่อ arr1 ที่มี 5 สมาชิก

 */

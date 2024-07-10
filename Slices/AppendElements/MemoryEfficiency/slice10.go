package main

import (
	f "fmt"
)

//ประสิทธิภาพของการใช้หน่วยความจำ (Memory Efficiency)
/*ใน Go เมื่อใช้ slice จะโหลดสมาชิกทั้งหมดของ array ที่อยู่
เบื้องหลังลงในหน่วยความจำ ซึ่งอาจทำให้ใช้หน่วยความจำมากเกินความจำเป็น
หาก array มีขนาดใหญ่ แต่คุณต้องการเพียงไม่กี่สมาชิกเท่านั้น ในกรณีนี้
การใช้ฟังก์ชัน copy() เพื่อสร้าง array ใหม่ที่มีเฉพาะสมาชิกที่ต้องการ
จะช่วยลดการใช้หน่วยความจำ
*/
//Syntax copy(dest, src)
/*ฟังก์ชัน copy() จะรับ slice สองตัวคือ dest และ src และ
คัดลอกข้อมูลจาก src ไปยัง dest โดยจะส่งคืนจำนวนสมาชิกที่ถูกคัดลอก
*/
//วิธีการใช้ฟังก์ชัน copy()
func main() {
	//ประกาศ slice ชื่อ numbers
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	//พิมพ์ค่าเริ่มต้นของ slice numbers
	f.Printf("numbers = %v\n", numbers)
	f.Printf("ความยาว = %d\n", len(numbers))
	f.Printf("ความจุ = %d\n", cap(numbers))

	// สร้าง slice ใหม่ที่มีเฉพาะสมาชิกที่ต้องการ ลบ (-) ค่าสมาชิกที่ต้องการ -10
	neededNumbers := numbers[:len(numbers)-10]
	// ฟังก์ชัน make() ใช้สำหรับสร้าง slice โดยระบุชนิดข้อมูล ความยาว และความจุ
	numberscopy := make([]int, len(neededNumbers))
	//ฟังก์ชัน copy() จะรับ slice สองตัว คือ dest และ src และคัดลอกข้อมูลจาก src ไปยัง dest โดยจะส่งคืนจำนวนสมาชิกที่ถูกคัดลอก
	copy(numberscopy, neededNumbers)

	//พิมพ์ค่าใหม่ของ slice numbersCopy หลังจากการคัดลอก
	f.Printf("numberCopy = %v\n", numberscopy)
	f.Printf("ความยาว = %d\n", len(numberscopy))
	f.Printf("ความจุ = %d\n", cap(numberscopy))
	//ความจุของ slice ใหม่จะน้อยกว่าความจุของ slice เดิมเนื่องจาก array ที่อยู่เบื้องหลัง slice ใหม่มีขนาดเล็กกว่า
}

/* สรุป
- การใช้ slice โดยตรงจาก array อาจทำให้ใช้หน่วยความจำมากเกินความจำเป็น
- ฟังก์ชัน copy() ช่วยให้คุณสามารถสร้าง slice ใหม่ที่มีเฉพาะสมาชิกที่ต้องการ
ทำให้ลดการใช้หน่วยความจำ
- copy(dest, src) ใช้ในการคัดลอกข้อมูลจาก slice หนึ่งไปยังอีก slice หนึ่ง
และส่งคืนจำนวนสมาชิกที่ถูกคัดลอก

*/
/* 1.การประกาศและพิมพ์ค่าเริ่มต้นของ slice
- numbers := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
ประกาศ slice ชื่อ numbers ที่มีสมาชิกเป็น 1 ถึง 15
- f.Printf("numbers = %v\n", numbers)
พิมพ์ค่าเริ่มต้นของ numbers
- f.Printf("ความยาว = %d\n", len(numbers))
พิมพ์ความยาวเริ่มต้นของ numbers ซึ่งคือ 15
- f.Printf("ความจุ = %d\n", cap(numbers))
พิมพ์ความจุเริ่มต้นของ numbers ซึ่งคือ 15

   2.การสร้าง slice ใหม่ที่มีเฉพาะสมาชิกที่ต้องการ
- neededNumbers := numbers[:len(numbers)-10]
สร้าง slice ชื่อ neededNumbers ที่มีสมาชิกตั้งแต่ 1 ถึง 5
โดยตัดออก 10 สมาชิกสุดท้ายของ numbers
- numbersCopy := make([]int, len(neededNumbers))
สร้าง slice ใหม่ชื่อ numberscopy ที่มีขนดเท่ากับ needeNumbers
- copy(numberscopy, neededNumbers)
คัดลออกสมาชิกจาก neededNumbers ไปยัง numberscopy

	3.การพิมพ์ค่าใหม่ของ slice หลังการคัดลอก
- f.Printf("numbersCopy = %v\n", numbersCopy)
พิมพ์ค่าใหม่ของ numberscopy หลังจากการคัดลอก ซึ่งคือ [1,2,3,4,5]
- f.Printf("ความยาว = %d\n", len(numberscopy))
พิมพ์ความยาวใหม่ของ numberscopy ซึ่งคือ 5
- f.Printf("ความจุ = %d\n", cap(numberscopy))
พิมพ์ความจุใหม่ของ numberscopy ซึ่งคือ 5

*/

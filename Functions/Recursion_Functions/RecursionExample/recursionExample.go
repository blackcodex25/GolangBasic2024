package main

import (
	f "fmt"
)

// ประกาศฟังก์ชันชื่อ factorial_recursion รับค่าพารามิเตอร์ชื่อ Boyroblox ชนิด float64 คืนค่าผลลัพธ์ชื่อ superBoy ชนิด float64
func factorial_recursion(Boyroblox float64) (superBoy float64) {
	if Boyroblox > 0 { // ตรวจสอบว่าค่า Boyroblox มากกว่า 0 หรือไม่
		// ถ้า Boyroblox มากกว่า 0, ให้ค่า superBoy เท่ากับ Boyroblox คูณกับผลลัพธ์ของการเรียกใช้ factorial_recursion(Boyroblox-1)
		superBoy = Boyroblox * factorial_recursion(Boyroblox-1)
	} else { // ถ้า Boyroblox ไม่มากกว่า 0 (คือเท่ากับ 0)
		superBoy = 1 // กำหนดค่า superBoy เป็น 1
	}
	return // คืนค่าผลลัพธ์ superBoy
}
func main() {
	// เรียกใช้ฟังก์ชัน factorial_recursion ด้วยค่าพารามิเตอร์ 5 และพิมพ์ผลลัพธ์ออกมา
	f.Println(factorial_recursion(5))
}

/* 1. ประกาศฟังก์ชัน factorial_recursion
ฟังก์ชันนี้รับพารามิเตอร์หึ่งตัวชื่อ Boyroblox ชนิด float64
และคืนค่าผลลัพธ์ชื่อ superBoy ชนิด float64

2. ตรวจสอบเงื่อนไขในฟังก์ชัน factorial_recursion
if Boyroblox > 0 {
	superBoy = Boyroblox * factorial_recursion(Boyroblox-1)
} else {
	superBoy = 1
}
ถ้า Boyroblox มากกว่า 0
- คำนวณค่า superBoy โดยการคูณ Boyroblox
กับผลลัพธ์ของการเรียกใช้ฟังก์ชัน factorial_recursion
โดยใช้ค่าของ Boyroblox - 1
- เช่น ถ้า Boyroblox เป็น 5, การคำนวณจะเป็น 5 * factorial_recursion(4)
ถ้า Boyroblox เท่ากับ 0
- กำหนดค่า superBoy เป็น 1 (นี่เป็นเงื่อนไขหยุด การเรียกใช้ฟังก์ชัน recursion)

3. คืนค่าผลลัพธ์จากฟังก์ชัน factorial_recursion
return
- คืนค่าผลลัพธ์ superBoy

4. ประกาศฟังก์ชัน factorial_recursion
ด้วยพารามิเตอร์ 5 และพิมพ์ผลลัพธ์ออกมาโดยใช้
f.Println()
*/
/* สรุปการทำงานของโปรแกรม
โปรแกรมจะคำนวณค่าแฟคทอเรียลของ 5 โดยใช้การเรียกฟังก์ชันแบบ recursion
ผลลัพธ์ที่ได้จะเป็น 120 (5! = 5 * 4 * 3 * 2 * 1)
*/
/* วิธีการคำนวณค่าแฟคทอเรียล
1. factorial_recursion(5) = 5 * factorial_recursion(4)
2. factorial_recursion(4) = 4 * factorial_recursion(3)
3. factorial_recursion(3) = 3 * factorial_recursion(2)
4. factorial_recursion(2) = 2 * factorial_recursion(1)
5. factorial_recursion(1) = 1 * factorial_recursion(0)
6. factorial_recursion(0) = 1 (เป็นเงื่อนไขหยุด)
7. ดังนั้น, ผลลัพธ์จะเป็น 5 * 4 * 3 * 2 * 1 = 120
*/

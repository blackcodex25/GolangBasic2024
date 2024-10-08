package main

import (
	f "fmt"
)

/*ในภาษา Go เราสามารถเปลี่ยนแปลงค่าของสมาชิกใน array ได้
โดยการระบุดัชนีของสมาชิกที่ต้องการเปลี่ยนแปลงแล้ว กำหนดค่าที่ต้องการ
เปลี่ยนแปลงแล้วกำหนดค่าที่ต้องการให้กับสมาชิกนั้น
*/
/*การเปลี่ยนแปลงค่าสมาชิกใน Array
เพื่อเปลี่ยนแปลงค่าของสมาชิกใน array ให้ระบุดัชนีของสมาชิกที่
ต้องการเปลี่ยนแปลงภายในวงเล็บสี่เหลี่ยมหลังชื่อ array จากนั้น
กำหนดค่าที่ต้องการให้กับสมาชิกนั้นเช่น
array_name[index] = new_value
*/
//แสดงการเปลี่ยนแปลงค่าของสมาชิกตัวที่สามใน array
func main() {
	prices := [3]int{10, 20, 30} // ประกาศและกกำหนดค่าเริ่มต้นให้กับ array
	prices[2] = 50               // เปลี่ยนแปลงค่าของสมาชิกตัวที่สามเป็น 50
	f.Println(prices)            // พิมพ์ค่าใน array หลังจากการเปลี่ยนแปลง
}

/* สรุป
การเปลี่ยนแปลงค่าของสมาชิกใน array ในภาษา Go สามารถ
ทำได้โดยการระบุดัชนีของสมาชิกที่ต้องการเปลี่ยนปลงแล้ว
กำหนดค่าที่ต้องการให้กับสมาชิกนั้น การทำเช่นนี้ช่วยให้เรา
สามารถอัปเดตค่าที่เก็บใน array ได้ตามต้องการ ทำให้การจัดการและการ
ประมวลผลข้อมูลในโปรแกรมมีความยืดหยุ่นมากขึ้น
*/

package main

import (
	f "fmt"
)

//การเปลี่ยนค่าของสมาชิกใน Slice
/*
เราสามารถเปลี่ยนค่าของสมาชิกเฉพาะใน slice ได้โดยการอ้างถึง
หมายเลขดัชนี (index number) และกำหนดค่าใหม่ให้กับสมาชิกนั้น
*/
//วิธีการเปลี่ยนค่าสมาชิกที่สามใน slice ชื่อ prices
func main() {
	prices := []int{10, 20, 30} //ประกาศ slice ชื่อ prices ที่มี 3 สมาชิก
	prices[2] = 50              // เปลี่ยนค่าของสมาชิกที่สามให้เป็น 50
	f.Println(prices[0])        // พิมพ์สมาชิกตัวแรกของ slice
	f.Println(prices[2])        // พิมพ์สมาชิกที่สามของ slice ที่ถูกเปลี่ยนค่าแล้ว
}

// การประกาศ slice
/* prices := []int{10,20,30}
ประกาศ slice ชื่อ prices ที่มีประเภทข้อมูล เป็น int
และมีค่าเริ่มต้นเป็น 10, 20, และ 30

 การเปลี่ยนแปลงค่าของสมาชิก
prices[2] = 50
เปลี่ยนค่าสมาชิกที่สาม (ดัชนี 2) ให้เป็น 50

 การพิมพ์สมาชิกของ slice
 f.Println(prices[0])
พิมพ์สมาชิกแรกของ prices ซึ่งมีค่าเป็น 10
 f.Println(prices[2])
พิมพ์สมาชิกที่สามของ prices ซึ่งมีค่าใหม่เป็น 50
*/
/* สรุป
- เราสารมารถเปลี่ยนค่าของสมาชิกใน slice ได้
โดยการอ้างถึงหมายเลขดัชนีและกำหนดค่าใหม่ให้กับสมาชิกนั้น
- ในตัวอย่างนี้
เราได้เปลี่ยนค่าของสมาชิกที่สามจาก 30 เป็น 50
- การเข้าถึงและการเปลี่ยนค่าของสมาชิกใน slice
ทำได้ง่ายเพียงแค่ระบุหมายเลขดัชนีและค่าใหม่ที่ต้องการ
*/

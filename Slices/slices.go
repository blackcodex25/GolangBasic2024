package main

import (
	f "fmt"
)

/*
Slices มีความคล้ายคลึงกับ array แต่มีความยืดหยุ่นและทรงพลังมากกว่า
เนื่องจากขนาดของ slice สามารถขยายหรือหดได้ตามต้องการ ต่างจาก array
ที่มีขนาดคงที่
*/
/*คุณสมบัติของ Slices
- ใช้เก็บหลายค่าในตัวแปรเดียวกันที่มีชนิดข้อมูลเดียวกัน
- ขนาดของ Slices สามารถขยายหรือหดได้
- มีวิธีสร้าง slices หลายแบบ
*/
//วิธีการสร้าง Slice
/*ใช้รูปแบบ []datatype{values}
   - สามารถสร้าง slice ได้โดยตรงด้วยการระบุชนิดข้อมูลและค่าที่ต้องการ
  สร้าง Slice จาก array
   - สามารถสร้าง slice โดยการระบุช่วงของสมาชิกจาก array
   ใช้ฟังก์ชัน make()
   - ฟังก์ชัน make() ใช้สำหรับสร้าง slice โดยระบุชนิด ข้อมูล ความยาว และความจุ
*/
func main() {
	//ใช้รูปแบบ []datatype{values}
	slice1 := []int{1, 2, 3, 4, 5} //สามารถสร้าง slice ได้โดยตรงด้วยการระบุชนิดข้อมูลและค่าที่ต้องการ
	f.Println(slice1)
	//สร้าง slice จาก array
	arr := [5]int{6, 7, 8, 9, 10}
	slice2 := arr[1:4] // สร้าง slice จากสมาชิกของ array ตั้งแต่ดัชนี 1 ถึง 3
	f.Println(slice2)
	//ใช้ฟังก์ชัน make()
	slice3 := make([]int, 3, 5) // สร้าง slice ที่มีความยาว 3 และความจุ 5
	f.Println(slice3)
}

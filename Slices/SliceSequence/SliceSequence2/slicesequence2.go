package slicesequence2

import (
	"cmp"
	"fmt"
	"slices"
)

/* นำเข้าแพ็กเกจ slices
แพ็กเกจนี้มีฟังก์ชันที่ใช้งานสะดวกสำหรับการจัดการกับ slices
ฟังก์ชัน SortedFunc ใช้สำหรับเรียงลำดับ slice ตามฟังก์ชันเปรียบเทียบที่กำหนดให้ผู้ใช้

นำเข้าแพ็กเกจ cmp
แพ็กเกจนี้ใช้สำหรับการเปรียบเทียบค่าต่างๆ
ฟังก์ชัน cmp.Compare(i, j) คืนค่า -1 ถ้า i น้อยกว่า j, 0 ถ้า i == j,
และ 1 ถ้า i มากกว่า j
*/
// เขียน sort ตัวเลขจากน้อยไปมาก
func Sort() {
	// ประกาศตัวแปร slice ชื่อ numbers ที่ประกอบด้วยตัวเลข
	numbers := []int{1, 2, 5, 4, 7}
	// การเรียงลำดับตัวเลขใน slice
	// slices.SortedFunc() ใช้เพื่อเรียงลำดับตัวเลข slice
	numbersSorted := slices.SortedFunc(slices.Values(numbers), func(i, j int) int {
		// ฟังก์ชัน cmp.Compare ถูกใช้ภายในฟังก์ชันเปรียบเทียบ
		// ที่ส่งเป็นอากิวเมนต์ให้ SortedFunc
		return cmp.Compare(i, j)
		/* ฟังก์ชัน cmp.Compare เปรียบเทียบค่าของ i และ j
		ถ้า i < j จะคืนค่า -1
		ถ้า i == j จะคืนค่า 0
		ถ้า i > j จะคืนค่า 1
		*/
	})
	// ผลลัพธ์ที่ถูกเรียงลำดับแล้วจะถูกพิมพ์ออกจอ
	fmt.Println(numbersSorted)
}

/* สรุป
โค้ดนี้ใช้ slice.SortedFunc และ cmp.Compare เพื่อเรียงลำดับ slice ของตัวเอง
ฟังก์ชัน cmp.Compare ช่วยในการเปรียบเทียบค่าเพื่อใช้ในการเรียงลำดับ
ผลลัพธ์ของการเรียงลำดับจะถูกพิมพ์ออกจอ
*/

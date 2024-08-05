package main

import (
	f "fmt"
	"strconv"
)

/*
1.การประกาศชนิด error ใหม่
การประกาศชนิด error ใหม่จะเกิดความขัดแย้งกับชนิด error
ที่มีอยู่แล้วใน Go ซึ่งอินเตอร์เฟซมาตรฐานสำหรับการ
จัดการข้อผิดพลาด ตัวอย่างบรรทัดที่ 14
*/
type error interface {
	Error() string
}

/*
	วิธีแก้คือ ลบการประกาศชนิด error

ใหม่เพื่อลดความขัดแย้งกับชนิด error ที่มีอยู่แล้วใน Go
ลบบรรทัดที่ 14 - 16 ออกข้อผิดพลาดโค้ดนี้จะหายไป
*/
func main() {
	i, err := strconv.Atoi("42")
	if err != nil {
		f.Printf("ไม่สามารถแปลงตัวเลขได้: %v\n", err)
		return
	}
	f.Println("จำนวนเต็มที่ถูกแปลง: ", i)
}

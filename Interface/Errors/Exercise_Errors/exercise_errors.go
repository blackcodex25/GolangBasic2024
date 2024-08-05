package main

import (
	"fmt"
	"math/cmplx"
)

// ประกาศประเภทชื่อ ErrNegativeSqrt ชนิด float64
type ErrNegativeSqrt float64

/*
	ประกาศเมธอด Error() และคืนค่าชนิด string

สำหรับ pointer ตัวรับค่า e (e ErrNegativeSqrt)
*/
func (e ErrNegativeSqrt) Error() string {
	// แปลงค่า ErrNegativeSqrt เป็น float64 ผ่านตัวรับ e
	// ประกาศฟังก์ชัน return คืนค่าสตริงและค่าของตัวรับ e
	// โดยใช้ฟังก์ชัน Sprintf()

	return fmt.Sprintf("ไม่สามารถ Sqrt จำนวน %v ได้", float64(e))
}

// ประกาศฟังก์ชัน Sqrt() กำหนดพารามิเตอร์ x ชนิด float64
// ประกาศคืนค่าสองชนิด float64 และ error
func Sqrt(x float64) (float64, error) {
	// คืนค่าจำนวนจริง float64 สำหรับค่าที่ยังไม่ได้คำนวณ 0 และ nil คืนค่า error
	return 0, nil
}
func main() {
	// ประกาศฟังก์ชันเชิงซ้อน cmplx.Sqrt() กำหนดค่าคำนวณ 2 พิมพ์ผลลัพธ์ออกจอ
	fmt.Println(cmplx.Sqrt(2))
	/* ประกาศฟังก์ชัน ErrNegativeSqrt กำหนดค่า -2
	เรียกใช้เมธอด Error() ของ ErrNegativeSqrt
	ซึ่งคืนค่าสตริงที่กำหนดถึงข้อผิดพลาด */
	fmt.Println(ErrNegativeSqrt(-2).Error())

}

/* สรุป
โค้ดนี้มีวัตถุประสงค์เพื่อแสดงการจัดการข้อผิดพลาดเมื่อ
พยายามคำนวนณรากที่สองของตัวเลขลบในภาษา Go โดยใช้
ชนิดข้อมูล ErrNegativeSqrt และเมธอด Error()
เพื่อคืนค่าข้อผิดพลาดในรูปแบบสตริง นอกจากนี้ยังใช้ฟังก์ชัน
cmplx.Sqrt() จากแพ็กเกจ math/cmplx
เพื่อคำนวณรากที่สองของตัวเลขเชิงซ้อนและแสดงผลลัพธ์ออกจอ

*/

package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

/* นำเข้าแพ็กเกจ image เพื่อใช้สำหรับการทำงานกับภาพ
นำเข้าแพ็กเกจ image/color เพื่อใช้สำหรับการจัดการกับสี
นำเข้าแพ็กเกจ golang.org/x/tour/pic เพื่อใช้ฟังก์ชัน
pic.ShowImage สำหรับแสดงภาพ
*/
/* ประกาศโครงสร้าง image ที่มีฟิลด์ w (ความกว้าง)
และ h (ความสูง)
*/
type Image struct {
	w, h int
}

/*การประกาศเมธอดสำหรับ image
เมธอด Bounds
*/
/* เมธอด Bounds คืนค่า ขอบเขตของภาพ
เป็น image.Rectangle จาก (0, 0) ถึง (w, h)
*/
func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.w, img.h)
}

/*
	เมธอด ColorModel คืนค่า รูปแบบสีของภาพเป็น

color.RGBAModel
*/
func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

/*
เมธอด At คืนค่าสีของพิกเซลที่ตำแหน่ง (x, y) เป็น
color.Color
ค่าสีที่ใช้เป็นสี RGBA โดยที่ค่าแดงและเขียวถูกกำหนดเป็น v
(ซึ่งคำนวณจาก (x + y) / 2) ค่าน้ำเงินเป็น 255 และค่า
ความทึบแสง (alpha) เป็น 255
*/
func (img Image) At(x, y int) color.Color {
	v := uint8((x + y) / 2)
	return color.RGBA{v, v, 255, 255}
}
func main() {
	/* สร้างภาพ Image ที่มีขนาด 256x256 พิกเซล
	ใช้ฟังก์ชัน pic.ShowImage เพื่อแสดงภาพ
	*/
	m := Image{w: 256, h: 256}
	pic.ShowImage(m)
}

/* การทำงานของโปรแกรม
1.การสร้างภาพ
m := Image{w: 256, h: 256}
สร้างภาพชนิด Image ที่มีความกว่างและความสูง 256 พิิกเซล

2.การแสดงภาพ
pic.ShowImage(m)
ใช้ฟังก์ชัน pic.ShowImage เพื่อแสดงภาพที่สร้างขึ้น
เมธอด At และการคำนวณสี
เมธอด At คืนค่าสีของพิกเซลที่ตำแหน่ง (x, y)
โดยคำนวณค่า v จาก (x + y) / 2 และใช้ค่า v สำหรับค่า
แดงและเขียว ค่าน้ำเงินและความทึบแสงถูกกำหนดเป็น 255
*/

/* ผลลัพธ์ที่ได้
เมื่อรันโปรแกรมจะได้ภาพขนาด 256x256 พิกเซล โดยแต่ละ
พิกเซลมีสีที่คำนวณจากตำแหน่ง (x, y) ของพิกเซล */

/* สรุป
โปรแกรมนี้สร้างโครงสร้าง Image ที่ทำงานตามอินเตอร์เฟซ
image.Image ในภาษา Go โดยการกำหนดเมธอด Bounds
ColorModel และ At และใช้ pic.ShowImage เพื่อแสดง
ภาพที่สร้างขึ้น เมธอด At คืนค่าสีของพิกเซลที่ตำแหน่ง
(x, y) ตามการคำนวณ (x + y) / 2 โดยค่าสีแดงและเขียว
ถูกกำหนดเป็น v ค่าน้ำเงินและความทึบแสงเป็น 255
*/

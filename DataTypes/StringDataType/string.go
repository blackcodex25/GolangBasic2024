package main

import (
	f "fmt"
)

// ประเภทข้อมูล String ใน Go
// ประเภทข้อมูล string ใช้ในการเก็บชุดของตัวอักษร (ข้อความ)
// โดยค่าของ string ต้องอยู่ในเครื่องหมายคำพูดคู่ (")

var txt1 string = "Hello!" //ประกาศเป็น string และกำหนดค่าเป็น "Hello!"
var txt2 string            //txt2 ถูกประกาศเป็น string แต่ไม่ได้กำหนดค่าเริ่มต้น ดังนั้นค่าของ txt2 จะเป็นค่าเริ่มต้นของประเภท string ซึ่งคือ "" (ข้อความว่าง)

func main() {
	txt3 := "Dev" //txt3 ถูกประกาศและกำหนดค่าเป็น "World 1" โดยใช้การประกาศแบบย่อ

	//การใช้ fmt.Printf เพื่อแสดงชนิด (%T) และค่า (%v) ของตัวแปรแต่ละตัวทำให้เราเห็นข้อมูลที่ถูกเก็บในตัวแปรเหล่านั้น
	f.Printf("Type: %T, value: %v\n", txt1, txt1)
	f.Printf("Type: %T, value: %v\n", txt2, txt2)
	f.Printf("Type: %T, value: %v\n", txt3, txt3)
}

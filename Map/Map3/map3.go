package main

import (
	f "fmt"
)

/*
	การสร้าง maps ในภาษา Go โดยใช้ฟังก์ชัน make()

สามารถทำได้อย่างมีประสิทธิภาพ ฟังก์ชัน make()
ช่วยให้เราสามารถสร้าง maps ว่างและกำหนด key ค่าให้กับ
maps นั้นได้ในภายหลัง

กาสร้าง Maps ด้วยฟังก์ชัน make()
การสร้าง maps ในภาษา Go สามารถทำได้โดยใช้ฟังก์ชัน make()
เพื่อเตรียมพื้นที่หน่วยความจำสำหรับเก็บข้อมูลใน map
*/

// การใช้ฟังก์ชัน make() เพื่อสร้าง maps

/* ไวยากรณ์การใช้ make() สร้าง Maps

var a = make(map[KeyType]ValueType)

b := make(map[KeyType]ValueType)
*/
// วิธีการสร้าง Maps ใน Go โดยใช้ฟังก์ชัน make()
// สร้าง map ชื่อ a โดยใช้ make() และกำหนดชนิดข้อมูลของ key และ value เป็น string:string
var a = make(map[string]string) // ขณะนี้ map ยังว่างเปล่า
func main() {
	// กำหนดค่าให้กับ map a
	a["brand"] = "Ford"
	a["model"] = "Mustang"
	a["year"] = "1964"
	// เราเพิ่มค่าลงใน map a โดยกำหนด key เป็น "brand", "model", และ "year" และกำหนดค่า value ตามลำดับ

	// ประกาศฟังก์ชัน map ชื่อ b โดยใช้ make() และกำหนดชนิดข้อมูลของ key และ value เป็น string:int
	b := make(map[string]int)
	//ประกาศฟังก์ชัน make() เพื่อสร้าง map ที่ชื่อว่า b โดยมีชนิดของ key เป็น string และชนิดของ value เป็น int

	// กำหนดค่าให้กับ map b
	b["Oslo"] = 1
	b["Bergen"] = 2
	b["Trondheim"] = 3
	b["Stavanger"] = 4
	// เราเพิ่มค่าลงใน map b โดยกำหนด key เป็นชื่อเมืองและกำหนดค่า value เป็นตัวเลขตามลำดับ

	// แสดงผลลัพธ์ของ map a
	f.Printf("a\t%v\n", a)
	// แสดงผลลัพธ์ของ map b
	f.Printf("b\t%v\n", b)

	// ประกาศฟังก์ชัน Println() เพื่อแสดงผลลัพธ์ของ map a และ b โดยใช้ %v เพื่อแสดงค่าของ map ทั้งหมด
}

/* Syntax Maps
maps[Keytype] ValueType {key1: value1, key2: value2, ...}
*/

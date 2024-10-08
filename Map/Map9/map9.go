package main

import (
	f "fmt"
)

// การวนลูปผ่าน map (lterate Over Maps) ใน ภาษา Go
/* เราสามารถใช้ range เพื่อวนลูปผ่าน map ในภาษา Go
การใช้ range กับ map ช่วยให้
เราสามารถเข้าถึงทั้งคีย์และค่าในแต่ละคู่ข้อมูลได้
*/
//วิธีการวนลูป map และพิมพ์คีย์และค่าของแต่ละคู่ข้อมูล
func main() {
	//ประกาศฟังก์ชัน map ชื่อ a กำหนดชนิดข้อมูล string และค่าชนิดข้อมูล int
	//เพิ่มคีย์และค่าเข้าไปใน map ซึ่งคือ "one": 1, "two": 2, "three": 3, "four": 4
	a := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}
	//เงื่อนไขวนลูป เรียกใช้งาน range เพื่อวนซ้ำผ่าน map ชื่อตัวแปร a
	for k, v := range a { //ประกาศตัวแปร k พร้อมได้รับคีย์ และ ประกาศตัวแปร v จะได้รับค่าที่เกี่ยวข้องกับคีย์นั้น
		f.Printf("%v : %v, ", k, v) //ประกาศฟังก์ชัน Printf() พิมพ์ค่าข้อมูลเริ่มต้นของตัวแปร k และตัวแปร v
	} // ในแต่ละรอบของการวนลูป, คีย์และค่าจะถูกพิมพ์ออกมาในรูปแบบ คีย์ : ค่า
}

/* Note:
ลำดับขององค์ประกอบใน map: คำสั่งใน map ไม่มีการจัดเรียง ดังนั้นลำดับขององค์
ประกอบที่แสดงผลอาจแตกต่างกันไปในการเรียกใช้งานแต่ละครั้ง
การวนลูป: การใช้ range จะทำให้การวนลูปสามารถเข้าถึงทุกคู่ key:value ใน map
*/
/*ผลลัพธ์
การพิมพ์คีย์และค่าของแต่ละคู่ข้อมูลใน map โดยลำดับขององค์ประกอบ
ที่แสดงผลอาจแตกต่างกันไป
	four : 4, one : 1, three : 3, two : 2,
การใช้ range กับ map ทำให้การเข้าถึงและจัดการข้อมูลใน map เป็นเรื่องง่ายและสะดวก
*/

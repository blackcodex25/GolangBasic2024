package main

/*
ฟังก์ชัน Printf() ใช้สำหรับการจัดรูปแบบ (format) อาร์กิวเมนต์
ตามการกำหนดรูปแบบ (formatting verb) ที่กำหนดไว้ แล้ว
จึงพิมพ์อาร์กิวเมนต์ออกมา
*/
/*
Verb การจัดรูปแบบที่ใช้บ่อย
%v: ใช้สำหรับพิมพ์ค่าของอาร์กิวเมนต์
%T: ใช้สำหรับพิมพ์ชนิดของอาร์กิวเมนต์
*/
import (
	f "fmt"
)

var i string = "Dev"
var j int = 15

func main() {
	f.Printf("i has value: %v and type: %T\n", i, i)
	f.Printf("i has Value: %v and type: %T", j, j)
	//ฟังก์ชัน Printf() ใช้เพื่อพิมพ์ค่าของ i และ j พร้อมกับชนิดของตัวแปรแต่ละตัว
	//  %v ใช้สำหรับพิมพ์ค่าของอาร์กิวเมนต์ และ %T ใช้สำหรับพิมพ์ชนิดของอาร์กิวเมนต์
	//  \n ใช้เพื่อขึ้นบรรทัดใหม่หลังจากพิมพ์ข้อความของ i
}

//ผลลัพธ์
// i has value: Dev and type: string
// i has Value: 15 and type: int

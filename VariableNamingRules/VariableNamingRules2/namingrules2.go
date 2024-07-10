package main

import (
	f "fmt"
)

/*
การใช้ Camel Case ช่วยให้ชื่อตัวแปรมีความอ่านง่ายขึ้น
เนื่องจากชื่อที่ใช้ตั้งแต่ต้นเป็นตัวพิมพ์เล็กและแต่ละคำ
ถัดมาเริ่มด้วยตัวพิมพ์ใหญ่
*/
//Camel Case ในรูปแบบ Camel Case,
//คำแต่ละคำที่ไม่ใช่คำแรกจะขึ้นต้นด้วยตัวอักษรใหญ่:
var myVariableName string

func main() {
	myVariableName = "John" //myVariableName เป็นตัวแปรที่ใช้รับค่าชื่อ "John"
	f.Println(myVariableName)
}

package main //การประกาศแพ็กเกจ (Package declaration)

//การนำเข้าแพ็กเกจ (Import packages)
//นำเข้าไฟล์ที่รวมอยู่ในแพ็กเกจ fmt
import (
	f "fmt"
)

// ในภาษา Go
/*:= การประกาศตัวแปร + การกำหนดค่าไปในตัว
  = มีไว้สำหรับกำหนดค่าเท่านั้น */
// การกดปุ่ม Enter จะเพิ่มเครื่องหมาย ";"
// ไว้ที่ท้ายบรรทัดโดยปริยาย (ซึ่งจะไม่ปรากฏในซอร์สโค้ด)
func main() {
	f.Println("Hello World!")
	f.Printf("%v", nil) // nil ค่าที่ยังไม่ได้กำหนดให้ตัวแปรอื่นหรือยังไม่เปลี่ยนแปลงเป็น type อื่น มันจะไม่มีtype

	// *การนำเข้าแพ็กเกจจากข้างนอก กรณีนี้คือโฟล์เดอร์ interface4
	/*b := &interface4.Bar{}
	interface4.Fizzy(b) */
}

//เครื่องหมายปีกกาซ้าย { ไม่สามารถอยู่ที่จุดเริ่มต้นของบรรทัดได้

/* int ค่าเริ่มต้นเก็บค่าได้สูงสุด 15000000000000000000
   int32 กำหนดค่าเป็น 32 bit เก็บค่าได้สูงสุด 1500000000
   int64 กำหนดค่าเป็น 64 bit เก็บค่าได้สูงสุด 1500000000000000000
*/

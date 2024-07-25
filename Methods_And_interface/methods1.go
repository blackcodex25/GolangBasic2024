package main

import (
	f "fmt"
	"math"
)

/* การใช้งาน method ในภาษา Go ซึ่งไม่มี classes แต่สามารถกำหนด method
บน type ได้โดย method เป็นฟังก์ชันที่มี receiver พิเศษที่ปรากฏในรายการ argument
ของฟังก์ชัน ระหว่าง keyword
func และชื่อ method เป็นอากิวเม้นต์ของตัวมันเอง
ในตัวอย่างนี้ method ชื่อตัวแปร abs
มี receiver ที่เป็น type Vertex ชื่อ v
รายละเอียดโค้ด
*/
/*นำเข้าแพ็กเกจ math เพื่อใช้ฟังก์ชันทางคณิตศาสตร์
 */
// ประกาศ Struct
type Vertex struct {
	X, Y float64
} // กำหนด struct ชื่อ vertex
// ซึ่งมีฟิลด์สองฟิลด์คือ x และ y ที่มีชนิดข้อมูลเป็น float64

// กำหนด method
// กำหนด method ชื่อ abs บน type Vertex

/*
v vertex เป็น receiver ของ method ชื่อ abs ซึ่งทำหน้าที่คล้ายกับ this
ในภาษาโปรแกรมเชิงวัตถุอื่นๆ
method ชื่อตัวแปร abs คืนค่าผลลัพธ์ชนิด float64
*/
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
	/*ภายใน method คำนวณค่าความยาวของเวกเตอร์โดยใช้สูตรคำนวณความยาว
	ของเวกเตอร์โดยใช้สูตรคำนวณความยาวของเส้นทะแยงมุมในระบบพิกัดคาร์ทีเซียน
	√(X^2 + Y^2) และคืนค่าผลลัพธ์ที่คำนวณได้
	*/
}
func main() {
	v := Vertex{3, 4}  // ประกาศตัวแปร v เป็น instance ของ Vertex โดยกำหนดค่า X เป็น 3 และ Y เป็น 4
	f.Println(v.Abs()) // เรียกใช้งาน method ชื่อตัวแปร abs บนตัวแปร v และพิมพ์ผลลัพธ์ที่ได้จาก method ชื่อตัวแปร abs
}

/*เมื่อรันโปรแกรม
จะได้ผลลัพธ์เป็น 5 เนื่องจาก √(3^2 + 4^2) = √(9 + 16) = √25 = 5
โค้ดแสดงให้เห็นวิธีการใช้ method ใน Go เพิ่มเพิ่มฟังก์ชันการทำงานให้กับ types
โดยไม่ต้องใช้ classes ตามแบบภาษาโปรแกรมเชิงวัตถุ
*/
/* สรุปหลักการทำงาน
1.นำเข้าแพ็กเกจ fmt และ math เพื่อใช้ฟังก์ชันพิมพ์ข้อความและคำนวณทางคณิตศาสตร์
2.ประกาศ struct ชื่อตัวแปร Vertex ที่มีสองฟิลด์ X และ Y สำหรับเก็บค่าพิกัดคาร์ทีเซียน
3.กำหนด method ชื่อตัวแปร abs บน type ชื่อตัวแปร Vertex เพื่อคำนวณความยาวของ
เวกเตอร์จากพิกัด X และ Y
4.ในฟังก์ชัน main สร้าง instance ของ Vertex
และเรียกใช้ method ชื่อตัวแปร abs เพื่อคำนวณและพิมพ์ค่าความยาวของเวกเตอร์
*/

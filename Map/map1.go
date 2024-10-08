package main

import (
	f "fmt"
)

/* การใช้งาน Maps ในภาษา Go
Maps ใช้เพื่อเก็บค่าข้อมูลในรูปแบบการจับคู่ key:value โดยที่
- แต่ละองค์ประกอบใน map ถูกจัดเก็บในรูปแบบการจับคู่ key กับ value
Map เป็นคอลเลกชันที่ไม่เรียงลำดับและเปลี่ยนแปลงได้
- ข้อมูลใน map จะไม่มีลำดับที่แน่นอนและสามารถเปลี่ยนแปลงได้ตลอดเวลา
Map ไม่อนุญาติให้มี key ซ้ำกัน
- ค่าของ key ใน map หนึ่งๆ จะต้องไม่ซ้ำกัน
ความยาวของ map
- ความยาวของ map คือ จำนวนคู่ของ key ที่มีอยู่ใน map นั้นๆ ซึ่งสามารถตรวจสอบได้ด้วยฟังก์ชัน
len()
ค่าเริ่มต้นของ map คือ nill หาก map ยังไม่ได้รับการ กำหนดค่าใดๆ ค่าเริ่มต้นจะเป็น nill
Maps อ้างอิงถึง hash table ที่อยู่เบื้องหลัง การจัดเก็บข้อมูลใน map ใช้โครงสร้าง hash table
เพื่อให้การเข้าถึงข้อมูลมีประสิทธิภาพ
*/
/* การสร้าง Maps
ภาษา Go มีวิธีการหลายแบบในการสร้าง Maps โดยสามารถสร้างได้โดยใช้ var และ :=
*/
/*การสร้าง Maps โดยใช้ var และ :=

การใช้ var
var a = map [KeyType] ValueType {key1: value1, key2: value2, ...}

การใช้ :=
ิb := map [KeyType] ValueType {key1: value1, key2: value2, ...}

*/
var a = map[int]string{1: "dafuq"}
var d = map[uint]float32{90: 98.5}

func main() {
	b := map[float64]int{7.95: 1}
	c := map[string]bool{"จริืงหรือเท็จ": true}
	f.Println(a, b, d, c)
}

/*์Note
ลำดับของข้อมูลใน map ที่กำหนดในโค้ดอาจจะไม่ตรงกับกับลำดับของข้อมูลที่จะแสดงในผลลัพธ์เนื่องจาก maps ใน Go
ถูกออกแบบมาเพื่อจัดเก็บข้อมูลในลำดับที่ทำให้การดึงข้อมูลมีประสิทธิภาพสูงสุด
*/

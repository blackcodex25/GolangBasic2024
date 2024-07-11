package main
import(f "fmt")
// รูปแบบที่ไม่ถูกต้อง if else
// *หากเราใช้คำสั่ง else ในบรรทัดใหม่ จะเกิดข้อผิดพลาด
func main() {
	temperature := 14
	if temperature > 15{
		f.Println("ข้างนอกมันอุ่น")
	} // ใช้คำสั่ง else ในบรรทัดใหม่ จะทำให้เกิด error 
	else {
		f.Println("ข้างในมันเย็น")
	}
}
/*สรุป
ในภาษา Go 
*/
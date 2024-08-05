package main

import (
	"io"
	"os"
	"strings"
)

/* รูปแบบที่พบบ่อยคือ io.Reader ที่ครอบ io.Reader อื่น
และปรับเปลี่ยนสตรีมในบางลักษณะ
ตัวอย่างเช่น ฟังก์ชัน gzip.NewReader รับ
io.Reader (สตรีมของข้อมูลที่ถูกบีบอัด)
และคืนค่า *gzip.Reader ที่ยังคงเป็น
io.Reader (สตรีมของข้อมูลที่ถูกคลายการบีบอัด)
*/
/* สร้าง rot13Reader ที่ใช้งาน io.Reader
และอ่านจาก io.Reader โดยปรับเปลี่ยนสตรีม
โดยใช้การเข้ารหัส rot13 กับตัวอักษรทั้งหมด
ประเภท rot13Reader ได้ถูกจัดเตรียมไว้ให้เราแล้ว ทำให้มัน
เป็น io.Reader โดยการใช้งานเมธอด Read ของมัน
*/
// ประกาศ struct ชื่อ rot13Reader กำหนดฟิลด์ชื่อ r ชนิด io.Reader
type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)
	for i := 0; i < n; i++ {
		rot13 := b[i]
		switch {
		case rot13 >= 'A' && rot13 <= 'Z':
			b[i] = 'A' + (rot13-'A'+13)%26
		case rot13 >= 'a' && rot13 <= 'z':
			b[i] = 'a' + (rot13-'a'+13)%26
		}
	}
	return n, err
}
func main() {
	s := strings.NewReader("LBH penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)

}

/*
buf := make([]byte, 8)
	if _, err := io.CopyBuffer(os.Stdout, s, buf); err != nil{
		log.Fatal(err)
	}
*/

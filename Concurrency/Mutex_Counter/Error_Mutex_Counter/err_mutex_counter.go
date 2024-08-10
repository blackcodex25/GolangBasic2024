package main

import (
	"errors"
	f "fmt"
	"sync"
	"time"
) /* ตัวอย่างการใช้ mutex ในการจัดการกับการเข้าถึงทรัพยากรที่ใช้
ร่วมกันระหว่างหลายๆ goroutine โดยเฉพาะเมื่อเรามีหลาย goroutine
ที่พยายามจะเข้าถึงหรือแก้ไขตัวแปรเดียวกัน หากไม่มีการจัดการที่เหมาะสม
อาจเกิดปัญหาเกี่ยวกับข้อมูลหรือ "data race" ได้ ดังนั้น
เราใช้ mutex เพื่อให้แน่ใจว่ามีเพียง goroutine เดียวเท่านั้นที่
สามารถเข้าถึงทรัพยากรในช่วงเวลาเดียวกันได้
*/
/*
	โครงสร้าง SafeCounter

เป็นโครงสร้างที่ปลอดภัยในการใช้งานพร้อมกัน (concurrent safe)
mu คือ Mutex จากแพ็กเกจ sync ซึ่งใช้เพื่อทำให้การเข้าถึง
ทรัพยากรเกิดขึ้นแบบ mutual exclusion (ไม่ให้มีการเข้าถึง
พร้อมกันได้)
v เป็น map ชนิด string:int ที่ใช้เก็บค่าจากตัวแปร counter
โดยใช้คีย์ที่เป็นสตริง
*/
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// ประกาศฟังก์ชัน Inc ใช้เพิ่มค่าของตัวนับ ( counter ) สำหรับ
// คีย์ที่กำหนด(key)
func (c *SafeCounter) Inc(key string) error {
	// c.mu.Lock จะทำการล็อก mutex ซึ่งจะทำให้มีเพียง gotoutine
	// เดียวเท่านั้นที่สามารถเข้าถึงโค้ดบล็อกนี้ในเวลาเดียวกัน
	c.mu.Lock()
	// จะปลดล็อก mutex ทำให้ goroutine อื่นสามารถเข้าถึงได้
	defer c.mu.Unlock()
	// ตรวจสอบว่า key ได้เพิ่มค่าไปแล้วจริงหรือไม่
	if _, exists := c.v[key]; !exists {
		return errors.New(" Error: ไม่มีค่า key ใน map ")
	}
	// เพิ่มค่า map ชื่อตัวแปร v เสร็จสิ้นแล้ว
	c.v[key]++

	return nil
} // ล็อกเพื่อให้แน่ใจว่าในแต่ละครั้งมีเพียง goroutine เดียวเท่านั้น
// ที่สามารถเข้าถึง map c.v ได้
// ฟังก์ชัน Value
// ฟังก์ชัน Value ใช้คืนค่าปัจจุบันของตัวนับสำหรับคีย์ที่กำหนด
// ฟังก์ชันนี้จะล็อก mutex โดยใช้ c.mu.Lock() ก่อนจะเข้าถึง map
func (c *SafeCounter) Value(key string) (int, error) {
	// ล็อกเพื่อให้แน่ใจว่าในแต่ละครั้งมีเพียง goroutine เดียวเท่านั้น
	// ที่สามารถเข้าถึง map c.v ได้
	c.mu.Lock()

	// ใช้ defer c.mu.Unlock() เพื่อให้แน่ใจว่า mutex
	// จะถูกปลดล็อกไม่ว่าจะเกิดอะไรขึ้นในโค้ดบล็อกนี้ defer
	// จะทำให้การปลดล็อกเกิดขึ้นทันทีเมื่อฟังก์ชัน value
	// สิ้นสุดการทำงาน
	defer c.mu.Unlock()
	// ตรวจสอบว่า key มีค่าอยู่ใน map หรือไม่
	val, exists := c.v[key]
	if !exists {
		return 0, errors.New("Error: ไม่มีค่า key ใน map")
	}
	return val, nil

}
func main() {
	// สร้างตัวแปร c จากโครงสร้าง SafeCounter และทำการ
	// สร้าง map สำหรับเก็บค่าตัวนับ
	c := SafeCounter{v: make(map[string]int)}
	c.v["somekey"] = 0
	/* ใช้ลูปเพื่อเริ่ม 1000 Goroutine โดยแต่ละ goroutine
	จะเรียกใช้ฟังก์ชัน Inc เพื่อเพิ่มค่าตัวนับสำหรับคีย์ somkey
	*/
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := c.Inc("somekey"); err != nil {
				f.Println("เกิดข้อผิดพลาดในการเพิ่ม: ", err)
			}
		}()
	}
	//go c.Inc("Somekey")
	/* เนื่องจาก goroutine ทำงานแบบ asynchronous
	(ไม่รอให้ทำงานเสร็จก่อนที่จะไปทำงานต่อไป) จึงใช้
	time.Sleep(time.Second) เพื่อรอให้ goroutine ทั้งหมด
	ทำงานเสร็จก่อนที่จะพิมพ์ค่าของตัวนับ
	*/
	time.Sleep(time.Second)

	/* ฟังก์ชัน Println(c.Value("somekey")) จะพิมพ์ค่าตัว
	นับปัจจุบันสำหรับคีย์ somekey
	*/
	if val, err := c.Value("somekey"); err != nil {
		f.Println("เกิดข้อผิดพลาดในการเรียกค่า: ", err)
	} else {
		f.Println("ค่าที่ควรจะเป็น: ", val)
	}

	//f.Println(c.Value("somekey"))

}

/* สรุป
Mutex: ใช้เพื่อให้แน่ใจว่ามีเพียง goroutine เดียวที่สามารถเข้า
ถึงทรัพยากรที่ใช้ร่วมกันได้ในเวลาเดียวกัน
Lock & Unlock: ใช้ในการล็อกและปลดล็อก mutex เพื่อ
ป้องกันปัญหาการเข้าถึงพร้อมกันจากหลายๆ goroutine
defer: ใช้เพื่อให้แน่ใจว่า Unlock จะถูกเรียกใช้เมื่อการทำงานของ
ฟังก์ชันเสร็จสิ้น ซึ่งช่วย้องกันปัญหาการลืมปลดล็อก mutex

โดยรวม โค้ดนี้แสดงให้เห็นถึงการใช้ sync.Mutex ใน Go เพื่อ
จัดการกับการเข้าถึงข้อมูลพร้อมกันจากหลายๆ goroutine อย่าง
ปลอดภัย
*/

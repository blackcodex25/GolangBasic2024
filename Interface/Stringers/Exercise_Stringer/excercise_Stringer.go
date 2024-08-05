package main

import (
	f "fmt"
	"strings"
)

/* โค้ดนี้ประกาศและใช้งาน struct แบบกำหนดเองสำหรับการ
แสดงที่อยู่ IP (IP address) โดยการกำหนดเมธอ string
เพื่อกำหนดรูปแบบการแสดงผลของ struct IPAddr และใช้
map เพื่อเก็บข้อมูลที่อยู่ IP ของ loopback และ googleDNS
*/
/* นำเข้าแพ็กเกจ strings
เพื่อใช้ฟังก์ชัน strings.Join() สำหรับการรวมสตริง */

/*
	ประกาศ struct ชื่อ IPAddr

ที่ประกอบด้วยอาร์เรย์ขนาด 4 ของชนิด byte
เพื่อแทนที่อยู่ IP
*/
type IPAddr [4]byte

/*
เมธอด String ถูกประกาศสำหรับ struct ชื่อ IPAddr เพื่อ
กำหนดรูปแบบการแสดงผล
*/
func (ip IPAddr) String() string {
	// สร้างสตริง slice ชื่อ parts ที่มีความยาวเท่ากับ ip
	parts := make([]string, len(ip))
	/* ใช้ลูป for เพื่อวนผ่านแต่ละไบต์ใน ip
	และแปลงเป็นสตริงด้วยฟังก์ชัน Sprintf("%d", b) */
	for i, b := range ip {
		parts[i] = f.Sprintf("%d", b)
	}
	/* รวมสตริงใน parts ด้วยจุด (.)
	โดยใช้ strings.Join(parts, ".") */
	return strings.Join(parts, ".")
}
func main() {
	/* ประกาศ map ชื่อตัวแปร hosts ที่มีคีย์เป็น string และค่าคือ IPAddr
	loopback ถูกกำหนดค่าเป็น {127, 0, 0, 1}
	googleDNS ถูกกำหนดค่าเป็น {8, 8, 8, 8}
	*/
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	/* ใช้ลูป for เพื่อวนผ่านแต่ละคู่ของคีย์และค่าใน map ชื่อตัวแปร hosts
	พิมพ์คีย์ (name) และค่า (ip) โดยใช้ Printf() ซึ่งจะเรียกเมธอด String ของ
	IPAddr เพื่อแสดงผลในรูปแบบที่กำหนด
	*/
	for name, ip := range hosts {
		f.Printf("%v: %v\n", name, ip)
	}
} // map คำสั่งไวยากรณ์ key:value
/* สรุป
โปรแกรมนี้ประกาศ struct ชื่อตัวแปร IPAddr เพื่อแทนที่อยู่ IP และเพิ่มเมธอด string
เพื่อกำหนดรูปแบบการแสดงผลที่อยู่ IP ในรูปแบบของสตริงที่คั่นด้วยจุด (.) ในฟังก์ชันหลัก
โปรแกรมสร้าง map เพื่อเก็บที่อยู่ IP ของ loopback และ googleDNS
และใช้ลูปเพื่อพิมพ์ที่อยู่ IP แต่ละรายการออกทางจอภาพ
*/

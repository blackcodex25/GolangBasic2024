package main

import "fmt"

// Struct 1 ประกาศฟังก์ชัน struct ชื่อ telephone
// กำหนดสมาชิก name, spec, price และ warranty
type telephone struct {
	name     string
	spec     string
	price    int
	warranty int
}

// Struct 2 ประกาศฟังก์ชัน struct ชื่อ computer
// และกำหนดสมาชิก cpu, ram, ssd, gpu, และ psu
type computer struct {
	cpu string
	ram int
	ssd string
	gpu string
	psu int
}

// Struct 2 ประกาศตัวแปรชื่อ comespec1 และตัวแปรตัวที่สองชื่อ comspec2
// รับค่าข้อมูลของ struct ชื่อ computer
var comspec1 computer
var comspec2 computer

func main() {
	//ประกาศตัวแปรชื่อ phonedetail รับค่าข้อมูล Struct ชื่อ telephone
	// และกำหนดค่าให้กับสมาชิก
	phonedetail := telephone{
		name:     "iphone 15 Pro Max",
		spec:     "CPU A17 PRO Ram 8 Capacity 256 GB",
		price:    59900,
		warranty: 1,
	} //ประกาศตัวแปรชื่อ phonedetail กำหนดค่าให้สมาชิกแต่ละตัวของ telephone
	// name, spec, price และ warranty

	// Struct 2
	// กำหนดค่าของสมาชิก cpu, ram, ssd, gpu และ psu ให้กับ comspec1
	comspec1.cpu = "i9 14900k"
	comspec1.ram = 64
	comspec1.ssd = "2TB"
	comspec1.gpu = "RTX 4090 TI 24 GB"
	comspec1.psu = 1200

	// Struct 2
	// กำหนดค่าของสมาชิก cpu, ram, ssd, gpu และ psu ให้กับ comspec2
	comspec2.cpu = "ryzen9 7950x3D"
	comspec2.ram = 64
	comspec2.ssd = "2TB"
	comspec2.gpu = "RX 7900 XTX 24 GB"
	comspec2.psu = 850

	//Struct 1 พิมพ์สมาชิกข้อมูลของตัวแปรชื่อ phonedetail ออกจอภาพ
	fmt.Println("ชื่อเครื่อง ", phonedetail.name)
	fmt.Println("สเปค ", phonedetail.spec)
	fmt.Println("ราคา ", phonedetail.price)
	fmt.Println("การรับประกัน ", phonedetail.warranty)
	// พิมพ์ค่าสมาชิกข้อมูล name, spec, price, warraty ของตัวแปร phonedetail

	// struct 2 พิมพ์สมาชิกข้อมูลของตัวแปรชื่อ comspec1 ออกจอภาพ
	fmt.Println("CPU = ", comspec1.cpu)
	fmt.Println("RAM = ", comspec1.ram)
	fmt.Println("SSD = ", comspec1.ssd)
	fmt.Println("GPU = ", comspec1.gpu)
	fmt.Println("PSU = ", comspec1.psu)
	//พิมพ์ค่าสมาชิกข้อมูล cpu, ram, ssd, gpu และ psu ของตัวแปร comspec1

	// struct 2 พิมพ์สมาชิกข้อมูลของตัวแปรชื่อ comspec2 ออกจอภาพ
	fmt.Println("CPU = ", comspec2.cpu)
	fmt.Println("RAM = ", comspec2.ram)
	fmt.Println("SSD = ", comspec2.ssd)
	fmt.Println("GPU = ", comspec2.gpu)
	fmt.Println("PSU = ", comspec2.psu)
	// พิมพ์ค่าสมาชิกข้อมูล cpu, ram, ssd, gpu และ psu ของตัวแปร comspec1
}

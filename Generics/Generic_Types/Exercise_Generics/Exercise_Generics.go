package main

import (
	f "fmt"
)

/*
	1.สร้างโครงสร้างลิงก์ลิสต์

สร้างโครงสร้าง Node ที่มีฟิลด์ value และ next
สร้างโครงสร้าง LinkedList ที่มีฟิลด์ head
2.เพิ่มฟังก์ชันในโครงสร้าง LinkedList
Add(value int): เพิ่มโหนดใหม่ที่มีค่า value ไปยัง
ท้ายลิสต์
Print(): พิมพ์ค่าของโหนดทั้งหมดในลิสต์
Find(value int) bool: ค้นหาว่ามีโหนดที่มีค่า value ในลิสต์หรือไม่
Delete(value int) bool: ลบโหนดที่มีค่า value จากลิสต์
และคืนค่า true ถ้าลบสำเร็จ, false ถ้าไม่พบโหนดที่มีค่า value
*/
type Node struct {
	value int
	next  *Node
}
type LinkedList struct {
	head *Node
}

// เพิ่มฟังก์ชันในโครงสร้าง struct ชื่อ LinkedList
// เพิ่มโหนดใหม่ที่มีค่า value ไปยังท้ายลิสต์
func (list *LinkedList) Add(value int) {
	newNode := &Node{value: value}
	if list.head == nil {
		list.head = newNode
		return
	}
	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode

}

// ฟังก์ชัน Println()
func (list *LinkedList) Print() {
	for newNode := list.head; newNode != nil; newNode = newNode.next {
		f.Println(newNode.value, " ")
	}
	f.Println()
}

func (list *LinkedList) Find(value int) bool {
	current := list.head
	for current != nil {
		if current.value == value {
			return true
		}
		current = current.next
	}
	return false
}

func (list *LinkedList) Delete(value int) bool {
	if list.head == nil {
		return false
	}
	if list.head.value == value {
		list.head = list.head.next
		return true
	}
	current := list.head
	for current.next != nil {
		if current.next.value == value {
			current.next = current.next.next
			return true
		}
		current = current.next
	}
	return false
}

func main() {
	// สร้างลิงก์ลิสต์ใหม่
	intlist := LinkedList{}

	// เพิ่มค่าเข้าไปในลิสต์
	intlist.Add(10)
	intlist.Add(20)
	intlist.Add(30)
	intlist.Add(40)

	// พิมพ์ค่าทั้งหมดในลิสต์
	f.Println("ลิสต์หลังจากเพิ่มค่า: ")
	intlist.Print()

	// ค้นหาค่าในลิสต์
	f.Println("หาค่า 20: ", intlist.Find(20))
	f.Println("หาค่า 50: ", intlist.Find(50))

	// ลบค่าออกจากลิสต์
	intlist.Delete(20)
	f.Println("ลิสต์หลังจากลบ 20: ")
	intlist.Print()

	// ลองลบค่าที่ไม่มีในลิสต์
	intlist.Delete(50)
	f.Println("ลิสต์หลังจากพยายามลบ 50: ")
	intlist.Print()

}

/* สรุป
โจทย์นี้ช่วยให้ผู้เริ่มต้นเข้าใจและฝึกฝนการสร้างและใช้งาน
singly linked list ในภาษา Go โดยการเพิ่มฟังก์ชันการทำงาน
ต่างๆ เช่น การเพิ่มโหนด การค้นหา และการลบโหนดในลิสต์
รวมถึงการทดสอบการทำงานของฟังก์ชันเหล่านี้ในฟังก์ชัน main

*/

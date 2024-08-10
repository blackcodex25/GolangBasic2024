package main

import (
	f "fmt"
) /* การสร้างและการเดินทางผ่าน Binary Tree
เป้าหมาย
1.สร้างโครงสร้างข้อมูล Binary Tree
2.เพิ่มฟังก์ชันเพื่อเพิ่มโหนดในต้นไม้
3.เพิ่มฟังก์ชันในการเดินทางผ่านต้นไม้ (Pre-order, In-order,
Post-order traversal)
4.พิมพ์ค่าของโหนดทั้งหมดที่ได้จากการเดินทางผ่านต้นไม้
*/
/*
	โครงสร้าง Tree เป็นโหนดใน Binary Tree ที่มีค่าประเภท int

อยู่ในฟิลด์ Value
ฟิลด์ Left และ Right เป็นพอยน์เตอร์ไปยังโหนดลูกซ้ายและขวา
ตามลำดับ
โครงสร้างนี้เป็นพื้นฐานของ Binary Tree ที่ใช้เก็บข้อมูลแบบเชิง
ลำดับที่มีความสัมพันธ์ทางซ้าย-ขวา
*/
type Tree struct {
	Value int
	Left  *Tree
	Right *Tree
}

// *insert แทรกโหนดใหม่ที่มีค่าที่กำหนดลงในแผนผังไบนารี
// ฟังก์ชัน Insert() ทำหน้าที่แทรกค่าใหม่ Value ลงใน Bianry Tree
func (n *Tree) Insert(value int) {
	// หาก value น้อยกว่าหรือเท่ากับ n.Value
	// ตรวจสอบว่ามีโหนดลูกซ้าย (n.Left) อยู่หรือไม่
	if value <= n.Value {
		// ถ้าไม่มี โหนดลูกซ้ายฟิลด์ Left เท่ากับ nil
		if n.Left == nil {
			// จะสร้างโหนดลูกซ้ายใหม่ที่มีค่า value และเชื่อมโยงไปที่ n.Left
			n.Left = &Tree{Value: value}
		} else { // ถ้ามีโหนดลูกซ้ายแล้ว จะเรียก Insert() แบบเรียกซ้ำ (recursive) ในโหนดลูกซ้ายนั้น
			n.Left.Insert(value)
		}
	} else { // ถ้า value มากกว่า n.Value
		if n.Right == nil { // ถ้าไม่มี โหนดลูกขวาฟิลด์ Right เท่ากับ nil
			// จะสร้างโหนดลูกขวาใหม่ที่มีค่า value และเชื่อมโยงไปที่ n.Right
			n.Right = &Tree{Value: value}
		} else { // ถ้ามีโหนดลูกขวาแล้ว จะเรียก Insert() แบบเรียกซ้ำ (recursive) ในโหนดลูกขวานั้น
			n.Right.Insert(value)
		}
	}
}

// *ฟังก์ชัน InOrder, PreOrder, และ PostOrder

/*
	InOrder Traversal:

1.เยี่ยมชมต้นไม้ย่อยซ้าย (n.Left)
2.พิมพ์ค่าในโหนดปัจจุบัน (n.Value)
3.เยี่ยมชมต้นไม้ย่อยขวา (n.Right)
จะพิมพ์ค่าทั้งหมดใน Binary Tree
เรียงจากน้อยไปมากสำหรับ Binary Search Tree (BST)
*/
func (n *Tree) InOrder() {
	if n == nil {
		return
	}
	n.Left.InOrder()
	f.Println(n.Value, " ")
	n.Right.InOrder()
}

/*
	PreOrder traversal:

1.พิมพ์ค่าในโหนดปัจจุบัน (n.Value)
2.เยี่ยมชมต้นไม้ย่อยซ้าย (n.Left)
3.เยี่ยมชมต้นไม้ย่อยขวา (n.Right)
จะพิมพ์ค่าทั้งหมดโดยเริ่มจากโหนด root ตามด้วยต้นไม้ย่อยซ้าย
และขวาตามลำดับ
*/
func (n *Tree) PreOrder() {
	if n == nil {
		return
	}
	f.Println(n.Value, " ")
	n.Left.PreOrder()
	n.Right.PreOrder()

}

/*
	PostOrder traversal:

1.เยี่ยมชมต้นไม้ย่อยซ้าย (n.Left)
2.เยี่ยมชมต้นไม้ย่อยขวา (n.Right)
3.พิมพ์ค่าในโหนดปัจจุบัน (n.Value)
จะพิมพ์ค่าทั้งหมดโดยเริ่มจากต้นไม้ย่อยซ้ายและขวา แล้วจึงตามด้วยโหนด root
*/
func (n *Tree) PostOrder() {
	if n == nil {
		return
	}
	n.Left.PostOrder()
	n.Right.PostOrder()
	f.Println(n.Value, " ")
}

func main() {
	// เริ่มต้นด้วยการสร้างต้นไม้ที่มี root เป็นโหนดที่มีค่า 10
	root := &Tree{Value: 10}
	/* จากนั้นจะเพิ่มค่าในชุด Values
	ได้แก่ 5, 15, 3, 7, 13, 17 ลงในต้นไม้ */
	values := []int{5, 15, 3, 7, 13, 17}
	for _, value := range values {
		root.Insert(value)
	}
	/* ฟังก์ชันเดินทางผ่านต้นไม้ (Inorder, PreOrder, PostOrder)
	จะถูกเรียกและพิมพ์ค่าในต้นไม้ตามลำดับที่กำหนด
	*/
	f.Println("InOrder: ")
	root.InOrder()
	f.Println()

	f.Println("PreOrder: ")
	root.PreOrder()
	f.Println()

	f.Println("PostOrder: ")
	root.PostOrder()
	f.Println()
}

/* สรุป
เราได้เรียนรู้วิธีการสร้างโครงสร้างข้อมูล Binary Tree ใน Go
และการเพิ่มโหนดลงในต้นไม้
เราได้เรียนรู้วิธีการเดินทางผ่านต้นไม้ในแบบต่างๆ และการพิมพ์
ค่าของโหนด
เราสามารถใช้ฟังก์ชันเหล่านี้ในการแก้ปัญหาเกี่ยวกับต้นไม้ในอนาคต
ได้
โค้ดนี้เป็นตัวอย่างการทำงานกับ Binary Tree โดยใช้ฟังก์ชัน
Insert() ในการเพิ่มโหนดและฟังก์ชันต่างๆ ในการเดินทางผ่านต้นไม้
แต่ละ traversal มีลักษณะการเดินทางและการพิมพ์ค่าแตกต่างกันไป
ซึ่งช่วยในการตรวจสอบหรือจัดการข้อมูลในต้นไม้ในรูปแบบที่ต้องการ
การใช้งานฟังก์ชันเหล่านี้ช่วยให้เราเข้าใจถึงการจัดการและการประยุกต์
ใช้งาน Binary Tree ในปัญหาต่างๆ ได้อย่างมีประสิทธิภาพ
*/

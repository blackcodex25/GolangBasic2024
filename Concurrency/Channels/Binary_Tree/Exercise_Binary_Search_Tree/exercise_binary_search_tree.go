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
type Tree struct {
	Value int
	Left  *Tree
	Right *Tree
}

// insert แทรกโหนดใหม่ที่มีค่าที่กำหนดลงในแผนผังไบนารี
func (n *Tree) Insert(value int) {
	if value <= n.Value {
		if n.Left == nil {
			n.Left = &Tree{Value: value}
		} else {
			n.Left.Insert(value)
		}
	} else {
		if n.Right == nil {
			n.Right = &Tree{Value: value}
		} else {
			n.Right.Insert(value)
		}
	}
}

func (n *Tree) InOrder() {
	if n == nil {
		return
	}
	n.Left.InOrder()
	f.Println(n.Value, " ")
	n.Right.InOrder()
}

func (n *Tree) PreOrder() {
	if n == nil {
		return
	}
	f.Println(n.Value, " ")
	n.Left.PreOrder()
	n.Right.PreOrder()

}

func (n *Tree) PostOrder() {
	if n == nil {
		return
	}
	n.Left.PostOrder()
	n.Right.PostOrder()
	f.Println(n.Value, " ")
}

func main() {

}

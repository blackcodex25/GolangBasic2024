package main

import (
	"golang.org/x/tour/tree"
)

/* ในภาษา Go เราสามารถใช้ concurrent และ channels
เพื่อสร้างฟังก์ชันที่ตรวจสอบว่าต้นไม้ไบนารีสองต้นมีลำดับของค่าที่เก็บไว้
เหมือนกันหรือไม่ ต่อไปนี้เป็นคำอธิบายเกี่ยวเ
*/
// ประกาศโครงสร้างข้อมูล struct tree กำหนดฟิลด์สามตัว
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func New(k int) *Tree
func Walk(t *tree.Tree, ch chan int)

func main() {

}

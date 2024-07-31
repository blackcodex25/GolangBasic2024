package main

import (
	"fmt"
)

type Foo interface {
	Fizz()
}
type Bar struct{}

func (b *Bar) Fizz() {
	fmt.Println("fizz")
}
func Fizzy(foo Foo) {
	foo.Fizz()
}
func main() {
	b := Bar{}
	Fizzy(b) // เกิดข้อผิดพลาด บรรทัดที่ 20
}

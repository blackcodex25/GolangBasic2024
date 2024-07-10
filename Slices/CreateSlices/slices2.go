package main

/*การสร้าง slice ด้วย []datatype{values}
slice ใน Go สามารถทำได้โดยใช้รูปแบบ
[]datatpe{values} ซึ่งเป็นวิธีที่ง่ายและทั่วไปในการประกาศ
และกำหนดค่าเริ่มต้นให้กับ slice
*/
//slice_name := []datatype{values} Syntax How to use?
func main() {

	//ตัวอย่างการประกาศและกำหนดค่าเริ่มต้นให้กับ slice

	//ประกาศ slice ที่เป็น int และมีความยาว และความจุเป็น 0
	myslice := []int{}
	//ประกาศ slice ที่เป็น int และมีความยาวและความจุเป็น 3
	myslice2 := []int{1, 2, 3}

}

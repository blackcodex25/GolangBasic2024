package main

import (
	f "fmt"
)

// ตัวดำเนินการ การกำหนดค่า (Assignment Operators)
// การกำหนดค่าใช้สำหรับกำหนดค่าให้กับตัวแปร
// การกำหนดค่า (=) เพื่อกำหนดค่า 10 ให้กับตัวแปร x
var x = 10 //ประกาศตัวแปร x กำหนดค่า (=) 10
func main() {
	f.Println(x) // 10
}

/*รายการตัวดำเนินการ การกำหนดค่า
 = กำหนดค่า (Assignment)
+= บวกและกำหนดค่า (Addition Assignment)
-= ลบและกำหนดค่า (Subtraction Assignment)
*= คูณและกำหนดค่า (Multiplication Assignment)
/= หารและกำหนดค่า (Division Assignment)
%= หารเอาเศษและกำหนดค่า (Modulus Assignment)
*/

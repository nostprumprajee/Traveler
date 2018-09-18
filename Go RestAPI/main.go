package main

// นำเข้า package fmt มาใช้งาน
import (
	"errors"
	"fmt"
	"os"
)

type human struct {
	name string
	age  int
}

func (h human) printInfo() {
	fmt.Println(h.name, h.age)
}
func main() {
	// ภายใต้ fmt มีฟังก์ชันชื่อ Println ให้เราสามารถพิมพ์ข้อความออกจอได้
	fmt.Println("Hello, Go")
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	// switch i {
	// case 0:
	// 	fmt.Println("Zero")
	// case 1:
	// 	fmt.Println("One")
	// case 2:
	// 	fmt.Println("Two")
	// case 3:
	// 	fmt.Println("Three")
	// case 4:
	// 	fmt.Println("Four")
	// case 5:
	// 	fmt.Println("Five")
	// default:
	// 	fmt.Println("Unknown Number")
	// }
	names := [3]string{"Somchai", "Somsree", "Somset"}

	for index, name := range names {
		fmt.Println(index, name)
	}
	printFullName("Nost", "Test")
	// เพราะว่าฟังก์ชันคืนค่าสองค่า เราจึงประกาศตัวแปรมารองรับได้พร้อมกันสองตัว
	result, err := divide(5, 3)

	// ตรวจสอบก่อนว่ามี error ไหม ถ้ามีก็จบโปรแกรมไปแบบไม่ค่อยสวยด้วย Exit(1)
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(result)

	somchai := human{name: "Somchai", age: 23}
	somchai.printInfo() // Somchai 23
}
func printFullName(firstName string, lastName string) {
	fmt.Println(firstName + " " + lastName)
}

// คืนค่า float และ error ออกไปพร้อมกันจากฟังก์ชัน
func divide(dividend float32, divisor float32) (float32, error) {
	if divisor == 0.0 {
		err := errors.New("Division by zero!")
		return 0.0, err
	}

	return dividend / divisor, nil
}

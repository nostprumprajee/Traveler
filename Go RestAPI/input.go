package main // ประกาศ package ครับผม
import "fmt" // อันนี้ประกาศเพื่อใช้ในการเเสดงผล รับค่า อะไรทำนองนั้นนะครับ

func main() {
	fmt.Println("Hello world This golang - programming By Nost 😚")

	// ประกาศตัวแปร
	var input float64
	var myname string

	fmt.Print("Input number : ")

	// input ต้องใช้ Scanf
	fmt.Scanf("%f", &input)
	fmt.Println("ผลรับของการรับค่า ", input)
	if input == 100 {

		fmt.Print("What is your name?  ")
		fmt.Scanf("%s", &myname)

	} else {
		fmt.Println("ERROR")

	}

	fmt.Scanf("%s", &myname)
	fmt.Println("Hello", myname)
}

package main // ประกาศ package ครับผม
import "fmt" // อันนี้ประกาศเพื่อใช้ในการเเสดงผล รับค่า อะไรทำนองนั้นนะครับ

func main() {
	fmt.Println("Hello world This go - programming By Nost 😚")

	// . Variable etc .
	var number int = 100
	number2 := 200.34
	str := "String"
	fact := true
	var arr = [13]string{"😟", "😡", "😕", "😫", "😨", "😾", "🐊", "🌵", "🌍", "🎃", "🐕", "🙀", "😧"}

	fmt.Println(" . Variable Show . ")
	fmt.Println("====================")
	fmt.Println()
	fmt.Println("fact -> ", fact)
	fmt.Println("int number -> ", number)
	fmt.Println("float number 2 -> ", number2)
	fmt.Println("String str -> ", str)
	fmt.Println("arr -> ", arr)

}

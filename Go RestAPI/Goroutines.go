package main

import "sync"

// โปรดสังเกต เราต้องรับพอยเตอร์ของ sync.WaitGroup เข้ามาด้วย
func searchFromFolder(keyword string, folder string, wg *sync.WaitGroup) {
	// ทำการค้นหา
	// เมื่อค้นหาเสร็จ ต้องแจ้งให้ WaitGroup ทราบว่าเราทำงานเสร็จแล้ว
	// WaitGroup จะได้นับถูกว่าเหลือ Goroutines ที่ต้องรออีกกี่ตัว
	wg.Done()
}

func search(keyword string) {
	folders := [3]string{"Document", "Image", "Library"}
	var wg sync.WaitGroup
	wg.Add(len(folders))
	for _, folder := range folders {
		// เราต้องส่ง reference ของ wg ไปด้วย เพื่อที่จะสั่ง Done
		go searchFromFolder(keyword, folder, &wg)
	}
	wg.Wait()
}

func main() {
	search("dog")
}

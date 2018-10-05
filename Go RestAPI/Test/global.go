// ใช้ค่าคงที่ ไม่ดีนะ
const port = 8080

// ใช้แบบนี้ดีกว่า
var defaultPort = 8080

// ใช้แบบนี้ก็ดีกว่า
type App struct {
	Port int
}
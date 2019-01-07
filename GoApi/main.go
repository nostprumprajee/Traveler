// main.go

package main

func main() {
	a := App{}
	a.Initialize("root", "Nost593214!", "rest_api_example")

	a.Run(":8000")
}

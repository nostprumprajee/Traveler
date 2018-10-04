package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	const jsonStream = `
		{"ID": 1, "Title": "title 1", "Done": false}
		{"ID": 2, "Title": "title 2", "Done": false}
	`
	type Task struct {
		ID    int64
		Title string
		Done  bool
	}

	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var task Task
		if err := dec.Decode(&task); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d: %s : %v\n", task.ID, task.Title, task.Done)
	}
}

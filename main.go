package main

import (
	"fmt"
	"math/rand"
	"time"
)

var cache = map[int]Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	MaxId := 10

	for i := 0; i < MaxId; i++ {
		id := rnd.Intn(MaxId) + 1

		// Two goroutines using anonymous function
		// and a hacky sleep to await their result
		go func(id int) {
			if b, ok := queryCache(id); ok {
				fmt.Println("From cache")
				fmt.Println(b)
			}
		}(id)

		go func(id int) {
			if b, ok := queryDatabase(id); ok {
				fmt.Println("From the database")
				fmt.Println(b)
			}
		}(id)

		fmt.Println("Book not found with id:", id)
		time.Sleep(150 * time.Millisecond)
	}
}

// A function that return multiple values
// That must be the same logic used to build functions that return a value
// and an error
func queryCache(id int) (Book, bool) {
	b, ok := cache[id]
	return b, ok
}

func queryDatabase(id int) (Book, bool) {
	time.Sleep(100 * time.Millisecond)

	// Iterate over a slice
	for _, b := range books {
		if b.ID == id {
			cache[id] = b
			return b, true
		}
	}

	return Book{}, false
}

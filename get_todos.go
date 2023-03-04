package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

const (
	todoServer = "https://jsonplaceholder.typicode.com/todos"
)

func getTodo(index int) (string, error) {
	time.Sleep(500 * time.Millisecond)
	resp, err := http.Get(fmt.Sprintf("%s/%d", todoServer, index))
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Error getting todo %d: %s", index, resp.Status)
	}
	defer resp.Body.Close()

	var todo struct {
		Title string `json:"title"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&todo); err != nil {
		return "", err
	}

	return todo.Title, nil
}

func main() {
	numTodos, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Error parsing number of todos: %v", err)
	}

	var wg sync.WaitGroup
	todos := make([]string, numTodos)

	for i := 0; i < numTodos; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			todo, err := getTodo(j + 1)
			if err != nil {
				log.Printf("Error getting todo %d: %v", j+1, err)
				return
			}
			todos[j] = todo
		}(i)
	}
	wg.Wait()

	fmt.Println("# Todos")
	for _, todo := range todos {
		if todo == "" {
			continue
		}
		fmt.Printf("* %s\n", todo)
	}
}

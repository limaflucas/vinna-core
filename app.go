package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"vinna.app/vinna-app/internal/frequency"
	"vinna.app/vinna-app/internal/habit"
	"vinna.app/vinna-app/internal/task"
	"vinna.app/vinna-app/pkg/db"
)

func main() {
	db_conn := db.GetConnection()
	var sqlR string
	err := db_conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&sqlR)
	if err != nil {
		fmt.Printf("QueryRow failed: %v\n", err)
		db.Close()
		os.Exit(1)
	}
	fmt.Println(sqlR)

	var kb string
	fmt.Scanln(&kb)
	for kb != "q" {
		freq, err := frequency.New("29,59", "*", "*", "*", "*")
		if err != nil {
			fmt.Printf("Error: %v/n", err)
			return
		}
		t1, err := task.New(uuid.New(), "a task", time.Now(), time.Now().Add(30), task.Recorded, *freq)
		if err != nil {
			fmt.Printf("Error: %v", err)
		}
		h1, err := habit.New(uuid.New(), "a habit", time.Now(), time.Now(), habit.Active, *freq)
		if err != nil {
			fmt.Printf("Error: %v", err)
		}

		fmt.Println(t1)
		fmt.Println(h1)
		fmt.Scanln(&kb)
	}
	db.Close()
}

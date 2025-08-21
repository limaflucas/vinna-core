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

	var kb string
	fmt.Scanln(&kb)
	for kb != "q" {
		db_conn := db.GetConnection()
		freq, error := frequency.New("29,59", "*", "*", "*", "*")
		if error != nil {
			fmt.Println(error)
			return
		}
		var t1 = task.New(uuid.New(), "a task", time.Now(), time.Now().Add(30), task.Recorded, *freq)
		var h1 = habit.New(uuid.New(), "a habit", time.Now(), time.Now(), habit.Active, *freq)
		fmt.Println(t1)
		fmt.Println(h1)

		var sqlR string
		err := db_conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&sqlR)
		if err != nil {
			fmt.Printf("QueryRow failed: %v\n", err)
			db.Close()
			os.Exit(1)
		}
		fmt.Println(sqlR)
		fmt.Scanln(&kb)
	}
	db.Close()
}

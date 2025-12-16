package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	// .env не обязателен; если файла нет — ошибка игнорируется
	_ = godotenv.Load()

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		// fallback — прямой DSN в коде с UTF8 кодировкой
		dsn = "postgres://postgres:P@ssw0rd@localhost:5432/todo?sslmode=disable&client_encoding=UTF8"
	}

	db, err := openDB(dsn)
	if err != nil {
		log.Fatalf("openDB error: %v", err)
	}
	defer db.Close()

	repo := NewRepo(db)

	// 1) Вставим пару задач по одной
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	titles := []string{"Сделать ПЗ №5", "Купить кофе", "Проверить отчёты"}
	for _, title := range titles {
		id, err := repo.CreateTask(ctx, title)
		if err != nil {
			log.Fatalf("CreateTask error: %v", err)
		}
		log.Printf("Inserted task id=%d (%s)", id, title)
	}

	// 2) Массовая вставка через транзакцию
	ctxMany, cancelMany := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelMany()

	batchTitles := []string{"Задача 1 из батча", "Задача 2 из батча", "Задача 3 из батча"}
	if err := repo.CreateMany(ctxMany, batchTitles); err != nil {
		log.Printf("CreateMany error: %v", err)
	} else {
		log.Println("Batch insert completed successfully")
	}

	// 3) Прочитаем ВСЕ задачи
	ctxList, cancelList := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelList()

	tasks, err := repo.ListTasks(ctxList)
	if err != nil {
		log.Fatalf("ListTasks error: %v", err)
	}

	// 4) Напечатаем все задачи
	fmt.Println("\n=== Все задачи ===")
	for _, t := range tasks {
		fmt.Printf("#%d | %-24s | done=%-5v | %s\n",
			t.ID, t.Title, t.Done, t.CreatedAt.Format(time.RFC3339))
	}

	// 5) Получим только НЕВЫПОЛНЕННЫЕ задачи
	ctxUndone, cancelUndone := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelUndone()

	undoneTasks, err := repo.ListDone(ctxUndone, false)
	if err != nil {
		log.Printf("ListDone (undone) error: %v", err)
	} else {
		fmt.Println("\n=== Невыполненные задачи ===")
		for _, t := range undoneTasks {
			fmt.Printf("#%d | %-24s | created: %s\n",
				t.ID, t.Title, t.CreatedAt.Format("2006-01-02 15:04"))
		}
	}

	// 6) Найдем задачу по конкретному ID
	ctxFind, cancelFind := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFind()

	// Например, ищем задачу с ID=2
	taskID := 2
	task, err := repo.FindByID(ctxFind, taskID)
	if err != nil {
		log.Printf("FindByID error for id=%d: %v", taskID, err)
	} else {
		fmt.Printf("\n=== Детали задачи #%d ===\n", taskID)
		fmt.Printf("ID:        %d\n", task.ID)
		fmt.Printf("Заголовок: %s\n", task.Title)
		fmt.Printf("Статус:    %v\n", task.Done)
		fmt.Printf("Создана:   %s\n", task.CreatedAt.Format(time.RFC3339))
		fmt.Printf("Создана:   %s назад\n", time.Since(task.CreatedAt).Round(time.Second))
	}

	// 7) Пример обновления задачи как выполненной
	// (для этого нужно добавить метод UpdateTask в репозиторий)
	fmt.Println("\n=== Проверка завершена ===")
}

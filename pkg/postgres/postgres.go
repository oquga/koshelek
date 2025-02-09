package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

// CheckConnection проверяет подключение к PostgreSQL
func CheckConnection(connStr string) error {
	// Открываем соединение с базой данных
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("failed to open database connection: %w", err)
	}
	defer db.Close()

	// Устанавливаем лимит времени для подключения
	db.SetConnMaxLifetime(time.Minute * 1)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	// Пытаемся подключиться к базе данных
	err = db.Ping()
	if err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	// Если подключение успешно, выводим сообщение
	log.Println("Successfully connected to PostgreSQL!")
	return nil
}

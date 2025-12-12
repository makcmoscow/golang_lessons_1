package main

import(
	// "fmt"
	"log"
	"context"
	"Bankstore/api"
	"github.com/jackc/pgx/v5/pgxpool"
	db "Bankstore/db/sqlc"
)

const (
    dbSource = "postgresql://app_user:pswd@localhost:5432/bankdb?sslmode=disable"
	serverAddress = "0.0.0.0:8080"

)

func main() {
	    // Соединение с БД
	var err error
    pool, err := pgxpool.New(context.Background(), dbSource)
    if err != nil {
        log.Fatal("can not connect to db", err)
    }
    // Закрываем соединение
    defer pool.Close()
    // Сщздаем необходимые экземпляры для работы
	store := db.NewStore(pool) // Хранилище
	server := api.NewServer(store) // роутинг и прочее

	err = server.Start(serverAddress)
		if err != nil {
			log.Fatal("Can't start server", err)
		}

}
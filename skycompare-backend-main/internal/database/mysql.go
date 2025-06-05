package database

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func parseDSN(databaseURL string) (string, error) {
	u, err := url.Parse(databaseURL)
	if err != nil {
		return "", err
	}

	user := u.User.Username()
	password, _ := u.User.Password()
	host := u.Host
	dbName := strings.TrimPrefix(u.Path, "/")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, host, dbName)
	return dsn, nil
}

func Connect() *sql.DB {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL no está configurada")
	}

	dsn, err := parseDSN(databaseURL)
	if err != nil {
		log.Fatal("Error parseando DSN:", err)
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error conectando a MySQL:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Error haciendo ping a MySQL:", err)
	}

	log.Println("Conexión a MySQL establecida")
	return db
}

package databases

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// BUAT KONEKSI KE DATABASE
func InitMariaDB() *sql.DB {
	// 1. Define the connection string (The "Key" to the DB)
	// username:password@tcp(host:port)/dbname
	dsn := "root:Yuna1070@tcp(127.0.0.1:3306)/sispa_db?parseTime=true"

	// 2. Open the connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("Failed to open database: %v", err))
	}

	// 3. Set connection pool limits (Stark-level optimization)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Hour)

	// 4. Test the connection
	if err := db.Ping(); err != nil {
		panic(fmt.Sprintf("Database is unreachable: %v", err))
	}

	fmt.Println("Successfully connected to MariaDB!")
	return db
}


package db

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	var err error
	dsn := os.Getenv("DB_URL")

	for i := 0; i < 10; i++ {  // retry 10 times
		DB, err = sql.Open("postgres", dsn)
		if err != nil {
			log.Println("DB open error:", err)
			time.Sleep(2 * time.Second)
			continue
		}

		err = DB.Ping()
		if err == nil {
			log.Println("✅ Connected to DB")
			return
		}

		log.Println("⏳ Waiting for DB...")
		time.Sleep(2 * time.Second)
	}

	log.Fatal("❌ Could not connect to DB")
}
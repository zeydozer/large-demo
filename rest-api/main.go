package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	_ "github.com/lib/pq"
)

var (
	db    	*sql.DB
	replica *sql.DB
	rdb   	*redis.Client
	ctx   	= context.Background()
)

func main() {
	dbURL := os.Getenv("DATABASE_URL")
	var err error
	db, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	// Db optimizasyonları
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)

	replica, err = sql.Open("postgres", os.Getenv("REPLICA_URL"))
	if err != nil {
		log.Fatal(err)
	}

	// Redis bağlantısı
	rdb = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	http.HandleFunc("/users/", getUserHandler)
	log.Println("API running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/users/"):]
	cacheKey := "user:" + id

	// 1. Önce Redis’ten oku
	val, err := rdb.Get(ctx, cacheKey).Result()
	if err == nil {
		// Cache’den geldi!
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(val))
		return
	}

	// 2. Cache’de yoksa veritabanından çek
	var name, email string
	err = replica.QueryRow("SELECT name, email FROM users WHERE id=$1", id).Scan(&name, &email)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	resp := map[string]string{"id": id, "name": name, "email": email}
	respJson, _ := json.Marshal(resp)

	// 3. Sonucu cache’e yaz (ör. 1 dakika süreyle)
	rdb.Set(ctx, cacheKey, respJson, 60*time.Second)

	w.Header().Set("Content-Type", "application/json")
	w.Write(respJson)
}

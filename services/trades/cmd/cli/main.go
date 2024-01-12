package main

import (
	"database/sql"
	"fmt"
	"log"
	"mabna/services/trades/configs"
	"mabna/shared/configext"
	"math/rand"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	var err error
	conf := new(configs.Configs)
	loadedConfigs, err := configext.LoadConfigs("", configs.DefaultConfig, true, conf)
	if err != nil {
		log.Fatal(err)
	}
	conf = loadedConfigs.(*configs.Configs)
	dsn := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=disable",
		conf.Postgres.User, conf.Postgres.Pass, conf.Postgres.Host, conf.Postgres.Port, conf.Postgres.DB)
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Usage: ./cli <number_of_records>")
		return
	}

	numRecords, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid input for number_of_records")
		return
	}

	start := time.Now()
	insertRandomTrades(numRecords)
	elapsed := time.Since(start)
	fmt.Printf("Inserted %d records in %s\n", numRecords, elapsed)
}

func insertRandomTrades(numRecords int) {
	rand.Seed(time.Now().UnixNano())

	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("Failed to start transaction: %v", err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`INSERT INTO Trade (ID, InstrumentId, DateEn, Open, High, Low, Close) VALUES ($1, $2, $3, $4, $5, $6, $7)`)
	if err != nil {
		log.Fatalf("Failed to prepare statement: %v", err)
	}
	defer stmt.Close()

	for i := 0; i < numRecords; i++ {
		id := rand.Int31()
		instrumentID := rand.Intn(2) + 1 // Random instrument ID between 1 and 2
		dateEn := randomDate()
		open := rand.Float64() * 1000
		high := open + (rand.Float64() * 100)
		low := open - (rand.Float64() * 100)
		close := low + (rand.Float64() * (high - low))

		_, err := stmt.Exec(id, instrumentID, dateEn, open, high, low, close)
		if err != nil {
			log.Printf("Failed to insert record: %v", err)
			continue
		}
	}

	if err := tx.Commit(); err != nil {
		log.Fatalf("Failed to commit transaction: %v", err)
	}
}

func randomDate() string {
	min := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0).Format("2006-01-02")
}

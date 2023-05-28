package cg_test

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"cg-new/internal/cg"
)

const (
	DEFAULT_QTY = 1000000
)

var (
	content []string
)

func TestMain(m *testing.M) {
	fmt.Println(time.Now())

	code := m.Run()
	os.Exit(code)
}

func TestCg(t *testing.T) {
	config := cg.NewConfig()
	config.Qty = DEFAULT_QTY

	planFile := cg.NewPlanFile(config)
	fmt.Println(planFile.FileName)
	binaryTree := cg.NewBinaryTree(config, planFile)
	binaryTree.Start()

	file, err := os.Open(planFile.FileName)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		file.Close()
		os.Remove(planFile.FileName)
	}()

	content = []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		t.Fatal(err)
	}

	t.Run("Count", checkCount)
	t.Run("Unique", checkUnique)
}

func checkCount(t *testing.T) {
	if len(content) != DEFAULT_QTY {
		t.Fatal()
	}
}

func checkUnique(t *testing.T) {
	const DB_NAME = "codes_test.sqlite"

	db, err := sql.Open("sqlite3", DB_NAME)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		db.Close()
		os.Remove(DB_NAME)
	}()

	db.Exec(`
DROP TABLE IF EXISTS codes
		`)
	db.Exec(`
CREATE TABLE codes (
	code TEXT NOT NULL UNIQUE
) STRICT
		`)

	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}

	stmt, err := tx.Prepare(`
INSERT INTO codes (code) VALUES (?)
		`)
	if err != nil {
		t.Fatal(err)
	}

	for _, s := range content {
		_, err = stmt.Exec(s)
		if err != nil {
			tx.Rollback()
			t.Fatalf("%s, code: %s", err, s)
		}
	}

	err = tx.Commit()
	if err != nil {
		t.Fatal(err)
	}
}

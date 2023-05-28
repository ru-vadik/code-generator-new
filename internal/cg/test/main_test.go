package cg_test

import (
	"bufio"
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"cg-new/internal/cg"
)

const (
	DEFAULT_QTY = 1000
	DB_NAME     = "codes_test.sqlite"
)

var (
	content []string
	db      *sql.DB
)

func TestMain(m *testing.M) {
	var err error
	db, err = sql.Open("sqlite3", DB_NAME)
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

	code := m.Run()
	os.Exit(code)
}

func TestCg(t *testing.T) {
	config := cg.NewConfig()
	planFile := cg.NewPlanFile(config)
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
			t.Fatal(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		t.Fatal(err)
	}
}

package db_test

import (
	"database/sql"
	"github.com/dumunari/go-ports-and-adapters/adapters/db"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	creatTable(Db)
	createProduct(Db)
}

func creatTable(db *sql.DB) {
	table := `CREATE TABLE products (
			"id" string, 
			"name" string, 
			"price" float, 
			"status" string
			);`
	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatalf(err.Error())
	}
	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `insert into products values("testeID", "Product Test", 0, "disabled")`
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatalf(err.Error())
	}
	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer Db.Close()

	productDb := db.NewProductDb(Db)
	product, err := productDb.Get("testeID")
	require.Nil(t, err)
	require.Equal(t, "Product Test", product.GetName())
	require.Equal(t, 0.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}
package database

import "testing"

func TestConnectSQLite(t *testing.T) {

	db, err := ConnectSQLite(":memory:")

	if err != nil {
		t.Fatalf("error conectando sqlite: %v", err)
	}

	if db == nil {
		t.Fatal("db no debería ser nil")
	}
}

package db

import (
  "database/sql"
  "fmt"
  pq "github.com/lib/pq"
)

var DB *sql.DB

func InitDB(connStr string) error{
  var err error
  DB, err = sql.Open("postgres", connStr)

  if err != nil {
    return fmt.Errorf("failed to open db: %w", err)
  }

  if err = DB.Ping(); err != nil{
      return fmt.Errorf("failed to ping db: %w", err)
  }

  return nil
}


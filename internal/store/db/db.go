package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"goth/internal/aws"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type tursoConn struct {
  Url    string
  Token  string
}

func open(dbName string) (*sql.DB, error) {
  sm, err := aws.GetSMClient("us-east-1")
  if err != nil { return nil, err }
  
  id := "turso"
  ss, err := sm.GetSecret(&id)
  
  conn := &tursoConn{}
  err = json.Unmarshal([]byte(*ss), conn)
  if err != nil { return nil, err }

  url := fmt.Sprintf("libsql://%s.turso.io?authToken=%s", dbName, conn.Token)
  db, err := sql.Open("libsql", url)
  return db, err
}

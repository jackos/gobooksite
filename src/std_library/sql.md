
---
```go
// ## Importing external packages not working with yaegi, however it works with gomacro. 
// This won't work with gobook until kernel changed to use gomacro
import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)
```
---
---
```go
func DoSomeInserts(ctx context.Context, db *sql.DB, value1 string, value2 string, value3 string) (err error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		// If the transaction fails roll back
		if err != nil {
			tx.Rollback()
		}
		err = tx.Commit()
	}()
	_, err = tx.ExecContext(ctx, "INSERT INTO area (building_id, name, description, created_at, updated_at) values (?, ?, ?, NULL, NULL)", value1, value2, value3)
	if err != nil {
		return err
	}
	// use tx to do more database inserts here
	return nil
}
```
---
---
```go
	db, err := sql.Open("mysql", "root:password@/test")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	// Use context.Background for the parent context
	err = DoSomeInserts(context.Background(), db, "2", "test", "test_description")
	if err != nil {
		log.Fatal(err)
	}
```
---

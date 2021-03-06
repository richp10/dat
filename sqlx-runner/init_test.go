package runner

import (
	"bytes"
	"database/sql"
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/richp10/dat"
	"github.com/richp10/dat/kvs"
	"github.com/richp10/dat/postgres"
)

var testDB *DB
var sqlDB *sql.DB

func init() {
	dat.Dialect = postgres.New()
	sqlDB = realDb()
	testDB = NewDB(sqlDB, "postgres")
	dat.Strict = false
	log.SetLevel(log.WarnLevel)

	Cache = kvs.NewMemoryKeyValueStore(1 * time.Second)
	//Cache, _ = kvs.NewDefaultRedisStore()
}

func beginTxWithFixtures() *Tx {
	installFixtures()
	c, err := testDB.Begin()
	if err != nil {
		panic(err)
	}
	return c
}

func quoteColumn(column string) string {
	var buffer bytes.Buffer
	dat.Dialect.WriteIdentifier(&buffer, column)
	return buffer.String()
}

func quoteSQL(sqlFmt string, cols ...string) string {
	args := make([]interface{}, len(cols))

	for i := range cols {
		args[i] = quoteColumn(cols[i])
	}

	return fmt.Sprintf(sqlFmt, args...)
}

func realDb() *sql.DB {
	driver := os.Getenv("DAT_DRIVER")
	if driver == "" {
		log.Fatal("env DAT_DRIVER is not set")
	}

	dsn := os.Getenv("DAT_DSN")
	if dsn == "" {
		log.Fatal("env DAT_DSN is not set")
	}

	db, err := sql.Open(driver, dsn)
	if err != nil {
		log.Fatal("Database error ", "err", err)
	}

	return db
}

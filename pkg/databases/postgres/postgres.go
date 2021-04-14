package postrges

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
)

type DB struct {
	Postgres *sql.DB
}

var dbConn = &DB{}

func ConnectPostgeSQL(host, port, username, pass, dbname string) (*DB, error) {
	fmt.Println("User is: ", username)
	fmt.Println("Password is: ", pass)
	fmt.Println("Host is: ", host)
	fmt.Println("Port is: ", port)
	fmt.Println("Dbname is: ", dbname)
	var conStr string
	if host == "" && port == "" && dbname == "" {
		return nil, errors.New("cannot estabished the connection")
	}
	conStr = fmt.Sprintf("host=%s user=%s password=%s port=%s database=%s",
		host,
		username,
		pass,
		port,
		dbname,
	)

	db, err := sql.Open("pgx", conStr)
	if err != nil {
		log.Println("Cannot connect to postgreSQL got error: ", err)
		return nil, err
	}
	configureConnectionPool(db)
	err = db.Ping()
	if err != nil {
		log.Println("PostgreSQL ping got error: ", err)
		return nil, err
	}

	dbConn.Postgres = db
	return dbConn, nil
}

func DisconnectPostgres(db *sql.DB) {
	db.Close()
	log.Println("Connected with postgres has closed")
}

func configureConnectionPool(dbPool *sql.DB) {
	// [START cloud_sql_postgres_databasesql_limit]

	// Set maximum number of connections in idle connection pool.
	dbPool.SetMaxIdleConns(5)

	// Set maximum number of open connections to the database.
	dbPool.SetMaxOpenConns(7)

	// [END cloud_sql_postgres_databasesql_limit]

	// [START cloud_sql_postgres_databasesql_lifetime]

	// Set Maximum time (in seconds) that a connection can remain open.
	dbPool.SetConnMaxLifetime(1800)

	// [END cloud_sql_postgres_databasesql_lifetime]
}

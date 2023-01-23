package postgres

import (
    "context"
    "database/sql"
    "fmt"
    "log"
    "os"
    "time"
)

var (
    DB *sql.DB
    counts int64
)
const dbTimeout = time.Second * 3
func openDB(dsn string)(*sql.DB, error){
    db, err := sql.Open("pgx", dsn)
    if err != nil {
        return nil, err
    }
    err = db.Ping()
    if err != nil {
        return nil, err
    }
    return db, nil
}
func ConnectToDB(){
    dsn := os.Getenv("DSN")
    for {
        connection, err := openDB(dsn)
        if err != nil {
            log.Println("Postgress not yet ready...")
            counts++
        } else {
            log.Println("Connected to Postgres")
            SetupDbTable(connection)
            DB = connection
            return
        }
        if counts > 10 {
            log.Println(err)
            return
        }
        log.Println("Waiting for 2 seconds...")
        time.Sleep(time.Second * 2)
        continue
    }
}
func SetupDbTable(conn *sql.DB){
    stmt := `CREATE TABLE IF NOT EXISTS public.users (
id int GENERATED ALWAYS AS IDENTITY NOT NULL,
name varchar(20) NOT NULL,
created timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,

CONSTRAINT PK_users PRIMARY KEY(id)
);
CREATE TABLE IF NOT EXISTS public.games (
id int GENERATED ALWAYS AS IDENTITY NOT NULL,
name varchar(20) NOT NULL,
created timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,

CONSTRAINT PK_games PRIMARY KEY(id)
);
CREATE TABLE IF NOT EXISTS public.combinations (
id int GENERATED ALWAYS AS IDENTITY NOT NULL,
name varchar(20) NOT NULL,

CONSTRAINT PK_combinations PRIMARY KEY(id)
);
`
    _, err := conn.Exec(stmt)
    if err != nil {
        fmt.Printf("couldn't run query:%s", err.Error())
        return
    }
    return
}




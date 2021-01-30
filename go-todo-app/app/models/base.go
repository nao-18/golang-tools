package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"go-tools/go-todo-app/config"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

//テーブル作成コード
var Db *sql.DB

var err error

const (
	tableNameUser = "users"
	tableNameTodo = "todos"
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	cmd := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		email STRING,
		password STRING,
		created_at DATETIME)`, tableNameUser)

	Db.Exec(cmd)

	cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT,
		user_id INTEGER,
		created_at DATETIME)`, tableNameTodo)
	Db.Exec(cmdT)
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (crytext string) {
	crytext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return crytext
}

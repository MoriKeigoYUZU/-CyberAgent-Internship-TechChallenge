package userModel

import (
	"database/sql"
	"log"

	"github.com/2009_proto_h_server/pkg/db"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name      string
	PassWord  string
	AuthToken string
	Coin      int32
}

// InsertUser データベースをレコードを登録する
func InsertUser(record *User) error {
	stmt, err := db.Conn.Prepare("INSERT INTO user (name, password, auth_token, coin) VALUES (?, ?, ?, ?); ")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(record.Name, record.PassWord, record.AuthToken, record.Coin)
	return err
}

// SelectUser １レコード分取得
func SelectUser(name, password string) (*User, error) {

	row := db.Conn.QueryRow("SELECT name, auth_token, coin FROM user WHERE name = ? AND password = ?; ", name, password)
	return convertToUser(row)

}

// SelectName トークンを用いて名前を取得
func SelectName(token string) (*User, error) {
	row := db.Conn.QueryRow("SELECT name FROM user WHERE auth_token = ?; ", token)
	return convertToName(row)
}

// SelectUserByAuthToken auth_tokenを条件にレコードを取得する
func SelectUserByAuthToken(authToken string) (*User, error) {
	row := db.Conn.QueryRow("SELECT id, auth_token, name, high_score, coin FROM user WHERE auth_token = ?; ", authToken)
	return convertToUser(row)
}

// convertToUser rowデータをUserデータへ変換する
func convertToUser(row *sql.Row) (*User, error) {
	user := User{}
	err := row.Scan(&user.Name, &user.AuthToken, &user.Coin)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}
	return &user, nil
}

// convertToName rowデータをUserデータへ変換する
func convertToName(row *sql.Row) (*User, error) {
	user := User{}
	err := row.Scan(&user.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}
	return &user, nil
}

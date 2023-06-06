package database

import (
	"database/sql"
	"fmt"
	"log"
	"test/service/model"

	_ "github.com/go-sql-driver/mysql"
)

var cn *sql.DB

func ConnectionDb(service model.Service) {
	log.Println("Connecting database...")

	connString := fmt.Sprintf("%v:%v@tcp(%v)/%v", service.Cfg.UserMysql, service.Cfg.PasswordMysql, service.Cfg.AddressMysql, service.Cfg.DbMysql)
	connection, err := sql.Open("mysql", connString)
	if err != nil {
		log.Fatalf("Dattabase no connected: %s", err)
	}
	log.Println("Database connected")
	if err = connection.Ping(); err != nil {
		log.Fatalf("Database no answer: %s", err)
	}
	log.Println("Database answer")

	cn = connection
}

func GetAuthorsByBook(title string) ([]string, error) {
	log.Println("Get authors")
	query := fmt.Sprintf("Select author from Books where title='%v'", title)
	rows, err := cn.Query(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var tmp string
	var res []string
	for rows.Next() {
		rows.Scan(&tmp)
		res = append(res, tmp)
	}
	log.Println("Get successful")
	return res, nil
}

func GetBooksByAuthor(name string) ([]string, error) {
	log.Println("Get books")
	query := fmt.Sprintf("Select title from Books where author='%v'", name)
	rows, err := cn.Query(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var tmp string
	var res []string
	for rows.Next() {
		rows.Scan(&tmp)
		res = append(res, tmp)
	}
	log.Println("Get successful")
	return res, nil
}

func DisconnectionDB() {
	cn.Close()
	log.Println("Database close")
}

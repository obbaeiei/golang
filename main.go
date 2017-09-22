package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// docker run --name some-mysql -e MYSQL_ROOT_PASSWORD=my-secret-pw -p 3306:3306 -d mysql
// docker run -it --link some-mysql:mysql --rm mysql sh -c 'exec mysql -h"$MYSQL_PORT_3306_TCP_ADDR" -P"$MYSQL_PORT_3306_TCP_PORT" -uroot -p"$MYSQL_ENV_MYSQL_ROOT_PASSWORD"'
// insert into accounts (name, email)
// values ("pallat", "yod.pallat@gmail.com");
// http://jinzhu.me/gorm/associations.html

type Account struct {
	Name  string
	Email string
}

func main() {
	db, err := gorm.Open("mysql", "root:my-secret-pw@/golang?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	user := Account{Name: "Jinzhu", Email: "fake1@gmail.com"}

	db.NewRecord(user) // => returns `true` as primary key is blank

	db.Create(&user)

	db.NewRecord(user)
}

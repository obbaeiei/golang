package main

import (
	"fmt"
	"log"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

//docker run --name some-mongo -d -p 27017:27017 mongo

type data struct {
	Name string
}

var (
	url        = "mongodb://localhost:27017"
	database   = "database"
	collection = "collection"
	query      = bson.M{"name": "pallat"}
	result     data
)

func main() {
	session, err := mgo.Dial(url)
	if err != nil {
		log.Fatal(err)
	}

	defer session.Close()

	c := session.DB(database).C(collection)

	// err = c.Insert(bson.M{"name": "pallat"})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	err = c.Find(query).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}

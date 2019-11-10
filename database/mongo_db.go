package database

import (
	"errors"
	"log"

	"github.com/globalsign/mgo"
)

type MongoDBRepository struct {
	Session        *mgo.Session
	DatabaseName   string
	CollectionName string
}

func (db MongoDBRepository) accessToDBCollection() *mgo.Collection {
	return db.Session.DB(db.DatabaseName).C(db.CollectionName)
}

func (db MongoDBRepository) Select(selector interface{}, output interface{}) error {
	if selector == nil {
		log.Println("Selector is nil")
	}

	if err := db.accessToDBCollection().Find(selector).All(output); err != nil {
		log.Println("MongoDB Find All has error:", err)
		return err
	}

	return nil
}

func (db MongoDBRepository) Insert(payload interface{}) error {
	if payload == nil {
		log.Println("No payload")
		return nil
	}

	if err := db.accessToDBCollection().Insert(payload); err != nil {
		log.Println("MongoDB Insert has error:", err)
		return err
	}

	return nil
}

func (db MongoDBRepository) Update(selector interface{}, updater interface{}) error {
	if selector == nil {
		errMessage := "Selector is nil"
		log.Println(errMessage)
		return errors.New(errMessage)
	}

	if err := db.accessToDBCollection().Update(selector, updater); err != nil {
		log.Println("MongoDB Update has error:", err)
		return err
	}

	return nil
}

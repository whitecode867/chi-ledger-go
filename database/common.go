package database

import (
	"github.com/globalsign/mgo"
)

type Session struct {
	MongoDBSession *mgo.Session
}

func (s *Session) CloseAllDatabases() {
	s.MongoDBSession.Close()
}

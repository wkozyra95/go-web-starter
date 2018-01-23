package db

import (
	conf "github.com/wkozyra95/go-web-starter/config"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func SetupDB(config conf.Config) (func() DB, error) {
	session, sessionErr := mgo.Dial(config.DbURL)
	if sessionErr != nil {
		return nil, sessionErr
	}

	return func() DB {
		sessionClone := session.Clone()
		db := session.DB("")
		return DB{
			session: sessionClone,
			User: func() Collection {
				return collection{
					collection: db.C(userCollection),
				}
			},
			Project: func() Collection {
				return collection{
					collection: db.C(projectCollection),
				}
			},
		}
	}, nil
}

type DB struct {
	session *mgo.Session
	User    func() Collection
	Project func() Collection
}

func (db DB) Close() {
	db.session.Close()
}

type Collection interface {
	Bulk() *mgo.Bulk
	Find(query bson.M) *mgo.Query
	FindId(id bson.ObjectId) *mgo.Query
	Insert(docs ...interface{}) error
	Pipe(pipeline interface{}) *mgo.Pipe
	Remove(selector bson.M) error
	RemoveAll(selector bson.M) (info *mgo.ChangeInfo, err error)
	RemoveId(id bson.ObjectId) error
	Update(selector bson.M, update interface{}) error
	UpdateAll(selector bson.M, update interface{}) (info *mgo.ChangeInfo, err error)
	UpdateId(id bson.ObjectId, update interface{}) error
	Upsert(selector bson.M, update interface{}) (info *mgo.ChangeInfo, err error)
	UpsertId(id bson.ObjectId, update interface{}) (info *mgo.ChangeInfo, err error)
}

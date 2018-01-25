package db

import (
	conf "github.com/wkozyra95/go-web-starter/config"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var log = conf.NamedLogger("db")

const OR = "$or"

const projectCollection = "project"
const userCollection = "user"

const PrimaryKey = "_id"
const ProjectForeignKey = "projectId"
const UserForeignKey = "userId"
const UserIdKeyUsername = "username"
const UserIdKeyEmail = "email"

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

func SetupDB(config conf.Config) (func() DB, error) {
	log.Info("Connecting to db ...")
	session, sessionErr := mgo.Dial(config.DbURL)
	if sessionErr != nil {
		log.Infof("Connection error: %s", sessionErr.Error())
		return nil, sessionErr
	}
	log.Info("Connected")

	log.Info("Ensure indices")
	ensureErr := ensureDBIndices(session.DB(""))
	if ensureErr != nil {
		log.Infof("Ensure indices  error: %s", ensureErr.Error())
		return nil, ensureErr
	}
	log.Info("Ensure success")

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

func ensureDBIndices(db *mgo.Database) error {
	ensureErrs := []error{
		db.C(userCollection).EnsureIndex(mgo.Index{
			Key: []string{PrimaryKey},
		}),
		db.C(userCollection).EnsureIndex(mgo.Index{
			Key:    []string{UserIdKeyUsername},
			Unique: true,
		}),
		db.C(projectCollection).EnsureIndex(mgo.Index{
			Key: []string{PrimaryKey},
		}),
		db.C(projectCollection).EnsureIndex(mgo.Index{
			Key: []string{UserForeignKey},
		}),
	}
	for _, err := range ensureErrs {
		if err != nil {
			return err
		}
	}
	return nil
}

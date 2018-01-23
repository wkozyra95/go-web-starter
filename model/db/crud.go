package db

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const projectCollection = "project"
const userCollection = "user"

const PrimaryKey = "_id"
const ProjectForeignKey = "projectId"
const UserForeignKey = "userId"

type Document struct {
	Id        bson.ObjectId `bson:"_id"`
	ProjectId bson.ObjectId `bson:"projectId"`
	UserId    bson.ObjectId `bson:"userId"`
}

type collection struct {
	collection *mgo.Collection
}

func (c collection) Bulk() *mgo.Bulk {
	return c.collection.Bulk()
}
func (c collection) Find(query bson.M) *mgo.Query {
	return c.collection.Find(query)
}
func (c collection) FindId(id bson.ObjectId) *mgo.Query {
	return c.collection.FindId(id)
}
func (c collection) Insert(docs ...interface{}) error {
	return c.collection.Insert(docs...)
}
func (c collection) Pipe(pipeline interface{}) *mgo.Pipe {
	return c.collection.Pipe(pipeline)
}
func (c collection) Remove(selector bson.M) error {
	return c.collection.Remove(selector)
}
func (c collection) RemoveAll(selector bson.M) (info *mgo.ChangeInfo, err error) {
	return c.collection.RemoveAll(selector)
}
func (c collection) RemoveId(id bson.ObjectId) error {
	return c.collection.RemoveId(id)
}
func (c collection) Update(selector bson.M, update interface{}) error {
	return c.collection.Update(selector, update)
}
func (c collection) UpdateAll(selector bson.M, update interface{}) (info *mgo.ChangeInfo, err error) {
	return c.collection.UpdateAll(selector, update)
}
func (c collection) UpdateId(id bson.ObjectId, update interface{}) error {
	return c.collection.UpdateId(id, update)
}
func (c collection) Upsert(selector bson.M, update interface{}) (info *mgo.ChangeInfo, err error) {
	return c.collection.Upsert(selector, update)
}
func (c collection) UpsertId(id bson.ObjectId, update interface{}) (info *mgo.ChangeInfo, err error) {
	return c.collection.UpsertId(id, update)
}

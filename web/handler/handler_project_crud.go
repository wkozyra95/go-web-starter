package handler

import (
	"fmt"
	"github.com/wkozyra95/go-web-starter/model"
	"github.com/wkozyra95/go-web-starter/model/db"
	"gopkg.in/mgo.v2/bson"
)

// ProjectCreate ...
func ProjectGet(id bson.ObjectId, context ActionContext) (model.Project, error) {
	accessOk, accessErr := validateReadRights(id, context.DB.Project(), context)
	if accessErr != nil {
		return model.Project{}, accessErr
	}
	if !accessOk {
		return model.Project{}, fmt.Errorf("")
	}
	project := model.Project{}
	getErr := context.DB.Project().FindId(id).One(&project)
	if getErr != nil {
		return model.Project{}, getErr
	}

	return project, nil
}

func ProjectGetAll(context ActionContext) ([]model.Project, error) {
	projects := []model.Project{}
	getErr := context.DB.Project().Find(bson.M{db.UserForeignKey: context.UserId}).All(&projects)
	if getErr != nil {
		return nil, getErr
	}
	return projects, nil
}

func ProjectCreate(project model.Project, context ActionContext) (bson.ObjectId, error) {
	project.Id = bson.NewObjectId()
	project.UserId = context.UserId
	err := context.DB.Project().Insert(project)
	if err != nil {
		return "", err
	}
	return project.Id, nil
}

func ProjectUpdate(project model.Project, context ActionContext) error {
	accessOk, accessErr := validateWriteRights(project.Id, context.DB.Project(), context)
	if accessErr != nil {
		return accessErr
	}
	if !accessOk {
		return fmt.Errorf("")
	}

	updateErr := context.DB.Project().UpdateId(project.Id, project)
	if updateErr != nil {
		return updateErr
	}
	return nil
}

func ProjectDelete(id bson.ObjectId, context ActionContext) error {
	accessOk, accessErr := validateWriteRights(id, context.DB.Project(), context)
	if accessErr != nil {
		return accessErr
	}
	if !accessOk {
		return fmt.Errorf("")
	}
	err := context.DB.Project().RemoveId(id)
	if err != nil {
		return err
	}
	return nil
}

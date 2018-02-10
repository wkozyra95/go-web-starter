package web

import (
	"net/http"

	"github.com/wkozyra95/go-web-starter/errors"
	"github.com/wkozyra95/go-web-starter/model"
	"github.com/wkozyra95/go-web-starter/model/mongo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (h *handler) getProjectsHandler(w http.ResponseWriter, r *http.Request) {
	userId := extractUserId(r.Context())
	db := extractDBSession(r.Context())

	projects := []mongo.Project{}
	projectsErr := db.Project().Find(bson.M{
		mongo.UserForeignKey: userId,
	}).All(&projects)
	if projectsErr == mgo.ErrNotFound {
		handleRequestErr(w, errors.NotFound)
		return
	}
	if projectsErr != nil {
		handleRequestErr(w, errors.InternalServerError)
		return
	}

	_ = writeJSONResponse(w, http.StatusOK, projects)
}

func (h *handler) getProjectHandler(w http.ResponseWriter, r *http.Request) {
	userId := extractUserId(r.Context())
	db := extractDBSession(r.Context())
	projectId := extractProjectId(r.Context())

	project := mongo.Project{}
	projectErr := db.Project().FindID(projectId).One(&project)
	if projectErr == mgo.ErrNotFound {
		handleRequestErr(w, errors.NotFound)
		return
	}
	if projectErr != nil {
		handleRequestErr(w, errors.InternalServerError)
		return
	}
	if project.UserID != userId {
		handleRequestErr(w, errors.Unauthorized)
		return
	}

	_ = writeJSONResponse(w, http.StatusOK, project)
}

func (h *handler) createProjectHandler(w http.ResponseWriter, r *http.Request) {
	userId := extractUserId(r.Context())
	db := extractDBSession(r.Context())

	projectInput := projectCreateInput(model.Project{})
	decodeErr := decodeJSONRequest(r, &projectInput)
	if decodeErr != nil {
		handleRequestErr(w, errors.Malformed)
		return
	}

	if err := projectInput.validate(); err != nil {
		handleRequestErr(w, err)
		return
	}

	project := projectInput.createProject(userId)
	insertErr := db.Project().Insert(project)
	if insertErr != nil {
		handleRequestErr(w, errors.InternalServerError)
		return
	}

	_ = writeJSONResponse(w, http.StatusCreated, struct {
		ID bson.ObjectId `json:"id"`
		model.Project
	}{
		ID:      project.ID,
		Project: project.Project,
	})
}

type projectCreateInput model.Project

func (pc projectCreateInput) validate() error {
	return nil
}

func (pc projectCreateInput) createProject(userId bson.ObjectId) mongo.Project {
	return mongo.Project{
		Project: model.Project(pc),
		ID:      bson.NewObjectId(),
		UserID:  userId,
	}
}

func (h *handler) updateProjectHandler(w http.ResponseWriter, r *http.Request) {

	_ = writeJSONResponse(w, http.StatusOK, []byte{})
}

type projectUpdateInput struct {
	Set struct {
		model.Project
	} `json:"set"`
	Delete []string `json:"delete"`
}

func (h *handler) deleteProjectHandler(w http.ResponseWriter, r *http.Request) {

	_ = writeJSONResponse(w, http.StatusOK, []byte{})
}

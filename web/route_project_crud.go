package web

import (
	"net/http"

	"github.com/wkozyra95/go-web-starter/model"
	"github.com/wkozyra95/go-web-starter/web/handler"
	"gopkg.in/mgo.v2/bson"
)

func getProjectsHandler(w http.ResponseWriter, r *http.Request, ctx requestCtx) error {
	userId := helperExtractUserId(r)
	context := handler.ActionContext{
		DB:     ctx.server.db(),
		UserId: userId,
	}

	projects, projectsErr := handler.ProjectGetAll(context)
	if projectsErr != nil {
		return projectsErr
	}

	_ = writeJSONResponse(w, http.StatusFound, projects)
	return nil
}

func getProjectHandler(w http.ResponseWriter, r *http.Request, ctx requestCtx) error {
	userId := helperExtractUserId(r)
	context := handler.ActionContext{
		DB:     ctx.server.db(),
		UserId: userId,
	}
	projectIdBin := ctx.chi.URLParam(paramProjectId)
	projectId := bson.ObjectIdHex(projectIdBin)

	project, projectErr := handler.ProjectGet(projectId, context)
	if projectErr != nil {
		return projectErr
	}

	_ = writeJSONResponse(w, http.StatusFound, project)
	return nil
}

func createProjectHandler(w http.ResponseWriter, r *http.Request, ctx requestCtx) error {
	userId := helperExtractUserId(r)
	context := handler.ActionContext{
		DB:     ctx.server.db(),
		UserId: userId,
	}
	var project model.Project
	decodeErr := decodeJSONRequest(r, &project)
	if decodeErr != nil {
		return decodeErr
	}

	projectId, projectErr := handler.ProjectCreate(project, context)
	if projectErr != nil {
		return projectErr
	}

	_ = writeJSONResponse(w, http.StatusFound, struct {
		ProjectId bson.ObjectId `json:"projectId"`
	}{
		ProjectId: projectId,
	})
	return nil
}

func updateProjectHandler(w http.ResponseWriter, r *http.Request, ctx requestCtx) error {
	userId := helperExtractUserId(r)
	context := handler.ActionContext{
		DB:     ctx.server.db(),
		UserId: userId,
	}
	projectIdBin := ctx.chi.URLParam(paramProjectId)
	projectId := bson.ObjectIdHex(projectIdBin)

	var project model.Project
	decodeErr := decodeJSONRequest(r, &project)
	if decodeErr != nil {
		return decodeErr
	}
	project.Id = projectId

	projectErr := handler.ProjectUpdate(project, context)
	if projectErr != nil {
		return projectErr
	}

	_ = writeJSONResponse(w, http.StatusOK, "")
	return nil
}

func deleteProjectHandler(w http.ResponseWriter, r *http.Request, ctx requestCtx) error {
	userId := helperExtractUserId(r)
	context := handler.ActionContext{
		DB:     ctx.server.db(),
		UserId: userId,
	}
	projectIdBin := ctx.chi.URLParam(paramProjectId)
	projectId := bson.ObjectIdHex(projectIdBin)

	projectErr := handler.ProjectDelete(projectId, context)
	if projectErr != nil {
		return projectErr
	}

	_ = writeJSONResponse(w, http.StatusOK, []byte{})
	return nil
}

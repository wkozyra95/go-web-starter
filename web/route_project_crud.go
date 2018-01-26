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
		DB:     ctx.db,
		UserId: userId,
	}

	projects, projectsErr := handler.ProjectGetAll(context)
	if projectsErr != nil {
		log.Warnf("request GET projects error [%s]", projectsErr.Error())
		return projectsErr
	}

	_ = writeJSONResponse(w, http.StatusFound, projects)
	return nil
}

func getProjectHandler(w http.ResponseWriter, r *http.Request, ctx requestCtx) error {
	userId := helperExtractUserId(r)
	context := handler.ActionContext{
		DB:     ctx.db,
		UserId: userId,
	}
	projectIdStr := ctx.chi.URLParam(paramProjectId)
	projectId := bson.ObjectIdHex(projectIdStr)

	project, projectErr := handler.ProjectGet(projectId, context)
	if projectErr != nil {
		log.Warnf("request GET project id(%s) error [%s]", projectIdStr, projectErr.Error())
		return projectErr
	}

	_ = writeJSONResponse(w, http.StatusFound, project)
	return nil
}

func createProjectHandler(w http.ResponseWriter, r *http.Request, ctx requestCtx) error {
	userId := helperExtractUserId(r)
	context := handler.ActionContext{
		DB:     ctx.db,
		UserId: userId,
	}
	var project model.Project
	decodeErr := decodeJSONRequest(r, &project)
	if decodeErr != nil {
		return requestMalformedErr("request CREATE project malformed")
	}

	projectId, projectErr := handler.ProjectCreate(project, context)
	if projectErr != nil {
		log.Warnf("request CREATE project error [%s]", projectErr.Error())
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
		DB:     ctx.db,
		UserId: userId,
	}
	projectIdStr := ctx.chi.URLParam(paramProjectId)
	projectId := bson.ObjectIdHex(projectIdStr)

	var project model.Project
	decodeErr := decodeJSONRequest(r, &project)
	if decodeErr != nil {
		return requestMalformedErr("request UPDATE project malformed")
	}
	project.Id = projectId

	projectErr := handler.ProjectUpdate(project, context)
	if projectErr != nil {
		log.Warnf("request UPDATE project id(%s) error [%s]", projectIdStr, projectErr.Error())
		return projectErr
	}

	_ = writeJSONResponse(w, http.StatusOK, []byte{})
	return nil
}

func deleteProjectHandler(w http.ResponseWriter, r *http.Request, ctx requestCtx) error {
	userId := helperExtractUserId(r)
	context := handler.ActionContext{
		DB:     ctx.db,
		UserId: userId,
	}
	projectIdStr := ctx.chi.URLParam(paramProjectId)
	projectId := bson.ObjectIdHex(projectIdStr)

	projectErr := handler.ProjectDelete(projectId, context)
	if projectErr != nil {
		log.Warnf("request DELETE project id(%s) error [%s]", projectIdStr, projectErr.Error())
		return projectErr
	}

	_ = writeJSONResponse(w, http.StatusOK, []byte{})
	return nil
}

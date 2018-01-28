package web

import (
	"net/http"

	"github.com/wkozyra95/go-web-starter/model"
	"github.com/wkozyra95/go-web-starter/web/handler"
	"gopkg.in/mgo.v2/bson"
)

func getProjectsHandler(w http.ResponseWriter, r *http.Request, ctx requestCtx) error {
	userID := helperExtractUserID(r)
	context := handler.ActionContext{
		DB:     ctx.db,
		UserID: userID,
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
	userID := helperExtractUserID(r)
	context := handler.ActionContext{
		DB:     ctx.db,
		UserID: userID,
	}
	projectIDStr := ctx.chi.URLParam(paramProjectID)
	projectID := bson.ObjectIdHex(projectIDStr)

	project, projectErr := handler.ProjectGet(projectID, context)
	if projectErr != nil {
		log.Warnf("request GET project id(%s) error [%s]", projectIDStr, projectErr.Error())
		return projectErr
	}

	_ = writeJSONResponse(w, http.StatusFound, project)
	return nil
}

func createProjectHandler(w http.ResponseWriter, r *http.Request, ctx requestCtx) error {
	userID := helperExtractUserID(r)
	context := handler.ActionContext{
		DB:     ctx.db,
		UserID: userID,
	}
	var project model.Project
	decodeErr := decodeJSONRequest(r, &project)
	if decodeErr != nil {
		return requestMalformedErr("request CREATE project malformed")
	}

	projectID, projectErr := handler.ProjectCreate(project, context)
	if projectErr != nil {
		log.Warnf("request CREATE project error [%s]", projectErr.Error())
		return projectErr
	}

	_ = writeJSONResponse(w, http.StatusFound, struct {
		ProjectID bson.ObjectId `json:"projectId"`
	}{
		ProjectID: projectID,
	})
	return nil
}

func updateProjectHandler(w http.ResponseWriter, r *http.Request, ctx requestCtx) error {
	userID := helperExtractUserID(r)
	context := handler.ActionContext{
		DB:     ctx.db,
		UserID: userID,
	}
	projectIDStr := ctx.chi.URLParam(paramProjectID)
	projectID := bson.ObjectIdHex(projectIDStr)

	var project model.Project
	decodeErr := decodeJSONRequest(r, &project)
	if decodeErr != nil {
		return requestMalformedErr("request UPDATE project malformed")
	}
	project.ID = projectID

	projectErr := handler.ProjectUpdate(project, context)
	if projectErr != nil {
		log.Warnf("request UPDATE project id(%s) error [%s]", projectIDStr, projectErr.Error())
		return projectErr
	}

	_ = writeJSONResponse(w, http.StatusOK, []byte{})
	return nil
}

func deleteProjectHandler(w http.ResponseWriter, r *http.Request, ctx requestCtx) error {
	userID := helperExtractUserID(r)
	context := handler.ActionContext{
		DB:     ctx.db,
		UserID: userID,
	}
	projectIDStr := ctx.chi.URLParam(paramProjectID)
	projectID := bson.ObjectIdHex(projectIDStr)

	projectErr := handler.ProjectDelete(projectID, context)
	if projectErr != nil {
		log.Warnf("request DELETE project id(%s) error [%s]", projectIDStr, projectErr.Error())
		return projectErr
	}

	_ = writeJSONResponse(w, http.StatusOK, []byte{})
	return nil
}

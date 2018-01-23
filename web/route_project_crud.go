package web

import (
	"github.com/wkozyra95/go-web-starter/model"
	"github.com/wkozyra95/go-web-starter/web/handler"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

func getProjectsHandler(w http.ResponseWriter, r *http.Request, ctx requestCtx) {
	userId := helperExtractUserId(r)
	context := handler.ActionContext{
		DB:     ctx.server.db(),
		UserId: userId,
	}

	projects, projectsErr := handler.ProjectGetAll(context)
	if projectsErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
	}

	_ = writeJSONResponse(w, http.StatusFound, projects)
}

func getProjectHandler(w http.ResponseWriter, r *http.Request, ctx requestCtx) {
	userId := helperExtractUserId(r)
	context := handler.ActionContext{
		DB:     ctx.server.db(),
		UserId: userId,
	}
	projectIdBin := ctx.chi.URLParam(paramProjectId)
	projectId := bson.ObjectIdHex(projectIdBin)

	project, projectErr := handler.ProjectGet(projectId, context)
	if projectErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
	}

	_ = writeJSONResponse(w, http.StatusFound, project)
}

func createProjectHandler(w http.ResponseWriter, r *http.Request, ctx requestCtx) {
	userId := helperExtractUserId(r)
	context := handler.ActionContext{
		DB:     ctx.server.db(),
		UserId: userId,
	}
	var project model.Project
	decodeErr := decodeJSONRequest(r, &project)
	if decodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
	}

	projectId, projectErr := handler.ProjectCreate(project, context)
	if projectErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
	}

	_ = writeJSONResponse(w, http.StatusFound, struct {
		ProjectId bson.ObjectId `json:"projectId"`
	}{
		ProjectId: projectId,
	})
}

func updateProjectHandler(w http.ResponseWriter, r *http.Request, ctx requestCtx) {
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
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
	}
	project.Id = projectId

	projectErr := handler.ProjectUpdate(project, context)
	if projectErr != nil {
		w.Write([]byte("Internal server error"))
		w.WriteHeader(http.StatusInternalServerError)
	}

	_ = writeJSONResponse(w, http.StatusOK, "")
}

func deleteProjectHandler(w http.ResponseWriter, r *http.Request, ctx requestCtx) {
	userId := helperExtractUserId(r)
	context := handler.ActionContext{
		DB:     ctx.server.db(),
		UserId: userId,
	}
	projectIdBin := ctx.chi.URLParam(paramProjectId)
	projectId := bson.ObjectIdHex(projectIdBin)

	projectErr := handler.ProjectDelete(projectId, context)
	if projectErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
	}

	_ = writeJSONResponse(w, http.StatusOK, "")
}

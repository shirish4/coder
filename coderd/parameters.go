package coderd

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"

	"github.com/coder/coder/database"
	"github.com/coder/coder/httpapi"
)

type ParameterScope string

const (
	ParameterOrganization ParameterScope = "organization"
	ParameterProject      ParameterScope = "project"
	ParameterUser         ParameterScope = "user"
	ParameterWorkspace    ParameterScope = "workspace"
)

// Parameter represents a set value for the scope.
type Parameter struct {
	ID                uuid.UUID                           `db:"id" json:"id"`
	CreatedAt         time.Time                           `db:"created_at" json:"created_at"`
	UpdatedAt         time.Time                           `db:"updated_at" json:"updated_at"`
	Scope             ParameterScope                      `db:"scope" json:"scope"`
	ScopeID           string                              `db:"scope_id" json:"scope_id"`
	Name              string                              `db:"name" json:"name"`
	SourceScheme      database.ParameterSourceScheme      `db:"source_scheme" json:"source_scheme"`
	DestinationScheme database.ParameterDestinationScheme `db:"destination_scheme" json:"destination_scheme"`
}

// CreateParameterRequest is used to create a new parameter value for a scope.
type CreateParameterRequest struct {
	Name              string                              `json:"name" validate:"required"`
	SourceValue       string                              `json:"source_value" validate:"required"`
	SourceScheme      database.ParameterSourceScheme      `json:"source_scheme" validate:"oneof=data,required"`
	DestinationScheme database.ParameterDestinationScheme `json:"destination_scheme" validate:"oneof=environment_variable provisioner_variable,required"`
}

func (api *api) postParameter(rw http.ResponseWriter, r *http.Request) {
	var createRequest CreateParameterRequest
	if !httpapi.Read(rw, r, &createRequest) {
		return
	}
	scope, scopeID, valid := readScopeAndID(rw, r)
	if !valid {
		return
	}
	_, err := api.Database.GetParameterValueByScopeAndName(r.Context(), database.GetParameterValueByScopeAndNameParams{
		Scope:   scope,
		ScopeID: scopeID,
		Name:    createRequest.Name,
	})
	if err == nil {
		httpapi.Write(rw, http.StatusConflict, httpapi.Response{
			Message: fmt.Sprintf("a parameter already exists in scope %q with name %q", scope, createRequest.Name),
		})
		return
	}
	if !errors.Is(err, sql.ErrNoRows) {
		httpapi.Write(rw, http.StatusInternalServerError, httpapi.Response{
			Message: fmt.Sprintf("get parameter value: %s", err),
		})
		return
	}
	parameterValue, err := api.Database.InsertParameterValue(r.Context(), database.InsertParameterValueParams{
		ID:                uuid.New(),
		Name:              createRequest.Name,
		CreatedAt:         database.Now(),
		UpdatedAt:         database.Now(),
		Scope:             scope,
		ScopeID:           scopeID,
		SourceScheme:      createRequest.SourceScheme,
		SourceValue:       createRequest.SourceValue,
		DestinationScheme: createRequest.DestinationScheme,
	})
	if err != nil {
		httpapi.Write(rw, http.StatusInternalServerError, httpapi.Response{
			Message: fmt.Sprintf("insert parameter value: %s", err),
		})
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(rw, r, convertParameterValue(parameterValue))
}

func (api *api) parameters(rw http.ResponseWriter, r *http.Request) {
	scope, scopeID, valid := readScopeAndID(rw, r)
	if !valid {
		return
	}
	parameterValues, err := api.Database.GetParameterValuesByScope(r.Context(), database.GetParameterValuesByScopeParams{
		Scope:   scope,
		ScopeID: scopeID,
	})
	if errors.Is(err, sql.ErrNoRows) {
		err = nil
	}
	if err != nil {
		httpapi.Write(rw, http.StatusInternalServerError, httpapi.Response{
			Message: fmt.Sprintf("get parameter values by scope: %s", err),
		})
		return
	}
	apiParameterValues := make([]Parameter, 0, len(parameterValues))
	for _, parameterValue := range parameterValues {
		apiParameterValues = append(apiParameterValues, convertParameterValue(parameterValue))
	}

	render.Status(r, http.StatusOK)
	render.JSON(rw, r, apiParameterValues)
}

func (api *api) deleteParameter(rw http.ResponseWriter, r *http.Request) {
	scope, scopeID, valid := readScopeAndID(rw, r)
	if !valid {
		return
	}
	name := chi.URLParam(r, "name")
	parameterValue, err := api.Database.GetParameterValueByScopeAndName(r.Context(), database.GetParameterValueByScopeAndNameParams{
		Scope:   scope,
		ScopeID: scopeID,
		Name:    name,
	})
	if errors.Is(err, sql.ErrNoRows) {
		httpapi.Write(rw, http.StatusNotFound, httpapi.Response{
			Message: fmt.Sprintf("parameter doesn't exist in the provided scope with name %q", name),
		})
		return
	}
	if err != nil {
		httpapi.Write(rw, http.StatusInternalServerError, httpapi.Response{
			Message: fmt.Sprintf("get parameter value: %s", err),
		})
		return
	}
	err = api.Database.DeleteParameterValueByID(r.Context(), parameterValue.ID)
	if err != nil {
		httpapi.Write(rw, http.StatusInternalServerError, httpapi.Response{
			Message: fmt.Sprintf("delete parameter: %s", err),
		})
		return
	}
	httpapi.Write(rw, http.StatusOK, httpapi.Response{
		Message: "parameter deleted",
	})
}

func convertParameterValue(parameterValue database.ParameterValue) Parameter {
	return Parameter{
		ID:                parameterValue.ID,
		CreatedAt:         parameterValue.CreatedAt,
		UpdatedAt:         parameterValue.UpdatedAt,
		Scope:             ParameterScope(parameterValue.Scope),
		ScopeID:           parameterValue.ScopeID,
		Name:              parameterValue.Name,
		SourceScheme:      parameterValue.SourceScheme,
		DestinationScheme: parameterValue.DestinationScheme,
	}
}

func readScopeAndID(rw http.ResponseWriter, r *http.Request) (database.ParameterScope, string, bool) {
	var scope database.ParameterScope
	switch chi.URLParam(r, "scope") {
	case string(ParameterOrganization):
		scope = database.ParameterScopeOrganization
	case string(ParameterProject):
		scope = database.ParameterScopeProject
	case string(ParameterUser):
		scope = database.ParameterScopeUser
	case string(ParameterWorkspace):
		scope = database.ParameterScopeWorkspace
	default:
		httpapi.Write(rw, http.StatusBadRequest, httpapi.Response{
			Message: fmt.Sprintf("invalid scope %q", scope),
		})
		return scope, "", false
	}

	return scope, chi.URLParam(r, "id"), true
}
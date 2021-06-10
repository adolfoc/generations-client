package moment_types

import (
	"encoding/json"
	"fmt"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

const (
	ResourceMomentType = "moment-types"
)

func getMomentTypesURL() string {
	return fmt.Sprintf("%s%s", handlers.GetAPIHostURL(), ResourceMomentType)
}

func getSchemaMomentTypeURL(momentSchemaID, momentID int) string {
	return fmt.Sprintf("%s%s/%d", handlers.GetAPIHostURL(), ResourceMomentType, momentID)
}

func getSchemaMomentType(w http.ResponseWriter, r *http.Request, momentSchemaID, momentTypeID int) (*model.MomentType, error) {
	url := getSchemaMomentTypeURL(momentSchemaID, momentTypeID)
	code, body, err := handlers.GetResource(w, r, url)

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	var momentType *model.MomentType
	err = json.Unmarshal(body, &momentType)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return momentType, nil
}

func getUrlGenerationSchemaID(w http.ResponseWriter, r *http.Request) (int, error) {
	return handlers.GetUrlIntParam("schema_id", w, r)
}

func getUrlMomentTypeID(w http.ResponseWriter, r *http.Request) (int, error) {
	return handlers.GetUrlIntParam("moment_type_id", w, r)
}

func buildMomentTypeRequest(momentType *model.MomentType) *model.MomentTypeRequest {
	mtr := &model.MomentTypeRequest{
		ID:          momentType.ID,
		SchemaID:    momentType.SchemaID,
		Name:        momentType.Name,
		Description: momentType.Description,
	}

	return mtr
}

func newMomentTypeRequest(schemaID int) *model.MomentTypeRequest {
	gr := &model.MomentTypeRequest{
		SchemaID: schemaID,
	}

	return gr
}

func makeMomentTypeRequest(r *http.Request) (*model.MomentTypeRequest, error) {
	normalizedID := handlers.GetIntFormValue(r, "inputID")
	normalizedSchemaID := handlers.GetIntFormValue(r, "inputSchemaID")
	name := handlers.GetStringFormValue(r, "inputName")
	description := handlers.GetStringFormValue(r, "inputDescription")

	mtr := &model.MomentTypeRequest{
		ID:          normalizedID,
		SchemaID:    normalizedSchemaID,
		Name:        name,
		Description: description,
	}

	return mtr, nil
}

func postMomentType(w http.ResponseWriter, r *http.Request, momentTypeRequest *model.MomentTypeRequest) (int, []byte, error) {
	url := getMomentTypesURL()

	payload, err := json.Marshal(momentTypeRequest)
	if err != nil {
		return 0, nil, err
	}

	if err != nil {
		return 0, nil, err
	}

	code, body, err := handlers.PostResource(w, r, url, payload)

	return code, body, nil
}

func patchMomentType(w http.ResponseWriter, r *http.Request, momentTypeRequest *model.MomentTypeRequest) (int, []byte, error) {
	url := getSchemaMomentTypeURL(momentTypeRequest.SchemaID, momentTypeRequest.ID)

	payload, err := json.Marshal(momentTypeRequest)
	if err != nil {
		return 0, nil, err
	}

	code, body, err := handlers.PatchResource(w, r, url, payload)
	if err != nil {
		return 0, nil, err
	}


	return code, body, nil
}

func deleteMomentType(w http.ResponseWriter, r *http.Request, schemaID, momentTypeID int) (int, []byte, error) {
	url := getSchemaMomentTypeURL(schemaID, momentTypeID)

	code, body, err := handlers.DeleteResource(w, r, url)
	if err != nil {
		return 0, nil, err
	}


	return code, body, nil
}

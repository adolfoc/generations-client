package generation_types

import (
	"encoding/json"
	"fmt"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

const (
	ResourceGenerationType = "generation-types"
)

func getGenerationTypesURL() string {
	return fmt.Sprintf("%s%s", handlers.GetAPIHostURL(), ResourceGenerationType)
}

func getSchemaGenerationTypeURL(generationSchemaID, generationID int) string {
	return fmt.Sprintf("%s%s/%d", handlers.GetAPIHostURL(), ResourceGenerationType, generationID)
}

func getSchemaGenerationType(w http.ResponseWriter, r *http.Request, generationSchemaID, generationTypeID int) (*model.GenerationType, error) {
	url := getSchemaGenerationTypeURL(generationSchemaID, generationTypeID)
	code, body, err := handlers.GetResource(w, r, url)

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	var generationType *model.GenerationType
	err = json.Unmarshal(body, &generationType)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return generationType, nil
}

func getUrlGenerationSchemaID(w http.ResponseWriter, r *http.Request) (int, error) {
	return handlers.GetUrlIntParam("schema_id", w, r)
}

func getUrlGenerationTypeID(w http.ResponseWriter, r *http.Request) (int, error) {
	return handlers.GetUrlIntParam("generation_type_id", w, r)
}

func buildGenerationTypeRequest(generationType *model.GenerationType) *model.GenerationTypeRequest {
	gr := &model.GenerationTypeRequest{
		ID:          generationType.ID,
		SchemaID:    generationType.SchemaID,
		Archetype:   generationType.Archetype,
		Description: generationType.Description,
	}

	return gr
}

func newGenerationTypeRequest(schemaID int) *model.GenerationTypeRequest {
	gr := &model.GenerationTypeRequest{
		SchemaID:             schemaID,
	}

	return gr
}

func makeGenerationTypeRequest(r *http.Request) (*model.GenerationTypeRequest, error) {
	normalizedID := handlers.GetIntFormValue(r, "inputID")
	normalizedSchemaID := handlers.GetIntFormValue(r, "inputSchemaID")
	archetype := handlers.GetStringFormValue(r, "inputArchetype")
	description := handlers.GetStringFormValue(r, "inputDescription")

	gtr := &model.GenerationTypeRequest{
		ID:          normalizedID,
		SchemaID:    normalizedSchemaID,
		Archetype:   archetype,
		Description: description,
	}

	return gtr, nil
}

func postGenerationType(w http.ResponseWriter, r *http.Request, generationTypeRequest *model.GenerationTypeRequest) (int, []byte, error) {
	url := getGenerationTypesURL()

	payload, err := json.Marshal(generationTypeRequest)
	if err != nil {
		return 0, nil, err
	}

	if err != nil {
		return 0, nil, err
	}

	code, body, err := handlers.PostResource(w, r, url, payload)

	return code, body, nil
}

func patchGenerationType(w http.ResponseWriter, r *http.Request, generationTypeRequest *model.GenerationTypeRequest) (int, []byte, error) {
	url := getSchemaGenerationTypeURL(generationTypeRequest.SchemaID, generationTypeRequest.ID)

	payload, err := json.Marshal(generationTypeRequest)
	if err != nil {
		return 0, nil, err
	}

	code, body, err := handlers.PatchResource(w, r, url, payload)
	if err != nil {
		return 0, nil, err
	}


	return code, body, nil
}

func deleteGenerationType(w http.ResponseWriter, r *http.Request, schemaID, generationTypeID int) (int, []byte, error) {
	url := getSchemaGenerationTypeURL(schemaID, generationTypeID)

	code, body, err := handlers.DeleteResource(w, r, url)
	if err != nil {
		return 0, nil, err
	}


	return code, body, nil
}

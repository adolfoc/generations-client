package generational_landscape

import (
	"encoding/json"
	"fmt"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

const (
	ResourceGenerationalLandscape = "generational-landscapes"
	ResourceSchemas               = "generation-schemas"
)

func getGenerationalLandscapesURL(generationSchemaID int) string {
	return fmt.Sprintf("%s%s", handlers.GetAPIHostURL(), ResourceGenerationalLandscape)
}

func getGenerationalLandscapeURL(generationSchemaID, generationalLandscapeID int) string {
	return fmt.Sprintf("%s%s/%d", handlers.GetAPIHostURL(), ResourceGenerationalLandscape, generationalLandscapeID)
}

func getMomentsForSchemaURL(schemaID int) string {
	return fmt.Sprintf("%s%s/%d/moments", handlers.GetAPIHostURL(), ResourceSchemas, schemaID)
}

func getGenerationalLandscape(w http.ResponseWriter, r *http.Request, generationSchemaID, generationalLandscapeID int) (*model.GenerationalLandscape, error) {
	url := getGenerationalLandscapeURL(generationSchemaID, generationalLandscapeID)
	code, body, err := handlers.GetResource(w, r, url)

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	var generationalLandscape *model.GenerationalLandscape
	err = json.Unmarshal(body, &generationalLandscape)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return generationalLandscape, nil
}

func getUrlGenerationSchemaID(w http.ResponseWriter, r *http.Request) (int, error) {
	return handlers.GetUrlIntParam("schema_id", w, r)
}

func getUrlGenerationalLandscapeID(w http.ResponseWriter, r *http.Request) (int, error) {
	return handlers.GetUrlIntParam("generational_landscape_id", w, r)
}

func getUrlGenerationID(w http.ResponseWriter, r *http.Request) (int, error) {
	return handlers.GetUrlIntParam("generation_id", w, r)
}

func getMomentsForSchema(w http.ResponseWriter, r *http.Request, schemaID int) (*model.HistoricalMoments, error) {
	url := getMomentsForSchemaURL(schemaID)
	code, body, err := handlers.GetResource(w, r, url)

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	var moments *model.HistoricalMoments
	err = json.Unmarshal(body, &moments)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return moments, nil
}

func buildGenerationalLandscapeRequest(generationalLandscape *model.GenerationalLandscape) *model.GenerationalLandscapeRequest {
	glr := &model.GenerationalLandscapeRequest{
		ID:                generationalLandscape.ID,
		GenerationID:      generationalLandscape.GenerationID,
		FormationMomentID: generationalLandscape.FormationMomentID,
		Description:       generationalLandscape.Description,
	}

	return glr
}

func newGenerationalLandscapeRequest(generationID int) *model.GenerationalLandscapeRequest {
	glr := &model.GenerationalLandscapeRequest{
		GenerationID:      generationID,
	}

	return glr
}

func makeGenerationalLandscapeRequest(r *http.Request) (*model.GenerationalLandscapeRequest, error) {
	normalizedID := handlers.GetIntFormValue(r, "inputID")
	normalizedGenerationID := handlers.GetIntFormValue(r, "inputGenerationID")
	normalizedFormationMomentID := handlers.GetIntFormValue(r, "inputFormationMomentID")
	description := handlers.GetStringFormValue(r, "inputDescription")

	glr := &model.GenerationalLandscapeRequest{
		ID:                normalizedID,
		GenerationID:      normalizedGenerationID,
		FormationMomentID: normalizedFormationMomentID,
		Description:       description,
	}

	return glr, nil
}

func postGenerationalLandscape(w http.ResponseWriter, r *http.Request, glRequest *model.GenerationalLandscapeRequest) (int, []byte, error) {
	url := getGenerationalLandscapesURL(0)

	payload, err := json.Marshal(glRequest)
	if err != nil {
		return 0, nil, err
	}

	if err != nil {
		return 0, nil, err
	}

	code, body, err := handlers.PostResource(w, r, url, payload)

	return code, body, nil
}

func patchGenerationalLandscape(w http.ResponseWriter, r *http.Request, auRequest *model.GenerationalLandscapeRequest) (int, []byte, error) {
	url := getGenerationalLandscapeURL(0, auRequest.ID)

	payload, err := json.Marshal(auRequest)
	if err != nil {
		return 0, nil, err
	}

	code, body, err := handlers.PatchResource(w, r, url, payload)
	if err != nil {
		return 0, nil, err
	}


	return code, body, nil
}

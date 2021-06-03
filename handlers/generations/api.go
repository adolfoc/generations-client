package generations

import (
	"encoding/json"
	"fmt"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

const (
	ResourceGenerationSchema = "generation-schemas"
	ResourceGenerations = "generations"
	ResourceMoments = "historical-moments"
)

func getSchemaGenerationsURL(generationSchemaID int) string {
	return fmt.Sprintf("%s%s/%d/generations", handlers.GetAPIHostURL(), ResourceGenerationSchema, generationSchemaID)
}

func getSchemaGenerationURL(generationSchemaID, generationID int) string {
	return fmt.Sprintf("%s%s/%d", handlers.GetAPIHostURL(), ResourceGenerations, generationID)
}

func getGenerationsURL() string {
	return fmt.Sprintf("%s%s", handlers.GetAPIHostURL(), ResourceGenerations)
}

func getGenerationURL(generationID int) string {
	return fmt.Sprintf("%s%s/%d", handlers.GetAPIHostURL(), ResourceGenerations, generationID)
}

func getCalculatedMomentURL(generationSchemaID, generationID int) string {
	return fmt.Sprintf("%s%s/%d/formation-landscape", handlers.GetAPIHostURL(), ResourceGenerations, generationID)
}

func getGenerationalLandscapeURL(generationSchemaID, generationID int) string {
	return fmt.Sprintf("%s%s/%d/generational-landscape", handlers.GetAPIHostURL(), ResourceGenerations, generationID)
}

func getMomentURL(momentID int) string {
	return fmt.Sprintf("%s%s/%d", handlers.GetAPIHostURL(), ResourceMoments, momentID)
}

func getGenerationPositionsURL(generationID int) string {
	return fmt.Sprintf("%s%s/%d/full-positions", handlers.GetAPIHostURL(), ResourceGenerations, generationID)
}

func getSchemaGenerations(w http.ResponseWriter, r *http.Request, generationSchemaID int) (*model.Generations, error) {
	url := getSchemaGenerationsURL(generationSchemaID)
	code, body, err := handlers.GetResource(w, r, url)

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	var generations *model.Generations
	err = json.Unmarshal(body, &generations)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return generations, nil
}

func getSchemaGeneration(w http.ResponseWriter, r *http.Request, generationSchemaID, generationID int) (*model.Generation, error) {
	url := getSchemaGenerationURL(generationSchemaID, generationID)
	code, body, err := handlers.GetResource(w, r, url)

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	var generation *model.Generation
	err = json.Unmarshal(body, &generation)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return generation, nil
}

func getCalculatedLandscape(w http.ResponseWriter, r *http.Request, generationSchemaID, generationID int) (*model.HistoricalMoment, error) {
	url := getCalculatedMomentURL(generationSchemaID, generationID)
	code, body, err := handlers.GetResource(w, r, url)

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	var calculatedMoment *model.HistoricalMoment
	err = json.Unmarshal(body, &calculatedMoment)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return calculatedMoment, nil
}

func getHistoricalMoment(w http.ResponseWriter, r *http.Request, momentID int) (*model.HistoricalMoment, error) {
	url := getMomentURL(momentID)
	code, body, err := handlers.GetResource(w, r, url)

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	var moment *model.HistoricalMoment
	err = json.Unmarshal(body, &moment)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return moment, nil
}

func getGenerationalLandscape(w http.ResponseWriter, r *http.Request, generationSchemaID, generationID int) (*model.GenerationalLandscape, error) {
	url := getGenerationalLandscapeURL(generationSchemaID, generationID)
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

func getPositions(w http.ResponseWriter, r *http.Request, generationID int) ([]*model.GenerationFullPosition, error) {
	url := getGenerationPositionsURL(generationID)
	code, body, err := handlers.GetResource(w, r, url)

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	var positions []*model.GenerationFullPosition
	err = json.Unmarshal(body, &positions)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return positions, nil
}

func getUrlGenerationSchemaID(w http.ResponseWriter, r *http.Request) (int, error) {
	return handlers.GetUrlIntParam("schema_id", w, r)
}

func getUrlGenerationID(w http.ResponseWriter, r *http.Request) (int, error) {
	return handlers.GetUrlIntParam("generation_id", w, r)
}

func buildGenerationRequest(generation *model.Generation) *model.GenerationRequest {
	gr := &model.GenerationRequest{
		ID:                   generation.ID,
		Name:                 generation.Name,
		SchemaID:             generation.SchemaID,
		TypeID:               generation.Type.ID,
		StartYear:            generation.StartYear,
		EndYear:              generation.EndYear,
		PlaceID:              generation.Place.ID,
		FormationLandscapeID: generation.FormationLandscapeID,
		Description:          generation.Description,
	}

	return gr
}

func newGenerationRequest(schemaID int) *model.GenerationRequest {
	gr := &model.GenerationRequest{
		SchemaID:             schemaID,
	}

	return gr
}

func makeGenerationRequest(r *http.Request) (*model.GenerationRequest, error) {
	normalizedID := handlers.GetIntFormValue(r, "inputID")
	normalizedSchemaID := handlers.GetIntFormValue(r, "inputSchemaID")
	name := handlers.GetStringFormValue(r, "inputName")
	normalizedTypeID := handlers.GetIntFormValue(r, "inputTypeID")
	normalizedStartYear := handlers.GetIntFormValue(r, "inputStartYear")
	normalizedEndYear := handlers.GetIntFormValue(r, "inputEndYear")
	normalizedPlaceID := handlers.GetIntFormValue(r, "inputPlaceID")
	normalizedFormationLandscapeID := handlers.GetIntFormValue(r, "inputLandscapeID")
	description := handlers.GetStringFormValue(r, "inputDescription")

	gr := &model.GenerationRequest{
		ID:                   normalizedID,
		Name:                 name,
		SchemaID:             normalizedSchemaID,
		TypeID:               normalizedTypeID,
		StartYear:            normalizedStartYear,
		EndYear:              normalizedEndYear,
		PlaceID:              normalizedPlaceID,
		FormationLandscapeID: normalizedFormationLandscapeID,
		Description:          description,
	}

	return gr, nil
}

func postGeneration(w http.ResponseWriter, r *http.Request, generationRequest *model.GenerationRequest) (int, []byte, error) {
	url := getGenerationsURL()

	payload, err := json.Marshal(generationRequest)
	if err != nil {
		return 0, nil, err
	}

	if err != nil {
		return 0, nil, err
	}

	code, body, err := handlers.PostResource(w, r, url, payload)

	return code, body, nil
}

func patchGeneration(w http.ResponseWriter, r *http.Request, auRequest *model.GenerationRequest) (int, []byte, error) {
	url := getGenerationURL(auRequest.ID)

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

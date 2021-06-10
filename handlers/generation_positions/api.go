package generation_positions

import (
	"encoding/json"
	"fmt"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

const (
	ResourceSchema             = "generation-schemas"
	ResourceGenerationPosition = "generation-positions"
	ResourceGeneration         = "generations"
	ResourceLifePhase          = "life-phases"
)

func getGenerationPositionsURL() string {
	return fmt.Sprintf("%s%s", handlers.GetAPIHostURL(), ResourceGenerationPosition)
}

func getSchemaGenerationPositionURL(generationSchemaID, generationPositionID int) string {
	return fmt.Sprintf("%s%s/%d", handlers.GetAPIHostURL(), ResourceGenerationPosition, generationPositionID)
}

func getGenerationURL(generationID int) string {
	return fmt.Sprintf("%s%s/%d", handlers.GetAPIHostURL(), ResourceGeneration, generationID)
}

func getSchemaLifePhasesURL(schemaID int) string {
	return fmt.Sprintf("%s%s/%d/life-phases", handlers.GetAPIHostURL(), ResourceSchema, schemaID)
}

func getLifePhaseURL(lifePhaseID int) string {
	return fmt.Sprintf("%s%s/%d", handlers.GetAPIHostURL(), ResourceLifePhase, lifePhaseID)
}

func getSchemaMomentsURL(schemaID int) string {
	return fmt.Sprintf("%s%s/%d/moments", handlers.GetAPIHostURL(), ResourceSchema, schemaID)
}

func getSchemaGenerationPosition(w http.ResponseWriter, r *http.Request, generationSchemaID, generationPositionID int) (*model.GenerationPosition, error) {
	url := getSchemaGenerationPositionURL(generationSchemaID, generationPositionID)
	code, body, err := handlers.GetResource(w, r, url)

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	var generationPosition *model.GenerationPosition
	err = json.Unmarshal(body, &generationPosition)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return generationPosition, nil
}

func getGeneration(w http.ResponseWriter, r *http.Request, generationID int) (*model.Generation, error) {
	url := getGenerationURL(generationID)
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

func getSchemaLifePhases(w http.ResponseWriter, r *http.Request, schemaID int) ([]*model.LifePhase, error ){
	url := getSchemaLifePhasesURL(schemaID)
	code, body, err := handlers.GetResource(w, r, url)

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	var lifePhases []*model.LifePhase
	err = json.Unmarshal(body, &lifePhases)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return lifePhases, nil
}

func getSchemaMoments(w http.ResponseWriter, r *http.Request, schemaID int) ([]*model.HistoricalMoment, error ){
	url := getSchemaMomentsURL(schemaID)
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

	return moments.HistoricalMoments, nil
}

func getLifePhase(w http.ResponseWriter, r *http.Request, lifePhaseID int) (*model.LifePhase, error) {
	url := getLifePhaseURL(lifePhaseID)
	code, body, err := handlers.GetResource(w, r, url)

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	var lifePhase *model.LifePhase
	err = json.Unmarshal(body, &lifePhase)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return lifePhase, nil
}

func getUrlGenerationSchemaID(w http.ResponseWriter, r *http.Request) (int, error) {
	return handlers.GetUrlIntParam("schema_id", w, r)
}

func getUrlGenerationID(w http.ResponseWriter, r *http.Request) (int, error) {
	return handlers.GetUrlIntParam("generation_id", w, r)
}

func getUrlGenerationPositionID(w http.ResponseWriter, r *http.Request) (int, error) {
	return handlers.GetUrlIntParam("generation_position_id", w, r)
}

func buildGenerationPositionRequest(generationPosition *model.GenerationPosition) *model.GenerationPositionRequest {
	gr := &model.GenerationPositionRequest{
		ID:           generationPosition.ID,
		MomentID:     generationPosition.MomentID,
		Name:         generationPosition.Name,
		Ordinal:      generationPosition.Ordinal,
		LifePhaseID:  generationPosition.LifePhase.ID,
		GenerationID: generationPosition.Generation.ID,
	}

	return gr
}

func newGenerationPositionRequest(generationID int) *model.GenerationPositionRequest {
	gr := &model.GenerationPositionRequest{
		GenerationID: generationID,
	}

	return gr
}

func makeGenerationPositionRequest(r *http.Request) (*model.GenerationPositionRequest, error) {
	normalizedID := handlers.GetIntFormValue(r, "inputID")
	momentID := handlers.GetIntFormValue(r, "inputMomentID")
	name := handlers.GetStringFormValue(r, "inputName")
	ordinal := handlers.GetIntFormValue(r, "inputOrdinal")
	lifePhaseID := handlers.GetIntFormValue(r, "inputLifePhaseID")
	generationID := handlers.GetIntFormValue(r, "inputGenerationID")

	gpr := &model.GenerationPositionRequest{
		ID:           normalizedID,
		MomentID:     momentID,
		Name:         name,
		Ordinal:      ordinal,
		LifePhaseID:  lifePhaseID,
		GenerationID: generationID,
	}

	return gpr, nil
}

func postGenerationPosition(w http.ResponseWriter, r *http.Request, gpRequest *model.GenerationPositionRequest) (int, []byte, error) {
	url := getGenerationPositionsURL()

	payload, err := json.Marshal(gpRequest)
	if err != nil {
		return 0, nil, err
	}

	if err != nil {
		return 0, nil, err
	}

	code, body, err := handlers.PostResource(w, r, url, payload)

	return code, body, nil
}

func patchGenerationPosition(w http.ResponseWriter, r *http.Request, gpRequest *model.GenerationPositionRequest) (int, []byte, error) {
	url := getSchemaGenerationPositionURL(0, gpRequest.ID)

	payload, err := json.Marshal(gpRequest)
	if err != nil {
		return 0, nil, err
	}

	code, body, err := handlers.PatchResource(w, r, url, payload)
	if err != nil {
		return 0, nil, err
	}


	return code, body, nil
}


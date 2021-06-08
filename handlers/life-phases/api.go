package life_phases

import (
	"encoding/json"
	"fmt"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

const (
	ResourceLifePhase = "life-phases"
)

func getLifePhasesURL() string {
	return fmt.Sprintf("%s%s", handlers.GetAPIHostURL(), ResourceLifePhase)
}

func getSchemaLifePhaseURL(momentSchemaID, lifePhaseID int) string {
	return fmt.Sprintf("%s%s/%d", handlers.GetAPIHostURL(), ResourceLifePhase, lifePhaseID)
}

func getSchemaLifePhase(w http.ResponseWriter, r *http.Request, lifePhaseSchemaID, lifePhaseID int) (*model.LifePhase, error) {
	url := getSchemaLifePhaseURL(lifePhaseSchemaID, lifePhaseID)
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

func getUrlLifePhaseID(w http.ResponseWriter, r *http.Request) (int, error) {
	return handlers.GetUrlIntParam("life_phase_id", w, r)
}

func buildLifePhaseRequest(lifePhase *model.LifePhase) *model.LifePhaseRequest {
	mtr := &model.LifePhaseRequest{
		ID:        lifePhase.ID,
		SchemaID:  lifePhase.SchemaID,
		Name:      lifePhase.Name,
		StartYear: lifePhase.StartYear,
		EndYear:   lifePhase.EndYear,
		Role:      lifePhase.Role,
	}

	return mtr
}

func newLifePhaseRequest(schemaID int) *model.LifePhaseRequest {
	gr := &model.LifePhaseRequest{
		SchemaID: schemaID,
	}

	return gr
}

func makeLifePhaseRequest(r *http.Request) (*model.LifePhaseRequest, error) {
	normalizedID := handlers.GetIntFormValue(r, "inputID")
	normalizedSchemaID := handlers.GetIntFormValue(r, "inputSchemaID")
	name := handlers.GetStringFormValue(r, "inputName")
	normalizedStartYear := handlers.GetIntFormValue(r, "inputStartYear")
	normalizedEndYear := handlers.GetIntFormValue(r, "inputEndYear")
	role := handlers.GetStringFormValue(r, "inputRole")

	lpr := &model.LifePhaseRequest{
		ID:        normalizedID,
		SchemaID:  normalizedSchemaID,
		Name:      name,
		StartYear: normalizedStartYear,
		EndYear:   normalizedEndYear,
		Role:      role,
	}

	return lpr, nil
}

func postLifePhase(w http.ResponseWriter, r *http.Request, lifePhaseRequest *model.LifePhaseRequest) (int, []byte, error) {
	url := getLifePhasesURL()

	payload, err := json.Marshal(lifePhaseRequest)
	if err != nil {
		return 0, nil, err
	}

	if err != nil {
		return 0, nil, err
	}

	code, body, err := handlers.PostResource(w, r, url, payload)

	return code, body, nil
}

func patchLifePhase(w http.ResponseWriter, r *http.Request, lifePhaseRequest *model.LifePhaseRequest) (int, []byte, error) {
	url := getSchemaLifePhaseURL(lifePhaseRequest.SchemaID, lifePhaseRequest.ID)

	payload, err := json.Marshal(lifePhaseRequest)
	if err != nil {
		return 0, nil, err
	}

	code, body, err := handlers.PatchResource(w, r, url, payload)
	if err != nil {
		return 0, nil, err
	}


	return code, body, nil
}

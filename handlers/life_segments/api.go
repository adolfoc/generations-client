package life_segments

import (
	"encoding/json"
	"fmt"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

const (
	ResourceLifeSegments = "life-segments"
	ResourcePersons      = "persons"
)

func getLifeSegmentURL(generalifeSegmentID int) string {
	return fmt.Sprintf("%s%s/%d", handlers.GetAPIHostURL(), ResourceLifeSegments, generalifeSegmentID)
}

func getPersonURL(personID int) string {
	return fmt.Sprintf("%s%s/%d", handlers.GetAPIHostURL(), ResourcePersons, personID)
}

func getLifeSegment(w http.ResponseWriter, r *http.Request, lifeSegmentID int) (*model.LifeSegment, error) {
	url := getLifeSegmentURL(lifeSegmentID)
	code, body, err := handlers.GetResource(w, r, url)

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	var lifeSegment *model.LifeSegment
	err = json.Unmarshal(body, &lifeSegment)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return lifeSegment, nil
}

func getPerson(w http.ResponseWriter, r *http.Request, personID int) (*model.Person, error) {
	url := getPersonURL(personID)
	code, body, err := handlers.GetResource(w, r, url)

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	var person *model.Person
	err = json.Unmarshal(body, &person)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return person, nil
}

func getUrlLifeSegmentID(w http.ResponseWriter, r *http.Request) (int, error) {
	return handlers.GetUrlIntParam("life_segment_id", w, r)
}

func buildLifeSegmentRequest(lifeSegment *model.LifeSegment) *model.LifeSegmentRequest {
	gr := &model.LifeSegmentRequest{
		ID:          lifeSegment.ID,
		PersonID:    lifeSegment.PersonID,
		LifePhaseID: lifeSegment.LifePhase.ID,
		Summary:     lifeSegment.Summary,
		Description: lifeSegment.Description,
	}

	return gr
}

func makeLifeSegmentRequest(r *http.Request) (*model.LifeSegmentRequest, error) {
	normalizedID := handlers.GetIntFormValue(r, "inputID")
	personID := handlers.GetIntFormValue(r, "inputPersonID")
	lifePhaseID := handlers.GetIntFormValue(r, "inputLifePhaseID")
	summary := handlers.GetStringFormValue(r, "inputSummary")
	description := handlers.GetStringFormValue(r, "inputDescription")

	lfRequest := &model.LifeSegmentRequest{
		ID:          normalizedID,
		PersonID:    personID,
		LifePhaseID: lifePhaseID,
		Summary:     summary,
		Description: description,
	}

	return lfRequest, nil
}

func patchLifeSegment(w http.ResponseWriter, r *http.Request, lsRequest *model.LifeSegmentRequest) (int, []byte, error) {
	url := getLifeSegmentURL(lsRequest.ID)

	payload, err := json.Marshal(lsRequest)
	if err != nil {
		return 0, nil, err
	}

	code, body, err := handlers.PatchResource(w, r, url, payload)
	if err != nil {
		return 0, nil, err
	}


	return code, body, nil
}

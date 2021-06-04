package persons

import (
	"encoding/json"
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

func AddSegments(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-persons", "AddSegments")

	if handlers.UserAuthenticated(w, r) == false {
		log.FailedReturn()
		handlers.RedirectToSessionExpired(w, r)
		return
	}

	personID, err := getUrlPersonID(w, r)
	if err != nil {
		log.FailedReturn()
		return
	}

	person, err := getPerson(w, r, personID)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	schemas, err := getGenerationSchemas(w, r)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	url := fmt.Sprintf("/persons/%d/generate-life-segments", personID)
	addSegmentForm, err := MakeAddSegmentsForm(w, r, url, GetLabel(PersonAddSegmentsPageTitleIndex),
		GetLabel(PersonAddSegmentSubmitLabelIndex), person, schemas, 0, handlers.ResponseErrors{})
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView(AddSegmentsFormTemplate, addSegmentForm, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}

func getGenerationSchemas(w http.ResponseWriter, r *http.Request) ([]*model.GenerationSchema, error) {
	url := fmt.Sprintf("%sgeneration-schemas", handlers.GetAPIHostURL())
	code, body, err := handlers.GetResource(w, r, url)

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	if err != nil {
		return nil, err
	}

	var generationSchemas *model.GenerationSchemas
	err = json.Unmarshal(body, &generationSchemas)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return generationSchemas.GenerationSchemas, nil
}


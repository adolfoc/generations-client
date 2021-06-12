package life_segments

import (
	"encoding/json"
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

func EditLifeSegment(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-life_segments", "EditLifeSegment")

	if handlers.UserAuthenticated(w, r) == false {
		log.FailedReturn()
		handlers.RedirectToSessionExpired(w, r)
		return
	}

	lifeSegmentID, err := getUrlLifeSegmentID(w, r)
	if err != nil {
		log.FailedReturn()
		return
	}

	lifeSegment, err := getLifeSegment(w, r, lifeSegmentID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	person, err := getPerson(w, r, lifeSegment.PersonID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	lifeSegmentRequest := buildLifeSegmentRequest(lifeSegment)

	url := fmt.Sprintf("/persons/%d/life-segments/%d/update", lifeSegment.PersonID, lifeSegmentID)
	generationForm, err := MakeLifeSegmentForm(w, r, url, GetLabel(LifeSegmentEditPageTitleIndex), "",
		GetLabel(LifeSegmentEditSubmitLabelIndex), lifeSegment, lifeSegmentRequest, person, handlers.ResponseErrors{})
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView(LifeSegmentFormTemplate, generationForm, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}

func EditLifeSegmentRetry(w http.ResponseWriter, r *http.Request, lsRequest *model.LifeSegmentRequest, errors handlers.ResponseErrors) {
	log := common.StartLog("handlers-life_segments", "EditLifeSegmentRetry")

	lifePhase, err := getLifePhase(w, r, lsRequest.LifePhaseID)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	person, err := getPerson(w, r, lsRequest.PersonID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	lifeSegment := &model.LifeSegment{
		ID:          lsRequest.ID,
		PersonID:    lsRequest.PersonID,
		LifePhase:   lifePhase,
		Summary:     lsRequest.Summary,
		Description: lsRequest.Description,
	}

	url := fmt.Sprintf("/life-segments/%d/update", lsRequest.ID)
	generationForm, err := MakeLifeSegmentForm(w, r, url, GetLabel(LifeSegmentEditPageTitleIndex), "",
		GetLabel(LifeSegmentEditSubmitLabelIndex), lifeSegment, lsRequest, person, errors)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView(LifeSegmentFormTemplate, generationForm, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
	return
}

const (
	LifePhaseURL = "life-phases"
)

func getLifePhase(w http.ResponseWriter, r *http.Request, lifePhaseID int) (*model.LifePhase, error) {
	url := fmt.Sprintf("%s%s/%d/generation-types", handlers.GetAPIHostURL(), LifePhaseURL, lifePhaseID)
	code, body, err := handlers.GetResource(w, r, url)

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	if err != nil {
		return nil, err
	}

	var lifePhase *model.LifePhase
	err = json.Unmarshal(body, &lifePhase)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return lifePhase, nil
}

package life_segments

import (
	"fmt"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

const (
	LifeSegmentFormTemplate = "life_segment_form"
)

type LifeSegmentForm struct {
	Ft          handlers.FormTemplate
	LifeSegment *model.LifeSegment
}

func makeLifeSegmentFormValues(lsRequest *model.LifeSegmentRequest) map[string]interface{} {
	formValues := make(map[string]interface{})
	formValues["ID"] = fmt.Sprintf("%d", lsRequest.ID)
	formValues["PersonID"] = fmt.Sprintf("%d", lsRequest.PersonID)
	formValues["LifePhaseID"] = fmt.Sprintf("%d", lsRequest.LifePhaseID)
	formValues["Summary"] = lsRequest.Summary
	formValues["Description"] = lsRequest.Description

	return formValues
}

func makeLifeSegmentErrorMessages(errors handlers.ResponseErrors) map[string]string {
	formErrorMessages := make(map[string]string)
	for _, error := range errors.Errors {
		if error.Source.Pointer == "data/attributes/person_id" {
			formErrorMessages["PersonID"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/life_phase_id" {
			formErrorMessages["LifePhaseID"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/summary" {
			formErrorMessages["Summary"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/description" {
			formErrorMessages["Description"] = error.Detail
		}
	}

	return formErrorMessages
}

func MakeLifeSegmentForm(w http.ResponseWriter, r *http.Request, url string, pageTitle, submitLabel string, lifeSegment *model.LifeSegment,
	lsRequest *model.LifeSegmentRequest, errors handlers.ResponseErrors) (*LifeSegmentForm, error) {

	formValues := makeLifeSegmentFormValues(lsRequest)
	formErrorMessages := makeLifeSegmentErrorMessages(errors)
	ft, err := handlers.MakeFormTemplate(r, url, pageTitle, submitLabel, formValues, formErrorMessages)
	if err != nil {
		return nil, err
	}

	lfTemplate := &LifeSegmentForm{
		Ft:          *ft,
		LifeSegment: lifeSegment,
	}

	return lfTemplate, nil
}

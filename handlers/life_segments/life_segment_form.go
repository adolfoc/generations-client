package life_segments

import (
	"fmt"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
	"strconv"
	"strings"
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
	formValues["ID"] = lsRequest.ID
	formValues["PersonID"] = lsRequest.PersonID
	formValues["LifePhaseID"] = lsRequest.LifePhaseID
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

func extractYear(date string) int {
	parts := strings.Split(date, "-")
	if len(parts) > 0 {
		year, _ := strconv.Atoi(parts[0])
		return year
	}

	return 0
}

func MakeLifeSegmentForm(w http.ResponseWriter, r *http.Request, url string, pageTitle, studyTitle, submitLabel string, lifeSegment *model.LifeSegment,
	lsRequest *model.LifeSegmentRequest, person *model.Person, errors handlers.ResponseErrors) (*LifeSegmentForm, error) {

	yearOfBirth := extractYear(person.BirthDate)
	fullPageTitle := fmt.Sprintf("%s: %d-%d años (%d-%d)",
		pageTitle, lifeSegment.LifePhase.StartYear, lifeSegment.LifePhase.EndYear,
		yearOfBirth + lifeSegment.LifePhase.StartYear, yearOfBirth + lifeSegment.LifePhase.EndYear)
	formValues := makeLifeSegmentFormValues(lsRequest)
	formErrorMessages := makeLifeSegmentErrorMessages(errors)
	ft, err := handlers.MakeFormTemplate(r, url, fullPageTitle, studyTitle, submitLabel, formValues, formErrorMessages)
	if err != nil {
		return nil, err
	}

	lfTemplate := &LifeSegmentForm{
		Ft:          *ft,
		LifeSegment: lifeSegment,
	}

	return lfTemplate, nil
}

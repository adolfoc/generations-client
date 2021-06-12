package moments

import (
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

const (
	MomentFormTemplate = "moment_form"
)

type MomentForm struct {
	Ft          handlers.FormTemplate
	SchemaID    int
	Moment      *model.HistoricalMoment
	MomentTypes []*model.MomentType
}

func makeMomentFormValues(momentRequest *model.HistoricalMomentRequest) map[string]interface{} {
	formValues := make(map[string]interface{})
	formValues["ID"] = momentRequest.ID
	formValues["Name"] = momentRequest.Name
	formValues["SchemaID"] = momentRequest.SchemaID
	formValues["TypeID"] = momentRequest.TypeID
	formValues["Start"] = momentRequest.Start
	formValues["End"] = momentRequest.End
	formValues["Summary"] = momentRequest.Summary
	formValues["Description"] = momentRequest.Description

	return formValues
}

func makeMomentErrorMessages(errors handlers.ResponseErrors) map[string]string {
	formErrorMessages := make(map[string]string)
	for _, error := range errors.Errors {
		if error.Source.Pointer == "data/attributes/name" {
			formErrorMessages["Name"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/schema_id" {
			formErrorMessages["SchemaID"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/moment_type_id" {
			formErrorMessages["TypeID"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/start_date" {
			formErrorMessages["Start"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/end_date" {
			formErrorMessages["End"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/summary" {
			formErrorMessages["Summary"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/description" {
			formErrorMessages["FormationLandscapeID"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/description" {
			formErrorMessages["Description"] = error.Detail
		}
	}

	return formErrorMessages
}

func MakeMomentForm(w http.ResponseWriter, r *http.Request, url string, pageTitle, studyTitle, submitLabel string,
	moment *model.HistoricalMoment, momentRequest *model.HistoricalMomentRequest, momentTypes []*model.MomentType,
	errors handlers.ResponseErrors) (*MomentForm, error) {

	formValues := makeMomentFormValues(momentRequest)
	formErrorMessages := makeMomentErrorMessages(errors)
	ft, err := handlers.MakeFormTemplate(r, url, pageTitle, studyTitle, submitLabel, formValues, formErrorMessages)
	if err != nil {
		return nil, err
	}

	momentTemplate := &MomentForm{
		Ft:          *ft,
		SchemaID:    momentRequest.SchemaID,
		Moment:      moment,
		MomentTypes: momentTypes,
	}

	return momentTemplate, nil
}


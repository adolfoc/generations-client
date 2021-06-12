package generational_landscape

import (
	"fmt"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

const (
	GenerationalLandscapeFormTemplate = "generational_landscape_form"
)

type GenerationalLandscapeForm struct {
	Ft                    handlers.FormTemplate
	SchemaID              int
	GenerationalLandscape *model.GenerationalLandscape
	FormationMoments      *model.HistoricalMoments
}

func makeGenerationalLandscapeFormValues(generationalLandscapeRequest *model.GenerationalLandscapeRequest) map[string]interface{} {
	formValues := make(map[string]interface{})
	formValues["ID"] = generationalLandscapeRequest.ID
	formValues["GenerationID"] = generationalLandscapeRequest.GenerationID
	formValues["FormationMomentID"] = generationalLandscapeRequest.FormationMomentID
	formValues["Description"] = generationalLandscapeRequest.Description

	return formValues
}

func makeGenerationalLandscapeErrorMessages(errors handlers.ResponseErrors) map[string]string {
	formErrorMessages := make(map[string]string)
	for _, error := range errors.Errors {
		if error.Source.Pointer == "data/attributes/generation_id" {
			formErrorMessages["GenerationID"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/formation_moment_id" {
			formErrorMessages["FormationMomentID"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/description" {
			formErrorMessages["Description"] = error.Detail
		}
	}

	return formErrorMessages
}

func MakeGenerationalLandscapeForm(w http.ResponseWriter, r *http.Request, url string, pageTitle, studyTitle, submitLabel string,
	generationalLandscape *model.GenerationalLandscape, glRequest *model.GenerationalLandscapeRequest, schemaID int,
	formationMoments *model.HistoricalMoments, generation *model.Generation, errors handlers.ResponseErrors) (*GenerationalLandscapeForm, error) {

	fullTitle := fmt.Sprintf("%s: %s (%d-%d)", pageTitle, generation.Name, generation.StartYear, generation.EndYear)
	formValues := makeGenerationalLandscapeFormValues(glRequest)
	formErrorMessages := makeGenerationalLandscapeErrorMessages(errors)
	ft, err := handlers.MakeFormTemplate(r, url, fullTitle, studyTitle, submitLabel, formValues, formErrorMessages)
	if err != nil {
		return nil, err
	}

	generationalLandscapeTemplate := &GenerationalLandscapeForm{
		Ft:                    *ft,
		SchemaID:              schemaID,
		GenerationalLandscape: generationalLandscape,
		FormationMoments:      formationMoments,
	}

	return generationalLandscapeTemplate, nil
}

package generations

import (
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

const (
	GenerationFormTemplate = "generation_form"
)

type GenerationForm struct {
	Ft                  handlers.FormTemplate
	SchemaID            int
	Generation          *model.Generation
	GenerationTypes     []*model.GenerationType
	Places              []*model.Place
	FormationLandscapes []*model.HistoricalMoment
}

func makeGenerationFormValues(generationRequest *model.GenerationRequest) map[string]interface{} {
	formValues := make(map[string]interface{})
	formValues["ID"] = generationRequest.ID
	formValues["Name"] = generationRequest.Name
	formValues["SchemaID"] = generationRequest.SchemaID
	formValues["TypeID"] = generationRequest.TypeID
	formValues["StartYear"] = generationRequest.StartYear
	formValues["EndYear"] = generationRequest.EndYear
	formValues["PlaceID"] = generationRequest.PlaceID
	formValues["FormationLandscapeID"] = generationRequest.FormationLandscapeID
	formValues["Description"] = generationRequest.Description

	return formValues
}

func makeGenerationErrorMessages(errors handlers.ResponseErrors) map[string]string {
	formErrorMessages := make(map[string]string)
	for _, error := range errors.Errors {
		if error.Source.Pointer == "data/attributes/name" {
			formErrorMessages["Name"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/schema_id" {
			formErrorMessages["SchemaID"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/generation_type_id" {
			formErrorMessages["TypeID"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/start_year" {
			formErrorMessages["StartYear"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/end_year" {
			formErrorMessages["EndYear"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/place_id" {
			formErrorMessages["PlaceID"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/formation_landscape_id" {
			formErrorMessages["FormationLandscapeID"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/description" {
			formErrorMessages["Description"] = error.Detail
		}
	}

	return formErrorMessages
}

func MakeGenerationForm(w http.ResponseWriter, r *http.Request, url string, pageTitle, submitLabel string, generation *model.Generation,
	generationRequest *model.GenerationRequest, generationTypes []*model.GenerationType, places []*model.Place,
	formationLandscapes []*model.HistoricalMoment, errors handlers.ResponseErrors) (*GenerationForm, error) {

	formValues := makeGenerationFormValues(generationRequest)
	formErrorMessages := makeGenerationErrorMessages(errors)
	ft, err := handlers.MakeFormTemplate(r, url, pageTitle, submitLabel, formValues, formErrorMessages)
	if err != nil {
		return nil, err
	}

	generationTemplate := &GenerationForm{
		Ft:                  *ft,
		SchemaID:            generationRequest.SchemaID,
		Generation:          generation,
		GenerationTypes:     generationTypes,
		Places:              places,
		FormationLandscapes: formationLandscapes,
	}

	return generationTemplate, nil
}

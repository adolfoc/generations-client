package generation_positions

import (
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

const (
	GenerationPositionFormTemplate = "generation_position_form"
)

type GenerationPositionForm struct {
	Ft                 handlers.FormTemplate
	SchemaID           int
	GenerationPosition *model.GenerationPosition
	LifePhases         []*model.LifePhase
	Moments            []*model.HistoricalMoment
}

func makeGenerationPositionFormValues(generationPositionRequest *model.GenerationPositionRequest) map[string]interface{} {
	formValues := make(map[string]interface{})
	formValues["ID"] = generationPositionRequest.ID
	formValues["MomentID"] = generationPositionRequest.MomentID
	formValues["Name"] = generationPositionRequest.Name
	formValues["Ordinal"] = generationPositionRequest.Ordinal
	formValues["LifePhaseID"] = generationPositionRequest.LifePhaseID
	formValues["GenerationID"] = generationPositionRequest.GenerationID

	return formValues
}

func makeGenerationPositionErrorMessages(errors handlers.ResponseErrors) map[string]string {
	formErrorMessages := make(map[string]string)
	for _, error := range errors.Errors {
		if error.Source.Pointer == "data/attributes/moment_id" {
			formErrorMessages["MomentID"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/name" {
			formErrorMessages["Name"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/ordinal" {
			formErrorMessages["Ordinal"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/life_phase_id" {
			formErrorMessages["LifePhaseID"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/generation_id" {
			formErrorMessages["GenerationID"] = error.Detail
		}
	}

	return formErrorMessages
}

func MakeGenerationPositionForm(w http.ResponseWriter, r *http.Request, url string, pageTitle, studyTitle, submitLabel string,
	generationPosition *model.GenerationPosition, gpRequest *model.GenerationPositionRequest, schemaID int,
	lifePhases []*model.LifePhase, moments []*model.HistoricalMoment, errors handlers.ResponseErrors) (*GenerationPositionForm, error) {

	formValues := makeGenerationPositionFormValues(gpRequest)
	formErrorMessages := makeGenerationPositionErrorMessages(errors)
	ft, err := handlers.MakeFormTemplate(r, url, pageTitle, studyTitle, submitLabel, formValues, formErrorMessages)
	if err != nil {
		return nil, err
	}

	generationPositionTemplate := &GenerationPositionForm{
		Ft:                 *ft,
		SchemaID:           schemaID,
		GenerationPosition: generationPosition,
		LifePhases:         lifePhases,
		Moments:            moments,
	}

	return generationPositionTemplate, nil
}


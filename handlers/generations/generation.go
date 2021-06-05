package generations

import (
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

type GenerationTemplate struct {
	Ct                        handlers.CommonTemplate
	SchemaID                  int
	Generation                *model.Generation
	HaveMoment                bool
	FormationMoment           *model.HistoricalMoment
	HaveCalculatedMoment      bool
	CalculatedFormationMoment *model.HistoricalMoment
	HaveLandscape             bool
	GenerationalLandscape     *model.GenerationalLandscape
	HavePositions             bool
	Positions                 []*model.GenerationFullPosition
	Cohort                    []*model.Person
}

func MakeGenerationTemplate(r *http.Request, pageTitle string, generation *model.Generation,
	formationMoment *model.HistoricalMoment, calculatedFormationMoment *model.HistoricalMoment,
	generationalLandscape *model.GenerationalLandscape, positions []*model.GenerationFullPosition, persons []*model.Person) (*GenerationTemplate, error) {
	ct, err := handlers.MakeCommonTemplate(r, pageTitle)
	if err != nil {
		return nil, err
	}

	haveMoment := false
	if formationMoment != nil && formationMoment.ID > 0 {
		haveMoment = true
	}

	haveCalculatedMoment := false
	if calculatedFormationMoment != nil && calculatedFormationMoment.ID > 0 {
		haveCalculatedMoment = true
	}

	haveLandscape := false
	if generationalLandscape != nil && generationalLandscape.ID > 0 {
		haveLandscape = true
	}

	havePositions := false
	if len(positions) > 0 {
		havePositions = true
	}
	generationTemplate := &GenerationTemplate{
		Ct:                        *ct,
		SchemaID:                  generation.SchemaID,
		Generation:                generation,
		HaveMoment:                haveMoment,
		FormationMoment:           formationMoment,
		HaveCalculatedMoment:      haveCalculatedMoment,
		CalculatedFormationMoment: calculatedFormationMoment,
		HaveLandscape:             haveLandscape,
		GenerationalLandscape:     generationalLandscape,
		HavePositions:             havePositions,
		Positions:                 positions,
		Cohort:                    persons,
	}

	return generationTemplate, nil
}

func GetGeneration(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-schemas", "GetGeneration")

	if handlers.UserAuthenticated(w, r) == false {
		log.FailedReturn()
		handlers.RedirectToSessionExpired(w, r)
		return
	}

	schemaID, err := getUrlGenerationSchemaID(w, r)
	if err != nil {
		log.FailedReturn()
		return
	}

	generationID, err := getUrlGenerationID(w, r)
	if err != nil {
		log.FailedReturn()
		return
	}

	generation, err := getSchemaGeneration(w, r, schemaID, generationID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	calculatedMoment, err := getCalculatedLandscape(w, r, schemaID, generationID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	generationalLandscape, err := getGenerationalLandscape(w, r, schemaID, generationID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	persons, err := getCohort(w, r, generation)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	var formationMoment *model.HistoricalMoment
	if generationalLandscape != nil && generationalLandscape.ID > 0 && generationalLandscape.FormationMomentID > 0 {
		formationMoment, err = getHistoricalMoment(w, r, generationalLandscape.FormationMomentID)
		if handlers.HandleError(w, r, err) {
			log.FailedReturn()
			return
		}
	}

	positions, err := getPositions(w, r, generationID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	ct, err := MakeGenerationTemplate(r, GetLabel(GenerationPageTitleIndex), generation,
		formationMoment, calculatedMoment, generationalLandscape, positions, persons)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView("generation", ct, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}

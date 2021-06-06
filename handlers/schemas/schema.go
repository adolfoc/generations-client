package schemas

import (
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

type GenerationSchemaTemplate struct {
	Ct                      handlers.CommonTemplate
	GenerationSchema        *model.GenerationSchema
	LifePhases              []*model.LifePhase
	GenerationTypes         []*model.GenerationType
	MomentTypes             []*model.MomentType
	AllowTemplateGeneration bool
}

func MakeGenerationSchemaTemplate(r *http.Request, pageTitle string, gs *model.GenerationSchema,
	lifePhases []*model.LifePhase, generationTypes []*model.GenerationType, momentTypes []*model.MomentType) (*GenerationSchemaTemplate, error) {
	ct, err := handlers.MakeCommonTemplate(r, pageTitle)
	if err != nil {
		return nil, err
	}

	allowTemplateGeneration := true
	if len(lifePhases) > 0 && len(generationTypes) > 0 && len(momentTypes) > 0 {
		allowTemplateGeneration = false
	}
	schemaTemplate := &GenerationSchemaTemplate{
		Ct:                      *ct,
		GenerationSchema:        gs,
		LifePhases:              lifePhases,
		GenerationTypes:         generationTypes,
		MomentTypes:             momentTypes,
		AllowTemplateGeneration: allowTemplateGeneration,
	}

	return schemaTemplate, nil
}

func GetGenerationSchema(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-schemas", "GetGenerationSchema")

	if handlers.UserAuthenticated(w, r) == false {
		log.FailedReturn()
		handlers.RedirectToSessionExpired(w, r)
		return
	}

	gsID, err := getUrlGenerationSchemaID(w, r)
	if err != nil {
		log.FailedReturn()
		return
	}

	generationSchema, err := getGenerationSchema(w, r, gsID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	lifePhases, err := getLifePhases(w, r, gsID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	generationTypes, err := getGenerationTypes(w, r, gsID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	momentTypes, err := getMomentTypes(w, r, gsID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	ct, err := MakeGenerationSchemaTemplate(r, GetLabel(GenerationSchemaPageTitleIndex), generationSchema, lifePhases, generationTypes, momentTypes)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView("schema", ct, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}

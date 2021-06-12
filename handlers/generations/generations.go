package generations

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

type GenerationsTemplate struct {
	Ct          handlers.CommonTemplate
	SchemaID    int
	Generations *model.Generations
	Pagination  *handlers.Pagination
}

func getPaginationBaseURL(generationSchemaID int) string {
	stem := fmt.Sprintf("/schemas/%d/generations", generationSchemaID)
	return stem + "?page=%d"
}

func MakeGenerationsTemplate(r *http.Request, pageTitle, studyTitle string, generationSchema, page int, generations *model.Generations) (*GenerationsTemplate, error) {
	ct, err := handlers.MakeCommonTemplate(r, pageTitle, studyTitle)
	if err != nil {
		return nil, err
	}

	pagination := handlers.MakePagination(generations.RecordCount, getPaginationBaseURL(generationSchema), page)

	generationTemplate := &GenerationsTemplate{
		Ct:          *ct,
		SchemaID:    generationSchema,
		Generations: generations,
		Pagination:  pagination,
	}

	return generationTemplate, nil
}

func GetSchemaGenerations(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-schemas", "GetSchemaGenerations")

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

	generationSchema, err := handlers.GetGenerationSchema(w, r, schemaID)
	if err != nil {
		log.FailedReturn()
		return
	}

	page := handlers.GetURLPageParameter(r)
	generations, err := getSchemaGenerations(w, r, schemaID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	ct, err := MakeGenerationsTemplate(r, GetLabel(GenerationIndexPageTitleIndex), generationSchema.MakeStudyTitle(),
		schemaID, page, generations)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView("generations", ct, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}

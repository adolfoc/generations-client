package schemas

import (
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

type TemplateGenerationSchemas struct {
	Ct                handlers.CommonTemplate
	GenerationSchemas *model.GenerationSchemas
	Pagination        *handlers.Pagination
}

func getPaginationBaseURL() string {
	return "/schemas/index?page=%d"
}

func MakeGenerationSchemasTemplate(r *http.Request, pageTitle, studyTitle string, schemas *model.GenerationSchemas, page int) (*TemplateGenerationSchemas, error) {
	ct, err := handlers.MakeCommonTemplate(r, pageTitle, studyTitle)
	if err != nil {
		return nil, err
	}

	pagination := handlers.MakePagination(schemas.RecordCount, getPaginationBaseURL(), page)

	generationSchemaTemplate := &TemplateGenerationSchemas{
		Ct:                *ct,
		GenerationSchemas: schemas,
		Pagination:        pagination,
	}

	return generationSchemaTemplate, err
}

func GetGenerationSchemas(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-schemas", "GetGenerationSchemas")

	if handlers.UserAuthenticated(w, r) == false {
		log.FailedReturn()
		handlers.RedirectToSessionExpired(w, r)
		return
	}

	page := handlers.GetURLPageParameter(r)
	affiliations, err := getGenerationSchemas(w, r)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	at, err := MakeGenerationSchemasTemplate(r, GetLabel(GenerationSchemaIndexPageTitleIndex), "", affiliations, page)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView("schemas", at, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}

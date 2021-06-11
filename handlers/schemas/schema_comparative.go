package schemas

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

type ComparativeRow struct {
	MomentName     string
	Span           string
	GenerationName string
	Role           string
	BirthYears     string
}

type ComparativeTemplate struct {
	Ct       handlers.CommonTemplate
	SchemaID int
	Rows     []*ComparativeRow
}

func buildRows(comparative *model.SchemaComparative) []*ComparativeRow {
	var rows []*ComparativeRow

	for _, item := range comparative.Items {
		momentName := item.Moment.Name
		momentSpan := fmt.Sprintf("%d--%d", item.Moment.StartYear(), item.Moment.EndYear())
		for _, pos := range item.Positions {
			row := &ComparativeRow{
				MomentName:     momentName,
				Span:           momentSpan,
				GenerationName: pos.Generation.Name,
				Role:           pos.Name,
				BirthYears:     fmt.Sprintf("%d--%d (%d)", pos.Generation.StartYear, pos.Generation.EndYear, pos.Generation.EndYear - pos.Generation.StartYear),
			}

			rows = append(rows, row)
			momentName = ""
			momentSpan = ""
		}
	}

	return rows
}

func MakeComparativeTemplate(r *http.Request, pageTitle string, sc *model.SchemaComparative, schemaID int) (*ComparativeTemplate, error) {
	ct, err := handlers.MakeCommonTemplate(r, pageTitle)
	if err != nil {
		return nil, err
	}

	comparativeTemplate := &ComparativeTemplate{
		Ct:       *ct,
		SchemaID: schemaID,
		Rows:     buildRows(sc),
	}

	return comparativeTemplate, nil
}

func GetComparative(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-schemas", "GetComparative")

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

	comparativeReport, err := getComparativeReport(w, r, schemaID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	ct, err := MakeComparativeTemplate(r, GetLabel(GenerationSchemaComparativePageTitleIndex), comparativeReport, schemaID)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView("comparative_report", ct, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}


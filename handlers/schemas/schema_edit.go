package schemas

import (
	"encoding/json"
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

func EditSchema(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-schemas", "EditSchema")

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

	schema, err := getGenerationSchema(w, r, schemaID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	places, err := getPlaces(w, r)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	schemaRequest := buildSchemaRequest(schema)

	url := fmt.Sprintf("/schemas/%d/update", schemaID)
	schemaForm, err := MakeSchemaForm(w, r, url, GetLabel(GenerationSchemaPageTitleIndex),
		GetLabel(GenerationSchemaEditSubmitLabelIndex), schema, schemaRequest, places, handlers.ResponseErrors{})
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView(SchemaFormTemplate, schemaForm, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}

func EditSchemaRetry(w http.ResponseWriter, r *http.Request, sRequest *model.GenerationSchemaRequest, errors handlers.ResponseErrors) {
	log := common.StartLog("handlers-schemas", "EditSchemaRetry")

	places, err := getPlaces(w, r)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	place := matchPlace(sRequest.PlaceID, places)

	schema := &model.GenerationSchema{
		ID:                    sRequest.ID,
		Name:                  sRequest.Name,
		Description:           sRequest.Description,
		StartYear:             sRequest.StartYear,
		EndYear:               sRequest.EndYear,
		MinimumGenerationSpan: sRequest.MinimumGenerationSpan,
		MaximumGenerationSpan: sRequest.MaximumGenerationSpan,
		Place:                 place,
	}

	url := fmt.Sprintf("/schemas/%d/update", sRequest.ID)
	schemaForm, err := MakeSchemaForm(w, r, url, GetLabel(GenerationSchemaPageTitleIndex),
		GetLabel(GenerationSchemaEditSubmitLabelIndex), schema, sRequest, places, errors)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView(SchemaFormTemplate, schemaForm, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
	return
}

func getPlaces(w http.ResponseWriter, r *http.Request) ([]*model.Place, error) {
	url := fmt.Sprintf("%splaces/", handlers.GetAPIHostURL())
	code, body, err := handlers.GetResource(w, r, url)

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	if err != nil {
		return nil, err
	}

	var places *model.Places
	err = json.Unmarshal(body, &places)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return places.Places, nil
}

func matchPlace(placeID int, places []*model.Place) *model.Place {
	for _, place := range places {
		if place.ID == placeID {
			return place
		}
	}

	return nil
}
package places

import (
	"encoding/json"
	"fmt"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

const (
	ResourcePlace = "places"
	ResourcePlaceType = "place-types"
)

func getPlacesURL(page int, column string) string {
	return fmt.Sprintf("%s%s?page[number]=%d&sort[column]=%s", handlers.GetAPIHostURL(), ResourcePlace, page, column)
}

func getSimplePlacesURL() string {
	return fmt.Sprintf("%s%s", handlers.GetAPIHostURL(), ResourcePlace)
}

func getPlaceURL(placeID int) string {
	return fmt.Sprintf("%s%s/%d", handlers.GetAPIHostURL(), ResourcePlace, placeID)
}

func getPlaceTypesURL() string {
	return fmt.Sprintf("%s%s", handlers.GetAPIHostURL(), ResourcePlaceType)
}

func getPlaces(w http.ResponseWriter, r *http.Request, page int, sortColumn string) (*model.Places, error) {
	url := getPlacesURL(page, sortColumn)
	code, body, err := handlers.GetResource(w, r, url)

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	var places *model.Places
	err = json.Unmarshal(body, &places)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return places, nil
}

func getUrlPlaceID(w http.ResponseWriter, r *http.Request) (int, error) {
	return handlers.GetUrlIntParam("place_id", w, r)
}

func getPlace(w http.ResponseWriter, r *http.Request, placeID int) (*model.Place, error) {
	url := getPlaceURL(placeID)
	code, body, err := handlers.GetResource(w, r, url)

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	var place *model.Place
	err = json.Unmarshal(body, &place)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return place, nil
}

func getPlaceTypes(w http.ResponseWriter, r *http.Request) ([]*model.PlaceType, error) {
	url := getPlaceTypesURL()
	code, body, err := handlers.GetResource(w, r, url)

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	var places []*model.PlaceType
	err = json.Unmarshal(body, &places)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return places, nil
}

func newPlaceRequest() *model.PlaceRequest {
	pRequest := &model.PlaceRequest{}

	return pRequest
}

func buildPlaceRequest(place *model.Place) *model.PlaceRequest {
	placeRequest := &model.PlaceRequest{
		ID:            place.ID,
		Name:          place.Name,
		PlaceTypeID:   place.PlaceTypeID,
		ParentPlaceID: place.ParentPlaceID,
		Start:         place.Start,
		End:           place.End,
		Summary:       place.Summary,
		Description:   place.Description,
	}

	return placeRequest
}

func makePlaceRequest(r *http.Request) (*model.PlaceRequest, error) {
	normalizedID := handlers.GetIntFormValue(r, "inputID")
	name := handlers.GetStringFormValue(r, "inputName")
	placeTypeID := handlers.GetIntFormValue(r, "inputPlaceTypeID")
	parentPlaceID := handlers.GetIntFormValue(r, "inputParentPlaceID")
	start := handlers.GetStringFormValue(r, "inputStart")
	end := handlers.GetStringFormValue(r, "inputend")
	summary := handlers.GetStringFormValue(r, "inputSummary")
	description := handlers.GetStringFormValue(r, "inputDescription")

	er := &model.PlaceRequest{
		ID:            normalizedID,
		Name:          name,
		PlaceTypeID:   placeTypeID,
		ParentPlaceID: parentPlaceID,
		Start:         start,
		End:           end,
		Summary:       summary,
		Description:   description,
	}

	return er, nil
}

func postPlace(w http.ResponseWriter, r *http.Request, placeRequest *model.PlaceRequest) (int, []byte, error) {
	url := getSimplePlacesURL()

	payload, err := json.Marshal(placeRequest)
	if err != nil {
		return 0, nil, err
	}

	if err != nil {
		return 0, nil, err
	}

	code, body, err := handlers.PostResource(w, r, url, payload)

	return code, body, nil
}

func patchPlace(w http.ResponseWriter, r *http.Request, placeRequest *model.PlaceRequest) (int, []byte, error) {
	url := getPlaceURL(placeRequest.ID)

	payload, err := json.Marshal(placeRequest)
	if err != nil {
		return 0, nil, err
	}

	code, body, err := handlers.PatchResource(w, r, url, payload)
	if err != nil {
		return 0, nil, err
	}

	return code, body, nil
}

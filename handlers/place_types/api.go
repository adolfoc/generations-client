package place_types

import (
	"encoding/json"
	"fmt"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

const (
	ResourcePlaceTypes = "place-types"
)

func getPlaceTypesURL(page int, column string) string {
	return fmt.Sprintf("%s%s?page[number]=%d&sort[column]=%s", handlers.GetAPIHostURL(), ResourcePlaceTypes, page, column)
}

func getSimplePlaceTypesURL() string {
	return fmt.Sprintf("%s%s", handlers.GetAPIHostURL(), ResourcePlaceTypes)
}

func getPlaceTypeURL(placeTypeID int) string {
	return fmt.Sprintf("%s%s/%d", handlers.GetAPIHostURL(), ResourcePlaceTypes, placeTypeID)
}

func getPlaceTypes(w http.ResponseWriter, r *http.Request, page int, column string) (*model.PlaceTypes, error) {
	url := getPlaceTypesURL(page, column)
	code, body, err := handlers.GetResource(w, r, url)

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	var placeTypes *model.PlaceTypes
	err = json.Unmarshal(body, &placeTypes)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return placeTypes, nil
}

func getUrlPlaceTypeID(w http.ResponseWriter, r *http.Request) (int, error) {
	return handlers.GetUrlIntParam("place_type_id", w, r)
}

func getPlaceType(w http.ResponseWriter, r *http.Request, placeTypeID int) (*model.PlaceType, error) {
	url := getPlaceTypeURL(placeTypeID)
	code, body, err := handlers.GetResource(w, r, url)

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	var placeType *model.PlaceType
	err = json.Unmarshal(body, &placeType)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return placeType, nil
}

func buildPlaceTypeRequest(placeType *model.PlaceType) *model.PlaceTypeRequest {
	mr := &model.PlaceTypeRequest{
		ID:          placeType.ID,
		Name:        placeType.Name,
		Description: placeType.Description,
	}

	return mr
}

func makePlaceTypeRequest(r *http.Request) (*model.PlaceTypeRequest, error) {
	normalizedID := handlers.GetIntFormValue(r, "inputID")
	name := handlers.GetStringFormValue(r, "inputName")
	description := handlers.GetStringFormValue(r, "inputDescription")

	ptr := &model.PlaceTypeRequest{
		ID:          normalizedID,
		Name:        name,
		Description: description,
	}

	return ptr, nil
}

func postPlaceType(w http.ResponseWriter, r *http.Request, placeTypeRequest *model.PlaceTypeRequest) (int, []byte, error) {
	url := getSimplePlaceTypesURL()

	payload, err := json.Marshal(placeTypeRequest)
	if err != nil {
		return 0, nil, err
	}

	if err != nil {
		return 0, nil, err
	}

	code, body, err := handlers.PostResource(w, r, url, payload)

	return code, body, nil
}

func patchPlaceType(w http.ResponseWriter, r *http.Request, placeTypeRequest *model.PlaceTypeRequest) (int, []byte, error) {
	url := getPlaceTypeURL(placeTypeRequest.ID)

	payload, err := json.Marshal(placeTypeRequest)
	if err != nil {
		return 0, nil, err
	}

	code, body, err := handlers.PatchResource(w, r, url, payload)
	if err != nil {
		return 0, nil, err
	}

	return code, body, nil
}

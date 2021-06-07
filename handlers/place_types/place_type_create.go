package place_types

import (
	"encoding/json"
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

func CreatePlaceType(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-place-types", "CreatePlaceType")

	placeRequest, err := makePlaceTypeRequest(r)
	if err != nil {
		log.FailedReturn()
		return
	}

	code, body, err := postPlaceType(w, r, placeRequest)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	if code == http.StatusUnprocessableEntity {
		responseErrors, _ := handlers.OnUpdateError(r, body, GetLabel(PlaceTypeCreateErrorsReceivedIndex))
		NewPlaceTypeRetry(w, r, placeRequest, *responseErrors)
		log.FailedReturn()
		return
	}

	if code == http.StatusCreated || code == http.StatusOK {
		var placeType *model.PlaceType
		_ = json.Unmarshal(body, &placeType)

		handlers.WriteSessionInfoMessage(r, GetLabel(PlaceTypeCreatedIndex))
		url := fmt.Sprintf("/place-types/%d", placeType.ID)
		http.Redirect(w, r, url, http.StatusMovedPermanently)
		return
	} else {
		fmt.Fprintf(w, "received %d", code)
		return
	}

	log.NormalReturn()
}

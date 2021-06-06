package place_types

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"net/http"
)

func UpdatePlaceType(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-place-types", "UpdatePlaceTypes")

	placeTypeRequest, err := makePlaceTypeRequest(r)
	if err != nil {
		log.FailedReturn()
		return
	}

	code, body, err := patchPlaceType(w, r, placeTypeRequest)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	if code == http.StatusUnprocessableEntity {
		responseErrors, _ := handlers.OnUpdateError(r, body, GetLabel(PlaceTypeUpdateErrorsReceivedIndex))
		EditPlaceTypeRetry(w, r, placeTypeRequest, *responseErrors)
		return
	}

	if code == http.StatusOK {
		handlers.WriteSessionInfoMessage(r, GetLabel(PlaceTypeUpdatedIndex))
		url := fmt.Sprintf("/place-types/%d", placeTypeRequest.ID)
		http.Redirect(w, r, url, http.StatusMovedPermanently)
		return
	} else {
		fmt.Fprintf(w, "received %d", code)
		log.FailedReturn()
		return
	}
}


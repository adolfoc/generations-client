package places

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"net/http"
)

func UpdatePlace(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-places", "UpdatePlace")

	placeRequest, err := makePlaceRequest(r)
	if err != nil {
		log.FailedReturn()
		return
	}

	code, body, err := patchPlace(w, r, placeRequest)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	if code == http.StatusUnprocessableEntity {
		responseErrors, _ := handlers.OnUpdateError(r, body, GetLabel(PlaceUpdateErrorsReceivedIndex))
		EditPlaceRetry(w, r, placeRequest, *responseErrors)
		return
	}

	if code == http.StatusOK {
		handlers.WriteSessionInfoMessage(r, GetLabel(PlaceUpdatedIndex))
		url := fmt.Sprintf("/places/%d", placeRequest.ID)
		http.Redirect(w, r, url, http.StatusMovedPermanently)
		return
	} else {
		fmt.Fprintf(w, "received %d", code)
		log.FailedReturn()
		return
	}
}


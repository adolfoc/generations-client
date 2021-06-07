package places

import (
	"encoding/json"
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

func CreatePlace(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-places", "CreatePlace")

	placeRequest, err := makePlaceRequest(r)
	if err != nil {
		log.FailedReturn()
		return
	}

	code, body, err := postPlace(w, r, placeRequest)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	if code == http.StatusUnprocessableEntity {
		responseErrors, _ := handlers.OnUpdateError(r, body, GetLabel(PlaceCreateErrorsReceivedIndex))
		NewPlaceRetry(w, r, placeRequest, *responseErrors)
		log.FailedReturn()
		return
	}

	if code == http.StatusCreated || code == http.StatusOK {
		var place *model.Place
		_ = json.Unmarshal(body, &place)

		handlers.WriteSessionInfoMessage(r, GetLabel(PlaceCreatedIndex))
		url := fmt.Sprintf("/places/%d", place.ID)
		http.Redirect(w, r, url, http.StatusMovedPermanently)
		return
	} else {
		fmt.Fprintf(w, "received %d", code)
		return
	}

	log.NormalReturn()
}

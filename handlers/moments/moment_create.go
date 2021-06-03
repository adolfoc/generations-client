package moments

import (
	"encoding/json"
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

func CreateMoment(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-moments", "CreateMoment")

	momentRequest, err := makeMomentRequest(r)
	if err != nil {
		log.FailedReturn()
		return
	}

	code, body, err := postMoment(w, r, momentRequest)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	if code == http.StatusUnprocessableEntity {
		responseErrors, _ := handlers.OnUpdateError(r, body, GetLabel(MomentCreateErrorsReceivedIndex))
		NewMomentRetry(w, r, momentRequest, *responseErrors)
		log.FailedReturn()
		return
	}

	if code == http.StatusCreated || code == http.StatusOK {
		var moment *model.HistoricalMoment
		_ = json.Unmarshal(body, &moment)

		handlers.WriteSessionInfoMessage(r, GetLabel(MomentCreatedIndex))
		url := fmt.Sprintf("/schemas/%d/moments/%d", moment.SchemaID, moment.ID)
		http.Redirect(w, r, url, http.StatusMovedPermanently)
		return
	} else {
		fmt.Fprintf(w, "received %d", code)
		return
	}
}

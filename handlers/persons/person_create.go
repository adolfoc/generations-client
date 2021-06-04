package persons

import (
	"encoding/json"
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-persons", "CreatePerson")

	eventRequest, err := makePersonRequest(r)
	if err != nil {
		log.FailedReturn()
		return
	}

	code, body, err := postPerson(w, r, eventRequest)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	if code == http.StatusUnprocessableEntity {
		responseErrors, _ := handlers.OnUpdateError(r, body, GetLabel(PersonCreateErrorsReceivedIndex))
		NewPersonRetry(w, r, eventRequest, *responseErrors)
		log.FailedReturn()
		return
	}

	if code == http.StatusCreated || code == http.StatusOK {
		var person *model.Person
		_ = json.Unmarshal(body, &person)

		handlers.WriteSessionInfoMessage(r, GetLabel(PersonCreatedIndex))
		url := fmt.Sprintf("/persons/%d", person.ID)
		http.Redirect(w, r, url, http.StatusMovedPermanently)
		return
	} else {
		fmt.Fprintf(w, "received %d", code)
		return
	}

	log.NormalReturn()
}

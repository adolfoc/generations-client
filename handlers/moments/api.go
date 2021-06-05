package moments

import (
	"encoding/json"
	"fmt"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

const (
	ResourceGenerationSchema = "generation-schemas"
	ResourceMoments          = "historical-moments"
	ResourcePositions        = "generation-positions"
	ResourceEvents           = "events"
	ResourcePersons          = "persons"
)

func getSchemaMomentsURL(generationSchemaID int) string {
	return fmt.Sprintf("%s%s/%d/moments", handlers.GetAPIHostURL(), ResourceGenerationSchema, generationSchemaID)
}

func getSchemaMomentURL(generationSchemaID, momentID int) string {
	return fmt.Sprintf("%s%s/%d", handlers.GetAPIHostURL(), ResourceMoments, momentID)
}

func getMomentsURL() string {
	return fmt.Sprintf("%s%s", handlers.GetAPIHostURL(), ResourceMoments)
}

func getGenerationURL(momentID int) string {
	return fmt.Sprintf("%s%s/%d", handlers.GetAPIHostURL(), ResourceMoments, momentID)
}

func getGenerationPositionsURL(generationSchemaID, momentID int) string {
	return fmt.Sprintf("%s%s/by-moment/%d", handlers.GetAPIHostURL(), ResourcePositions, momentID)
}

func getMomentsBetweenURL(startYear, endYear int) string {
	return fmt.Sprintf("%s%s/between/%d/%d", handlers.GetAPIHostURL(), ResourceEvents, startYear, endYear)
}

func getPeopleAliveBetweenURL(startYear, endYear int) string {
	return fmt.Sprintf("%s%s/alive-during/%d/%d", handlers.GetAPIHostURL(), ResourcePersons, startYear, endYear)
}

func getSchemaMoments(w http.ResponseWriter, r *http.Request, generationSchemaID int) (*model.HistoricalMoments, error) {
	url := getSchemaMomentsURL(generationSchemaID)
	code, body, err := handlers.GetResource(w, r, url)

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	var moments *model.HistoricalMoments
	err = json.Unmarshal(body, &moments)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return moments, nil
}

func getSchemaMoment(w http.ResponseWriter, r *http.Request, schemaID, momentID int) (*model.HistoricalMoment, error) {
	url := getSchemaMomentURL(schemaID, momentID)
	code, body, err := handlers.GetResource(w, r, url)

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	var moment *model.HistoricalMoment
	err = json.Unmarshal(body, &moment)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return moment, nil
}

func getGenerationPositions(w http.ResponseWriter, r *http.Request, schemaID, momentID int) ([]*model.GenerationPosition, error){
	url := getGenerationPositionsURL(schemaID, momentID)
	code, body, err := handlers.GetResource(w, r, url)

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	var positions []*model.GenerationPosition
	err = json.Unmarshal(body, &positions)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return positions, nil
}

func getEvents(w http.ResponseWriter, r *http.Request, moment *model.HistoricalMoment) ([]*model.Event, error){
	url := getMomentsBetweenURL(moment.StartYear(), moment.EndYear())
	code, body, err := handlers.GetResource(w, r, url)

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	var events *model.Events
	err = json.Unmarshal(body, &events)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return events.Events, nil
}

func getContemporaries(w http.ResponseWriter, r *http.Request, moment *model.HistoricalMoment) ([]*model.Person, error){
	url := getPeopleAliveBetweenURL(moment.StartYear(), moment.EndYear())
	code, body, err := handlers.GetResource(w, r, url)

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	var persons *model.Persons
	err = json.Unmarshal(body, &persons)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return persons.Persons, nil
}

func getUrlGenerationSchemaID(w http.ResponseWriter, r *http.Request) (int, error) {
	return handlers.GetUrlIntParam("schema_id", w, r)
}

func getUrlMomentID(w http.ResponseWriter, r *http.Request) (int, error) {
	return handlers.GetUrlIntParam("moment_id", w, r)
}

func buildMomentRequest(moment *model.HistoricalMoment) *model.HistoricalMomentRequest {
	mr := &model.HistoricalMomentRequest{
		ID:          moment.ID,
		Name:        moment.Name,
		SchemaID:    moment.SchemaID,
		TypeID:      moment.Type.ID,
		Start:       moment.Start,
		End:         moment.End,
		Summary:     moment.Summary,
		Description: moment.Description,
	}

	return mr
}

func newMomentRequest(schemaID int) *model.HistoricalMomentRequest {
	mtr := &model.HistoricalMomentRequest{
		SchemaID:    schemaID,
	}

	return mtr
}

func makeMomentRequest(r *http.Request) (*model.HistoricalMomentRequest, error) {
	normalizedID := handlers.GetIntFormValue(r, "inputID")
	normalizedSchemaID := handlers.GetIntFormValue(r, "inputSchemaID")
	name := handlers.GetStringFormValue(r, "inputName")
	normalizedTypeID := handlers.GetIntFormValue(r, "inputTypeID")
	start := handlers.GetStringFormValue(r, "inputStart")
	end := handlers.GetStringFormValue(r, "inputEnd")
	summary := handlers.GetStringFormValue(r, "inputSummary")
	description := handlers.GetStringFormValue(r, "inputDescription")

	mr := &model.HistoricalMomentRequest{
		ID:          normalizedID,
		Name:        name,
		SchemaID:    normalizedSchemaID,
		TypeID:      normalizedTypeID,
		Start:       start,
		End:         end,
		Summary:     summary,
		Description: description,
	}

	return mr, nil
}

func postMoment(w http.ResponseWriter, r *http.Request, momentRequest *model.HistoricalMomentRequest) (int, []byte, error) {
	url := getMomentsURL()

	payload, err := json.Marshal(momentRequest)
	if err != nil {
		return 0, nil, err
	}

	if err != nil {
		return 0, nil, err
	}

	code, body, err := handlers.PostResource(w, r, url, payload)

	return code, body, nil
}

func patchMoment(w http.ResponseWriter, r *http.Request, momentRequest *model.HistoricalMomentRequest) (int, []byte, error) {
	url := getGenerationURL(momentRequest.ID)

	payload, err := json.Marshal(momentRequest)
	if err != nil {
		return 0, nil, err
	}

	code, body, err := handlers.PatchResource(w, r, url, payload)
	if err != nil {
		return 0, nil, err
	}


	return code, body, nil
}

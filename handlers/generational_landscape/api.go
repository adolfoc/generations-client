package generational_landscape

import (
	"encoding/json"
	"fmt"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

const (
	ResourceGenerationalLandscape = "generational-landscapes"
	ResourceSchemas               = "generation-schemas"
	ResourceTangibles             = "tangibles"
	ResourceIntangibles           = "intangibles"
)

func getGenerationalLandscapesURL(generationSchemaID int) string {
	return fmt.Sprintf("%s%s", handlers.GetAPIHostURL(), ResourceGenerationalLandscape)
}

func getGenerationalLandscapeURL(generationSchemaID, generationalLandscapeID int) string {
	return fmt.Sprintf("%s%s/%d", handlers.GetAPIHostURL(), ResourceGenerationalLandscape, generationalLandscapeID)
}

func getMomentsForSchemaURL(schemaID int) string {
	return fmt.Sprintf("%s%s/%d/moments", handlers.GetAPIHostURL(), ResourceSchemas, schemaID)
}

func getAddTangibleURL(generationalLandscapeID int) string {
	return fmt.Sprintf("%s%s/%d/tangible", handlers.GetAPIHostURL(), ResourceGenerationalLandscape, generationalLandscapeID)
}

func getTangibleURL(tangibleID int) string {
	return fmt.Sprintf("%s%s/%d", handlers.GetAPIHostURL(), ResourceTangibles, tangibleID)
}

func getAddIntangibleURL(generationalLandscapeID int) string {
	return fmt.Sprintf("%s%s/%d/intangible", handlers.GetAPIHostURL(), ResourceGenerationalLandscape, generationalLandscapeID)
}

func getIntangibleURL(intangibleID int) string {
	return fmt.Sprintf("%s%s/%d", handlers.GetAPIHostURL(), ResourceIntangibles, intangibleID)
}

func getGenerationalLandscape(w http.ResponseWriter, r *http.Request, generationSchemaID, generationalLandscapeID int) (*model.GenerationalLandscape, error) {
	url := getGenerationalLandscapeURL(generationSchemaID, generationalLandscapeID)
	code, body, err := handlers.GetResource(w, r, url)

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	var generationalLandscape *model.GenerationalLandscape
	err = json.Unmarshal(body, &generationalLandscape)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return generationalLandscape, nil
}

func getUrlGenerationSchemaID(w http.ResponseWriter, r *http.Request) (int, error) {
	return handlers.GetUrlIntParam("schema_id", w, r)
}

func getUrlGenerationalLandscapeID(w http.ResponseWriter, r *http.Request) (int, error) {
	return handlers.GetUrlIntParam("generational_landscape_id", w, r)
}

func getUrlGenerationID(w http.ResponseWriter, r *http.Request) (int, error) {
	return handlers.GetUrlIntParam("generation_id", w, r)
}

func getMomentsForSchema(w http.ResponseWriter, r *http.Request, schemaID int) (*model.HistoricalMoments, error) {
	url := getMomentsForSchemaURL(schemaID)
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

func buildGenerationalLandscapeRequest(generationalLandscape *model.GenerationalLandscape) *model.GenerationalLandscapeRequest {
	glr := &model.GenerationalLandscapeRequest{
		ID:                generationalLandscape.ID,
		GenerationID:      generationalLandscape.GenerationID,
		FormationMomentID: generationalLandscape.FormationMomentID,
		Description:       generationalLandscape.Description,
	}

	return glr
}

func newGenerationalLandscapeRequest(generationID int) *model.GenerationalLandscapeRequest {
	glr := &model.GenerationalLandscapeRequest{
		GenerationID:      generationID,
	}

	return glr
}

func getTangible(w http.ResponseWriter, r *http.Request, tangibleID int) (*model.Tangible, error) {
	url := getTangibleURL(tangibleID)
	code, body, err := handlers.GetResource(w, r, url)

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	var tangible *model.Tangible
	err = json.Unmarshal(body, &tangible)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return tangible, nil
}

func getUrlTangibleID(w http.ResponseWriter, r *http.Request) (int, error) {
	return handlers.GetUrlIntParam("tangible_id", w, r)
}

func makeGenerationalLandscapeRequest(r *http.Request) (*model.GenerationalLandscapeRequest, error) {
	normalizedID := handlers.GetIntFormValue(r, "inputID")
	normalizedGenerationID := handlers.GetIntFormValue(r, "inputGenerationID")
	normalizedFormationMomentID := handlers.GetIntFormValue(r, "inputFormationMomentID")
	description := handlers.GetStringFormValue(r, "inputDescription")

	glr := &model.GenerationalLandscapeRequest{
		ID:                normalizedID,
		GenerationID:      normalizedGenerationID,
		FormationMomentID: normalizedFormationMomentID,
		Description:       description,
	}

	return glr, nil
}

func buildTangibleRequest(tangible *model.Tangible) *model.TangibleRequest {
	tr := &model.TangibleRequest{
		ID:          tangible.ID,
		LandscapeID: tangible.LandscapeID,
		Name:        tangible.Name,
		Description: tangible.Description,
	}

	return tr
}

func postGenerationalLandscape(w http.ResponseWriter, r *http.Request, glRequest *model.GenerationalLandscapeRequest) (int, []byte, error) {
	url := getGenerationalLandscapesURL(0)

	payload, err := json.Marshal(glRequest)
	if err != nil {
		return 0, nil, err
	}

	if err != nil {
		return 0, nil, err
	}

	code, body, err := handlers.PostResource(w, r, url, payload)

	return code, body, nil
}

func patchGenerationalLandscape(w http.ResponseWriter, r *http.Request, auRequest *model.GenerationalLandscapeRequest) (int, []byte, error) {
	url := getGenerationalLandscapeURL(0, auRequest.ID)

	payload, err := json.Marshal(auRequest)
	if err != nil {
		return 0, nil, err
	}

	code, body, err := handlers.PatchResource(w, r, url, payload)
	if err != nil {
		return 0, nil, err
	}


	return code, body, nil
}

func newTangibleRequest(landscapeID int) *model.TangibleRequest {
	tr := &model.TangibleRequest{
		LandscapeID: landscapeID,
	}

	return tr
}

func makeTangibleRequest(r *http.Request) (*model.TangibleRequest, error) {
	normalizedID := handlers.GetIntFormValue(r, "inputID")
	normalizedLandscapeID := handlers.GetIntFormValue(r, "inputLandscapeID")
	name := handlers.GetStringFormValue(r, "inputName")
	description := handlers.GetStringFormValue(r, "inputDescription")

	glr := &model.TangibleRequest{
		ID:          normalizedID,
		LandscapeID: normalizedLandscapeID,
		Name:        name,
		Description: description,
	}

	return glr, nil
}

func postTangible(w http.ResponseWriter, r *http.Request, tRequest *model.TangibleRequest) (int, []byte, error) {
	url := getAddTangibleURL(tRequest.LandscapeID)

	payload, err := json.Marshal(tRequest)
	if err != nil {
		return 0, nil, err
	}

	if err != nil {
		return 0, nil, err
	}

	code, body, err := handlers.PostResource(w, r, url, payload)

	return code, body, nil
}

func patchTangible(w http.ResponseWriter, r *http.Request, tRequest *model.TangibleRequest) (int, []byte, error) {
	url := getTangibleURL(tRequest.ID)

	payload, err := json.Marshal(tRequest)
	if err != nil {
		return 0, nil, err
	}

	code, body, err := handlers.PatchResource(w, r, url, payload)
	if err != nil {
		return 0, nil, err
	}


	return code, body, nil
}

func getIntangible(w http.ResponseWriter, r *http.Request, intangibleID int) (*model.Intangible, error) {
	url := getTangibleURL(intangibleID)
	code, body, err := handlers.GetResource(w, r, url)

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	var intangible *model.Intangible
	err = json.Unmarshal(body, &intangible)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return intangible, nil
}

func getUrlIntangibleID(w http.ResponseWriter, r *http.Request) (int, error) {
	return handlers.GetUrlIntParam("intangible_id", w, r)
}

func makeIntangibleRequest(r *http.Request) (*model.IntangibleRequest, error) {
	normalizedID := handlers.GetIntFormValue(r, "inputID")
	normalizedLandscapeID := handlers.GetIntFormValue(r, "inputLandscapeID")
	name := handlers.GetStringFormValue(r, "inputName")
	description := handlers.GetStringFormValue(r, "inputDescription")

	itr := &model.IntangibleRequest{
		ID:          normalizedID,
		LandscapeID: normalizedLandscapeID,
		Name:        name,
		Description: description,
	}

	return itr, nil
}

func buildIntangibleRequest(intangible *model.Intangible) *model.IntangibleRequest {
	tr := &model.IntangibleRequest{
		ID:          intangible.ID,
		LandscapeID: intangible.LandscapeID,
		Name:        intangible.Name,
		Description: intangible.Description,
	}

	return tr
}

func newIntangibleRequest(landscapeID int) *model.IntangibleRequest {
	glr := &model.IntangibleRequest{
		LandscapeID: landscapeID,
	}

	return glr
}

func postIntangible(w http.ResponseWriter, r *http.Request, itRequest *model.IntangibleRequest) (int, []byte, error) {
	url := getAddIntangibleURL(itRequest.LandscapeID)

	payload, err := json.Marshal(itRequest)
	if err != nil {
		return 0, nil, err
	}

	if err != nil {
		return 0, nil, err
	}

	code, body, err := handlers.PostResource(w, r, url, payload)

	return code, body, nil
}

func patchIntangible(w http.ResponseWriter, r *http.Request, itRequest *model.IntangibleRequest) (int, []byte, error) {
	url := getIntangibleURL(itRequest.ID)

	payload, err := json.Marshal(itRequest)
	if err != nil {
		return 0, nil, err
	}

	code, body, err := handlers.PatchResource(w, r, url, payload)
	if err != nil {
		return 0, nil, err
	}


	return code, body, nil
}


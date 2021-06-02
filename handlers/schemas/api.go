package schemas

import (
	"encoding/json"
	"fmt"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

const (
	ResourceGenerationSchema = "generation-schemas"
)

func getGenerationSchemasURL() string {
	return fmt.Sprintf("%s%s", handlers.GetAPIHostURL(), ResourceGenerationSchema)
}

func getGenerationSchemaURL(generationSchemaID int) string {
	return fmt.Sprintf("%s%s/%d", handlers.GetAPIHostURL(), ResourceGenerationSchema, generationSchemaID)
}

func getGenerationSchemaLifePhasesURL(generationSchemaID int) string {
	return fmt.Sprintf("%s%s/%d/life-phases", handlers.GetAPIHostURL(), ResourceGenerationSchema, generationSchemaID)
}

func getGenerationSchemaGenerationTypesURL(generationSchemaID int) string {
	return fmt.Sprintf("%s%s/%d/generation-types", handlers.GetAPIHostURL(), ResourceGenerationSchema, generationSchemaID)
}

func getGenerationSchemaMomentTypesURL(generationSchemaID int) string {
	return fmt.Sprintf("%s%s/%d/moment-types", handlers.GetAPIHostURL(), ResourceGenerationSchema, generationSchemaID)
}

func getGenerationSchemas(w http.ResponseWriter, r *http.Request) (*model.GenerationSchemas, error) {
	url := getGenerationSchemasURL()
	code, body, err := handlers.GetResource(w, r, url)

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	var generationSchemas *model.GenerationSchemas
	err = json.Unmarshal(body, &generationSchemas)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return generationSchemas, nil
}

func getGenerationSchema(w http.ResponseWriter, r *http.Request, gsID int) (*model.GenerationSchema, error) {
	url := getGenerationSchemaURL(gsID)
	code, body, err := handlers.GetResource(w, r, url)

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	var generationSchema *model.GenerationSchema
	err = json.Unmarshal(body, &generationSchema)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return generationSchema, nil
}

func getUrlGenerationSchemaID(w http.ResponseWriter, r *http.Request) (int, error) {
	return handlers.GetUrlIntParam("schema_id", w, r)
}

func getLifePhases(w http.ResponseWriter, r *http.Request, gsID int) ([]*model.LifePhase, error) {
	url := getGenerationSchemaLifePhasesURL(gsID)
	code, body, err := handlers.GetResource(w, r, url)

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	var lifePhases []*model.LifePhase
	err = json.Unmarshal(body, &lifePhases)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return lifePhases, nil
}

func getGenerationTypes(w http.ResponseWriter, r *http.Request, gsID int) ([]*model.GenerationType, error) {
	url := getGenerationSchemaGenerationTypesURL(gsID)
	code, body, err := handlers.GetResource(w, r, url)

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	var generationTypes []*model.GenerationType
	err = json.Unmarshal(body, &generationTypes)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return generationTypes, nil
}

func getMomentTypes(w http.ResponseWriter, r *http.Request, gsID int) ([]*model.MomentType, error) {
	url := getGenerationSchemaMomentTypesURL(gsID)
	code, body, err := handlers.GetResource(w, r, url)

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	var momentTypes []*model.MomentType
	err = json.Unmarshal(body, &momentTypes)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return momentTypes, nil
}

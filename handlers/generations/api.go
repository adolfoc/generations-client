package generations

import (
	"encoding/json"
	"fmt"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

const (
	ResourceGenerationSchema = "generation-schemas"
	ResourceGenerations = "generations"
)

func getSchemaGenerationsURL(generationSchemaID int) string {
	return fmt.Sprintf("%s%s/%d/generations", handlers.GetAPIHostURL(), ResourceGenerationSchema, generationSchemaID)
}

func getSchemaGenerationURL(generationSchemaID, generationID int) string {
	return fmt.Sprintf("%s%s/%d", handlers.GetAPIHostURL(), ResourceGenerations, generationID)
}

func getSchemaGenerations(w http.ResponseWriter, r *http.Request, generationSchemaID int) (*model.Generations, error) {
	url := getSchemaGenerationsURL(generationSchemaID)
	code, body, err := handlers.GetResource(w, r, url)

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	var generations *model.Generations
	err = json.Unmarshal(body, &generations)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return generations, nil
}

func getSchemaGeneration(w http.ResponseWriter, r *http.Request, generationSchemaID, generationID int) (*model.Generation, error) {
	url := getSchemaGenerationURL(generationSchemaID, generationID)
	code, body, err := handlers.GetResource(w, r, url)

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	var generation *model.Generation
	err = json.Unmarshal(body, &generation)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return generation, nil
}

func getUrlGenerationSchemaID(w http.ResponseWriter, r *http.Request) (int, error) {
	return handlers.GetUrlIntParam("schema_id", w, r)
}

func getUrlGenerationID(w http.ResponseWriter, r *http.Request) (int, error) {
	return handlers.GetUrlIntParam("generation_id", w, r)
}


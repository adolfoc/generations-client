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
)

func getSchemaMomentsURL(generationSchemaID int) string {
	return fmt.Sprintf("%s%s/%d/moments", handlers.GetAPIHostURL(), ResourceGenerationSchema, generationSchemaID)
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

func getUrlGenerationSchemaID(w http.ResponseWriter, r *http.Request) (int, error) {
	return handlers.GetUrlIntParam("schema_id", w, r)
}



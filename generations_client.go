package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/handlers/authentication"
	"github.com/adolfoc/generations-client/handlers/events"
	"github.com/adolfoc/generations-client/handlers/generations"
	"github.com/adolfoc/generations-client/handlers/life_segments"
	"github.com/adolfoc/generations-client/handlers/moments"
	"github.com/adolfoc/generations-client/handlers/persons"
	"github.com/adolfoc/generations-client/handlers/place_types"
	"github.com/adolfoc/generations-client/handlers/places"
	"github.com/adolfoc/generations-client/handlers/schemas"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var EnvironmentLoaded bool = false

func loadEnvironment() bool {
	if EnvironmentLoaded == true {
		return true
	}

	root, err := os.Getwd()
	fullPath := fmt.Sprintf("%s/.env", root)
	err = godotenv.Load(fullPath)
	if err != nil {
		fmt.Printf("error loading environment %q\n", err.Error())
		return false
	}

	EnvironmentLoaded = true
	return true
}

func getServerSpec() string {
	if loadEnvironment() == false {
		return ""
	}

	host := os.Getenv("GENERATIONS_CLIENT_HOST")
	port := os.Getenv("GENERATIONS_CLIENT_PORT")
	return fmt.Sprintf("%s:%s", host, port)
}

func useTLS() bool {
	if loadEnvironment() == false {
		return false
	}

	tls := os.Getenv("GENERATIONS_USE_TLS")
	if len(tls) > 0 && tls == "1" {
		return true
	}

	return false
}

func getHTTPSCredentials() (string, string) {
	if loadEnvironment() == false {
		return "", ""
	}

	certFile := os.Getenv("GENERATIONS_CLIENT_CERT_FILE")
	keyFile := os.Getenv("GENERATIONS_CLIENT_KEY_FILE")
	return certFile, keyFile
}

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second * 15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	r := makeRouter()

	hostParams := getServerSpec()
	fmt.Printf("Server will listen on %s\n", hostParams)
	srv := &http.Server{
		Addr:         hostParams,
		// Good practice to set timeouts to avoid Slowloris attacks.
		//WriteTimeout: time.Second * 15,
		//ReadTimeout:  time.Second * 15,
		//IdleTimeout:  time.Second * 60,
		Handler: r, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if useTLS() == true {
			cert, key := getHTTPSCredentials()
			fmt.Printf("Using cert %q and key %q\n", cert, key)
			if err := srv.ListenAndServeTLS(cert, key); err != nil {
				log.Println(err)
			}
		} else {
			if err := srv.ListenAndServe(); err != nil {
				log.Println(err)
			}
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}

func makeRouter() *mux.Router {
	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	r.HandleFunc("/", authentication.Authenticate).Methods("GET")
	r.HandleFunc("/request-authentication", authentication.PerformAuthentication).Methods("POST")
	r.HandleFunc("/authenticate", authentication.Authenticate).Methods("POST")
	r.HandleFunc("/logout", authentication.Logout).Methods("GET")
	r.HandleFunc("/session-expired", handlers.SessionExpiredHandler).Methods("GET")
	r.HandleFunc("/general-error", handlers.GetErrorPage).Methods("GET")

	r.HandleFunc("/schemas/index", schemas.GetGenerationSchemas).Methods("GET")
	r.HandleFunc("/schemas/{schema_id:[0-9]+}", schemas.GetGenerationSchema).Methods("GET")
	r.HandleFunc("/schemas/{schema_id:[0-9]+}/generate-template", schemas.GenerateTemplate).Methods("GET")

	r.HandleFunc("/schemas/{schema_id:[0-9]+}/generations", generations.GetSchemaGenerations).Methods("GET")
	r.HandleFunc("/schemas/{schema_id:[0-9]+}/generations/{generation_id:[0-9]+}", generations.GetGeneration).Methods("GET")
	r.HandleFunc("/schemas/{schema_id:[0-9]+}/generations/{generation_id:[0-9]+}/edit", generations.EditGeneration).Methods("GET")
	r.HandleFunc("/schemas/{schema_id:[0-9]+}/generations/{generation_id:[0-9]+}/update", generations.UpdateGeneration).Methods("POST")
	r.HandleFunc("/schemas/{schema_id:[0-9]+}/generations/new", generations.NewGeneration).Methods("GET")
	r.HandleFunc("/schemas/{schema_id:[0-9]+}/generations/create", generations.CreateGeneration).Methods("POST")

	r.HandleFunc("/schemas/{schema_id:[0-9]+}/moments", moments.GetSchemaMoments).Methods("GET")
	r.HandleFunc("/schemas/{schema_id:[0-9]+}/moments/{moment_id:[0-9]+}", moments.GetMoment).Methods("GET")
	r.HandleFunc("/schemas/{schema_id:[0-9]+}/moments/{moment_id:[0-9]+}/edit", moments.EditMoment).Methods("GET")
	r.HandleFunc("/schemas/{schema_id:[0-9]+}/moments/{moment_id:[0-9]+}/update", moments.UpdateMoment).Methods("POST")
	r.HandleFunc("/schemas/{schema_id:[0-9]+}/moments/new", moments.NewMoment).Methods("GET")
	r.HandleFunc("/schemas/{schema_id:[0-9]+}/moments/create", moments.CreateMoment).Methods("POST")

	r.HandleFunc("/persons/index", persons.GetPersons).Methods("GET")
	r.HandleFunc("/persons/{person_id:[0-9]+}", persons.GetPerson).Methods("GET")
	r.HandleFunc("/persons/new-person", persons.NewPerson).Methods("GET")
	r.HandleFunc("/persons/{person_id:[0-9]+}/create", persons.CreatePerson).Methods("POST")
	r.HandleFunc("/persons/{person_id:[0-9]+}/edit", persons.EditPerson).Methods("GET")
	r.HandleFunc("/persons/{person_id:[0-9]+}/update", persons.UpdatePerson).Methods("POST")
	r.HandleFunc("/persons/{person_id:[0-9]+}/add-life-segments", persons.AddSegments).Methods("GET")
	r.HandleFunc("/persons/{person_id:[0-9]+}/generate-life-segments", persons.GenerateLifeSegments).Methods("POST")

	r.HandleFunc("/persons/{person_id:[0-9]+}/life-segments/{life_segment_id:[0-9]+}/edit", life_segments.EditLifeSegment).Methods("GET")
	r.HandleFunc("/persons/{person_id:[0-9]+}/life-segments/{life_segment_id:[0-9]+}/update", life_segments.UpdateLifeSegment).Methods("POST")

	r.HandleFunc("/events/index", events.GetEvents).Methods("GET")
	r.HandleFunc("/events/{event_id:[0-9]+}", events.GetEvent).Methods("GET")
	r.HandleFunc("/events/new", events.NewEvent).Methods("GET")
	r.HandleFunc("/events/create", events.CreateEvent).Methods("POST")
	r.HandleFunc("/events/{event_id:[0-9]+}/edit", events.EditEvent).Methods("GET")
	r.HandleFunc("/events/{event_id:[0-9]+}/update", events.UpdateEvent).Methods("POST")

	r.HandleFunc("/places/index", places.GetPlaces).Methods("GET")
	r.HandleFunc("/places/{place_id:[0-9]+}", places.GetPlace).Methods("GET")
	r.HandleFunc("/places/new", places.NewPlace).Methods("GET")
	r.HandleFunc("/places/create", places.CreatePlace).Methods("POST")
	r.HandleFunc("/places/{place_id:[0-9]+}/edit", places.EditPlace).Methods("GET")
	r.HandleFunc("/places/{place_id:[0-9]+}/update", places.UpdatePlace).Methods("POST")

	r.HandleFunc("/place-types/index", place_types.GetPlaceTypes).Methods("GET")
	r.HandleFunc("/place-types/{place_type_id:[0-9]+}", place_types.GetPlaceType).Methods("GET")
	//r.HandleFunc("/places/new", place_types.NewPlaceType).Methods("GET")
	//r.HandleFunc("/places/create", place_types.CreatePlaceType).Methods("POST")
	r.HandleFunc("/place-types/{place_type_id:[0-9]+}/edit", place_types.EditPlaceType).Methods("GET")
	r.HandleFunc("/place-types/{place_type_id:[0-9]+}/update", place_types.UpdatePlaceType).Methods("POST")

	return r
}


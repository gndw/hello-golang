package main

import "hello-golang/app"

// "hello-golang/app"
// "hello-golang/app/master"
// "hello-golang/v1/functions/startup"
// "hello-golang/v1/repositories/http_gorilla_mux"


func main() {

	// Creating the App Instance
	app := &app.App{}
	app.New()

	// Initialization Process / Startup
	
	
	
	
	
	// startup := &startup.Startup{}
	// startup.Register(mux)
	// startup.Start()

	// r := mux.NewRouter()

	// r.HandleFunc("/health", handlers.HealthHandler).Methods("GET")
	// r.HandleFunc("/auth/login", handlers.LoginHandler).Methods("POST")
	// r.HandleFunc("/auth/register", handlers.RegisterHandler).Methods("POST")

	// dataroute := r.PathPrefix("/data").Subrouter()
	// dataroute.Use(middlewares.AuthorizationMiddleware)
	// dataroute.HandleFunc("/self", handlers.DataSelfHandler).Methods("POST")

	// r.Use(middlewares.ContentTypeJSONMiddleware)
	// r.Use(middlewares.ReadBodyFromHTTPRequestMiddleware)

	// log.Fatal(http.ListenAndServe(":3000", r))
}

package main

func main() {

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

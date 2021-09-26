package controllers

import "github.com/Stuart6970/e-comm-api/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// CatalogItem Route
	s.Router.HandleFunc("/items", middlewares.SetMiddlewareJSON(s.CreateCatalogItem)).Methods("POST")
	s.Router.HandleFunc("/items", middlewares.SetMiddlewareJSON(s.GetCatalogItems)).Methods("GET")
	s.Router.HandleFunc("/items/{id}", middlewares.SetMiddlewareJSON(s.GetCatalogItem)).Methods("GET")
	s.Router.HandleFunc("/items/{id}", middlewares.SetMiddlewareJSON(s.UpdateCatalogItem)).Methods("PUT")
	s.Router.HandleFunc("/items/{id}", middlewares.SetMiddlewareJSON(s.DeleteCatalogItem)).Methods("DELETE")
}

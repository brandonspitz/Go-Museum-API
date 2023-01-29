package routes

import (
	"github.com/brandonspitz/GO-Museum/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterMuseumRoutes = func(router *mux.Router) {
	router.HandleFunc("/artifact/", controllers.CreateArtifact).Methods("POST")
	router.HandleFunc("/artifact/", controllers.GetArtifact).Methods("GET")
	router.HandleFunc("/artifact/{artifactId}", controllers.GetArtifactById).Methods("GET")
	router.HandleFunc("/artifact/{artifactId}", controllers.UpdateArtifact).Methods("PUT")
	router.HandleFunc("/artifact/{artifactId}", controllers.DeleteArtifact).Methods("DELETE")
}

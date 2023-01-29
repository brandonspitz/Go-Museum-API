package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/brandonspitz/GO-Museum/pkg/models"
	"github.com/brandonspitz/GO-Museum/pkg/utils"
	"github.com/gorilla/mux"
)

var NewArtifact models.Artifact

func GetArtifact(w http.ResponseWriter, r *http.Request) {
	newArtifacts := models.GetAllArtifacts()
	res, _ := json.Marshal(newArtifacts)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetArtifactById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	artifactId := vars["artifactId"]
	ID, err := strconv.ParseInt(artifactId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	artifactDetails, _ := models.GetArtifactById(ID)
	res, _ := json.Marshal(artifactDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateArtifact(w http.ResponseWriter, r *http.Request) {
	CreateArtifact := &models.Artifact{}
	utils.ParseBody(r, CreateArtifact)
	b := CreateArtifact.CreateArtifact()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteArtifact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	artifactId := vars["artifactId"]
	ID, err := strconv.ParseInt(artifactId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	artifact := models.DeleteArtifact(ID)
	res, _ := json.Marshal(artifact)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateArtifact(w http.ResponseWriter, r *http.Request) {
	var updateArtifact = &models.Artifact{}
	utils.ParseBody(r, updateArtifact)
	vars := mux.Vars(r)
	artifactId := vars["artifactId"]
	ID, err := strconv.ParseInt(artifactId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	artifactDetails, db := models.GetArtifactById(ID)
	if updateArtifact.Site != "" {
		artifactDetails.Site = updateArtifact.Site
	}
	if updateArtifact.Feature != 0 {
		artifactDetails.Feature = updateArtifact.Feature
	}
	if updateArtifact.Type != "" {
		artifactDetails.Type = updateArtifact.Type
	}
	db.Save(&artifactDetails)
	res, _ := json.Marshal(artifactDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

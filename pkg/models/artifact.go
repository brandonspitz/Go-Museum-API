package models

import (
	"github.com/brandonspitz/GO-Museum/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Artifact struct {
	gorm.Model
	Site    string `gorm: ""json:"site"`
	Feature int    `json: "feature"`
	Type    string `json: "type"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Artifact{})
}

func (b *Artifact) CreateArtifact() *Artifact {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllArtifacts() []Artifact {
	var Artifacts []Artifact
	db.Find(&Artifacts)
	return Artifacts
}

func GetArtifactById(Id int64) (*Artifact, *gorm.DB) {
	var getArtifact Artifact
	db := db.Where("ID=?", Id).Find(&getArtifact)
	return &getArtifact, db
}

func DeleteArtifact(Id int64) Artifact {
	var artifact Artifact
	db.Where("ID=?", Id).Delete(artifact)
	return artifact
}

package receta

import "gopkg.in/mgo.v2/bson"

type Receta struct {
	ID           bson.ObjectId `bson:"_id" json:":id"`
	Nombre       string        `bson:"nombre" json:"nombre"`
	Descripcion  string        `bson:"descripcion" json:"descripcion"`
	Ingredientes []string      `bson:"ingredientes" json:"ingredientes"`
	Pasos        []string      `bson:"pasos" json:"pasos"`
}

type Recetas []Receta

package receta

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Repository struct{}

const SERVER = "localhost:27017"
const DBNAME = "restapi"
const DOCNAME = "recetas"

func (r Repository) GetRecetas() Recetas {
	recetas, err := mgo.Dial(SERVER)
	if err != nil {
		fmt.Println("Error al establecer conexión con el servidor de Mongo:", err)
	}
	defer recetas.Close()
	c := recetas.DB(DBNAME).C(DOCNAME)
	results := Recetas{}
	if err := c.Find(bson.M{}).All(&results); err != nil {
		fmt.Println("Error al recibir las recetas:", err)
	}
	return results
}

func (r Repository) GetReceta(nombre string) Recetas {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		fmt.Println("Error al establecer conexión con el servidor de Mongo:", err)
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME)
	results := Recetas{}
	if err := c.Find(bson.M{"nombre": nombre}).All(&results); err != nil {
		fmt.Println("Error al recibir la receta:", err)
	}
	return results
}

func (r Repository) AddReceta(receta Receta) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()
	receta.ID = bson.NewObjectId()
	session.DB(DBNAME).C(DOCNAME).Insert(receta)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func (r Repository) UpdateReceta(receta Receta) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()
	session.DB(DBNAME).C(DOCNAME).UpdateId(receta.ID, receta)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func (r Repository) DeleteReceta(nombre string) string {
	session, err := mgo.Dial(SERVER)
	defer session.Close()
	if err = session.DB(DBNAME).C(DOCNAME).Remove(bson.M{"nombre": nombre}); err != nil {
		log.Fatal(err)
		return "INTERNAL ERR"
	}
	return "OK"
}

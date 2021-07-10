package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type Person struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

var dbConnect *pg.DB

func CreatePersonTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	err := db.CreateTable(&Person{}, opts)

	if err != nil {
		log.Printf("Error while creating todo tables, Reason: %v \n", err)
		return err
	}

	log.Println("Survey Table Created. ")
	return nil
}

func InitiateDB(db *pg.DB) {
	dbConnect = db
}

func GetAllPeople(c *gin.Context) {
	var people []Person
	err := dbConnect.Model(&people).Select()

	if err != nil {
		log.Printf("Error while getting all people, Reason: %v \n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Surveys",
		"data":    people,
	})
	return
}

func CreatePerson(c *gin.Context) {
	var person Person
	c.BindJSON(&person)
	err := dbConnect.Insert(&person)
	if err != nil {
		log.Printf("Error while inserting person, Reason: %v \n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "something went wrong",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Survey Created Successfully",
	})
	return
}

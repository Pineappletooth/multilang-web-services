package main

import (
		"net/http"

		"github.com/gin-gonic/gin"
		"pineappletooth/encryptednotes/database"
)
type NoteDB database.Connection
type NoteJson struct {
		Title string `json:"title"`
		Note string `json:"note"`
}

func (n Note) String() string {
		return n.Note + n.Title
}
func (db NoteDB) CreateNote (c *gin.Context) {
		var note NoteJson
		
    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&note); err != nil {
        return
    }
		db.DB.Create(&Note{
			Title: note.Title,
			Note: note.Note,
		})
		 c.JSON(http.StatusCreated, note)
}

func (db NoteDB) ReadNote (c *gin.Context) {
		var notes []Note
		results := db.DB.Find(&notes)
		if results.Error != nil {
		return
	}
		c.IndentedJSON(http.StatusOK, notes)	
}

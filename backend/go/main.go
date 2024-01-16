package main

import (
  "github.com/gin-gonic/gin"
  "gorm.io/gorm"
  "gorm.io/driver/sqlite"
  // "net/http"
  // "strconv"
)


func main() {
   db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
   if err != nil {
     panic("failed to connect database")
   }

  // Migrate the schema
  db.AutoMigrate(&Note{})
    noteController := NoteDB{
      DB: db,
    }
    router := gin.Default()
    router.POST("/add", noteController.CreateNote)
    router.GET("/read",noteController.ReadNote)
    router.Run("localhost:8080")
}

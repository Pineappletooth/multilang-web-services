package main

import (
  "gorm.io/gorm"
)

type Note struct {
  gorm.Model
  Userid uint
  Title string
  Note string
}

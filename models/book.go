package models

import gm "github.com/jinzhu/gorm"

type Book struct {
	gm.Model
	Author    string
	Name      string
	PageCount int
}

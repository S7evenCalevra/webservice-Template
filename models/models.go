package models

import "gorm.io/gorm"

//*
// struct assumes a question / answer data modeling for boilerplate. Note standard JSON struct tags to tell which fields correlate to our struct fields
//In the database, both the question and answer columns will be of type TEXT
//In the database, neither column should be allowed to have NULL values
//We also set the default value for each column to NULL so that we can return an error if the user does not provide their own values when they create a new Fact
//*/

type Fact struct {
	gorm.Model
	Question string `json:"question" gorm:"text;not null;default:null`
	Answer   string `json:"answer" gorm:"text;not null;default:null`
}

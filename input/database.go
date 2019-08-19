package input

import (
	df "../common"
	"github.com/jinzhu/gorm"
)


// Dummy code to show how to implement get data from DB - need to have DB before running


var db *gorm.DB
type DatabaseLoader struct {}

func (service DatabaseLoader) GetPost() ([]df.Post, error) {
	var err error
	db, err = gorm.Open("mysql", "root@tcp(localhost:3306)/sys?parseTime=True")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var posts []df.Post
	err = db.Find(&posts).Error
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (service DatabaseLoader) GetComment() ([]df.Comment, error) {
	var err error
	db, err = gorm.Open("mysql", "root@tcp(localhost:3306)/sys?parseTime=True")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var comments []df.Comment
	err = db.Find(&comments).Error
	if err != nil {
		return nil, err
	}

	return comments, nil
}

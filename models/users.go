package models

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
)

var (
	//Error not found is returned when there is nothing in database
	ErrNotFound = errors.New("models: resource not found")
)

func NewUserService(connectionInfo string) (*UserService, error) {
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		panic(err)
	}
	// defer db.Close()
	return &UserService{
		db: db,
	}, nil
}

type UserService struct {
	db *gorm.DB
}

func (us *UserService) ById(id uint) (*User, error) {
	var user User
	err := us.db.Where("id =?", id).First(&user).Error
	switch err {
	case nil:
		return &user, nil
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

//Closes the userservice db connection
func (us *UserService) Close() error {
	return us.db.Close()
}

//Destructive reset
func (us *UserService) DestructiveReset() {
	us.db.DropTableIfExists(&User{})
	us.db.AutoMigrate(&Users{})
}

type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;unique_index"`
}
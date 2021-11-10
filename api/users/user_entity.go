package users

import (
	"fmt"
	"time"

	"github.com/pranotobudi/go-simple-ecommerce/common"
	"github.com/pranotobudi/go-simple-ecommerce/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
}

type UserRepository interface {
	FreshUserMigrator()
	UserDataSeed()
	AddUser(entity User) (*User, error)
	GetUserById(id uint) (*User, error)
	GetUserByEmail(email string) (*User, error)
	GetUserByUsername(username string) (*User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository() *userRepository {
	db := database.GetDBInstance()
	return &userRepository{db}
}

func (r *userRepository) FreshUserMigrator() {
	r.db.AutoMigrate(User{})

	// Create Fresh User Table
	if (r.db.Migrator().HasTable(&User{})) {
		fmt.Println("User table exist")
		r.db.Migrator().DropTable(&User{})
		fmt.Println("Drop User table")
	}
	r.db.Migrator().CreateTable(&User{})
	fmt.Println("Create User table")

}

func (r *userRepository) UserDataSeed() {
	statement := "INSERT INTO users (username, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"

	r.db.Exec(statement, "username1", "emailusername1@gmail.com", common.GeneratePassword("username1Password"), time.Now(), time.Now())
}

func (r *userRepository) AddUser(entity User) (*User, error) {
	err := r.db.Create(&entity).Error
	if err != nil {
		return nil, err
	}
	err = r.db.First(&entity).Error
	if err != nil {
		return nil, err
	}
	fmt.Printf("INSIDE User REPOSITORY AddEntity: %+v \n", entity)
	return &entity, nil
}

func (r *userRepository) GetUserById(id uint) (*User, error) {
	var entity User
	err := r.db.First(&entity, "id=?", id).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *userRepository) GetUserByEmail(email string) (*User, error) {
	var entity User
	err := r.db.First(&entity, "email=?", email).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *userRepository) GetUserByUsername(username string) (*User, error) {
	var entity User
	err := r.db.First(&entity, "username=?", username).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

package model

import (
	"gorm.io/gorm"

	"github.com/rajasoun/gorm-client/db/v0"
)

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"not null"`
	Email string `gorm:"not null;unique"`
}

type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository returns a new UserRepository
func NewUserRepository() (*UserRepository, error) {
	db := db.GetConnection()
	err := db.AutoMigrate(&User{})
	if err != nil {
		return nil, err
	}

	return &UserRepository{db: db}, nil
}

// NewUserRepository returns a new UserRepository with connection pool
func NewUserRepositoryWithConnectionPool() (*UserRepository, error) {
	db := db.GetConnectionWithConnectionPoll()
	err := db.AutoMigrate(&User{})
	if err != nil {
		return nil, err
	}

	return &UserRepository{db: db}, nil
}

// Close closes the database connection
func (r *UserRepository) Create(user *User) error {
	result := r.db.Create(user)
	return result.Error
}

// FindAll returns all users
func (r *UserRepository) FindAll() ([]*User, error) {
	var users []*User
	result := r.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// FindByID returns all users
func (r *UserRepository) FindByID(id uint) (*User, error) {
	var user User
	result := r.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// Updates a user
func (r *UserRepository) Update(user *User) error {
	result := r.db.Save(user)
	return result.Error
}

// Delete a user
func (r *UserRepository) Delete(user *User) error {
	result := r.db.Delete(user)
	return result.Error
}

// Find total number of users
func (r *UserRepository) Count() (int64, error) {
	var count int64
	result := r.db.Model(&User{}).Count(&count)
	return count, result.Error
}

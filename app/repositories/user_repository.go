package repositories

import (
	"github.com/swiftee19/GiziNusa-BackEnd/entities"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repository *UserRepository) CreateUser(user *entities.User) (*entities.User, error) {
	if err := repository.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (repository *UserRepository) FindUserByEmail(email string) (*entities.User, error) {
	var user entities.User

	if err := repository.db.Where("email = ?", email).First(&user).Error; err != nil {
		// If no user is found, return nil with an error
		if err == gorm.ErrRecordNotFound {
			return nil, nil // No user found, returning nil
		}
		// If there's any other error, return the error
		return nil, err
	}

	// User found, returning user
	return &user, nil
}

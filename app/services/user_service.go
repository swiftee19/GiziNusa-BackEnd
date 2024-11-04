package services

import (
	"errors"

	"github.com/swiftee19/GiziNusa-BackEnd/entities"
	"github.com/swiftee19/GiziNusa-BackEnd/repositories"
	"golang.org/x/crypto/bcrypt"
)

// UserService defines the methods related to user operations
type UserService struct {
	repository *repositories.UserRepository
}

// NewUserService initializes a new user service
func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{repository: userRepository}
}

// CreateUser creates a new user with hashed password
func (userService *UserService) CreateUser(name, email, password string) (*entities.User, error) {
	// Check if user already exists by email
	user, err := userService.repository.FindUserByEmail(email)
	if err != nil {
		return nil, err // Return the error if there's an issue querying the DB
	}

	if user != nil {
		return nil, errors.New("email already in use") // Email is already registered
	}

	// Hash the password
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return nil, err
	}

	// Create the user model
	newUser := &entities.User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}

	// Save the user in the database using the repository
	createdUser, err := userService.repository.CreateUser(newUser)
	if err != nil {
		return nil, err // Return error if DB save fails
	}

	// Return the newly created user
	return createdUser, nil
}

// HashPassword hashes the plain text password.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash compares plain text password with the hashed password.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

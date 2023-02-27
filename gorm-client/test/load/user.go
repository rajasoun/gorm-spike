package load

import (
	"math/rand"

	"github.com/icrowley/fake"
	"github.com/rajasoun/gorm-client/model"
)

func randString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

type UserRepository struct {
	repo *model.UserRepository
}

// Creates a New User Repository
func NewUserRepository() (UserRepository, error) {
	repo, err := model.NewUserRepository()
	if err != nil {
		return UserRepository{}, err
	}
	return UserRepository{repo: repo}, nil
}

// Creates a New User Repository with connection pool
func NewUserRepositoryWithConnectionPool() (UserRepository, error) {
	repo, err := model.NewUserRepositoryWithConnectionPool()
	if err != nil {
		return UserRepository{}, err
	}
	return UserRepository{repo: repo}, nil
}

// Creates a new user in the database.
func (r *UserRepository) NewUser() error {
	user := &model.User{
		Name:  fake.FullName(),
		Email: randString(10) + fake.EmailAddress(),
	}
	err := r.repo.Create(user)
	if err != nil {
		return err
	}
	return nil
}

// Concurrently creates users in the database.
func (r *UserRepository) CreateUsers(numGoroutines int, numUsersToCreatePerGoroutine int) error {
	// Create a channel to communicate the result of user creation
	resultChan := make(chan error)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			for j := 0; j < numUsersToCreatePerGoroutine; j++ {
				err := r.NewUser()
				if err != nil {
					resultChan <- err
					return
				}
			}
			resultChan <- nil
		}()
	}
	// Wait for all goroutines to finish
	for i := 0; i < numGoroutines; i++ {
		err := <-resultChan
		if err != nil {
			return err
		}
	}
	return nil
}

// Get the total number of users in the database
func (r *UserRepository) GetTotalUsers() int64 {
	users, err := r.repo.Count()
	if err != nil {
		return 0
	}
	return users
}

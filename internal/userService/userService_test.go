package userService

import (
	"MarketPlace_Pet/internal/models"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateNewUser(t *testing.T) {
	tests := []struct {
		name      string
		input     models.User
		mockSetup func(m *MockUserRepository, input models.User)
		wantErr   bool
	}{
		{
			name:  "успешное создание",
			input: models.User{Email: "test@example.com", Name: "Test User", Password: "password123"},
			mockSetup: func(m *MockUserRepository, input models.User) {
				m.On("CreateNewUser", mock.MatchedBy(func(u models.User) bool {
					return u.Email == input.Email && u.Name == input.Name
				})).Return(input, nil)
			},
			wantErr: false,
		},
		{
			name:  "ошибка создания",
			input: models.User{Email: "bad@example.com", Name: "Bad User", Password: "badpassword"},
			mockSetup: func(m *MockUserRepository, input models.User) {
				m.On("CreateNewUser", mock.AnythingOfType("models.User")).Return(models.User{}, errors.New("email already exists"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockUserRepository)
			tt.mockSetup(mockRepo, tt.input)

			service := NewUserService(mockRepo, nil)
			_, err := service.CreateNewUser(tt.input.Email, tt.input.Name, "", tt.input.Password)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser_returnValidUser(t *testing.T) {
	// Arrange
	// Act
	user, err := NewUser("João da Silva", "joao@email.com", "123456")

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.UserID)
	assert.NotEmpty(t, user.Password)
}

func TestUser_ValidatetPassword(t *testing.T) {
	// Arrage
	user, err := NewUser("João da Silva", "joao@email.com", "123456")

	// Act

	// Assert
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("1234567"))
}

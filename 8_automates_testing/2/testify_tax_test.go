package tax

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	tax, err := CalculateTax(1000)
	assert.Equal(t, 10.0, tax)
	assert.Error(t, err, "amount must be greater than 0")

}

func TestCalculateTaxAndSave(t *testing.T) {
	// arrange
	repository := &TaxRepositoryMock{}
	repository.On("SaveTax", 10.0).Return(nil)
	repository.On("SaveTax", 0.0).Return(errors.New("error saving tax"))

	// act
	err := CalculateTaxAndSave(1000, repository)

	// assert
	assert.Nil(t, err)

	// act 2
	err = CalculateTaxAndSave(0.0, repository)

	// assert 2
	assert.Error(t, err, "error saving tax")

	repository.AssertExpectations(t)
}

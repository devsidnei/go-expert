package tax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTaxAndSave(t *testing.T) {
	// Crie um mock da interface Repository
	repo := new(TaxRepositoryMock)

	// Expectativa: SaveTax deve retornar erro nil para amount 1000
	repo.On("SaveTax", 10.0).Return(nil)

	// Chame a função CalculateTaxAndSave com o mock para cada cenário
	err := CalculateTaxAndSave(1000, repo)
	assert.NoError(t, err)

	// Verifique se as expectativas foram atendidas
	repo.AssertExpectations(t)
}

func TestCalculateTax(t *testing.T) {
	// Test when amount is less than or equal to 0 with error message
	t.Run("AmountLessThanOrEqualZeroWithErrorMessage", func(t *testing.T) {
		tax, err := CalculateTax(0)
		assert.Error(t, err)
		assert.Equal(t, "amount must be greater than 0", err.Error())
		assert.Equal(t, 0.0, tax)
	})

	// Test when amount is between 0 and 1000
	t.Run("AmountBetween0And1000", func(t *testing.T) {
		tax, err := CalculateTax(500)
		assert.NoError(t, err)
		assert.Equal(t, 5.0, tax)
	})

	// Test when amount is equal to 1000
	t.Run("AmountEqualTo1000", func(t *testing.T) {
		tax, err := CalculateTax(1000)
		assert.NoError(t, err)
		assert.Equal(t, 10.0, tax)
	})

	// Test when amount is between 1000 and 20000
	t.Run("AmountBetween1000And20000", func(t *testing.T) {
		tax, err := CalculateTax(15000)
		assert.NoError(t, err)
		assert.Equal(t, 10.0, tax)
	})

	// Test when amount is equal to 20000
	t.Run("AmountEqualTo20000", func(t *testing.T) {
		tax, err := CalculateTax(20000)
		assert.NoError(t, err)
		assert.Equal(t, 10.0, tax)
	})

	// Test when amount is greater than 20000
	t.Run("AmountGreaterThan20000", func(t *testing.T) {
		tax, err := CalculateTax(25000)
		assert.NoError(t, err)
		assert.Equal(t, 20.0, tax)
	})

	// Test when amount is negative with error message
	t.Run("NegativeAmountWithErrorMessage", func(t *testing.T) {
		tax, err := CalculateTax(-1000)
		assert.Error(t, err)
		assert.Equal(t, "amount must be greater than 0", err.Error())
		assert.Equal(t, 0.0, tax)
	})
}

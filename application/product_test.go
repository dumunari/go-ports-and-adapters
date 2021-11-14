package application_test

import (
	"github.com/dumunari/go-ports-and-adapters/application"
	"github.com/stretchr/testify/require"
	uuid "github.com/satori/go.uuid"
	"testing"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Drumsticks"
	product.Status = application.DISABLED
	product.Price = 35

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0

	err = product.Enable()
	require.Equal(t, "product price must be greather than 0", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Drumsticks"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 35

	err = product.Disable()
	require.Equal(t, "product price must be 0 in order to have the product disabled", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Drumsticks"
	product.Status = application.DISABLED
	product.Price = 35

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "product status must be enabled or disabled", err.Error())

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "product price must be greater or equal than 0", err.Error())
}


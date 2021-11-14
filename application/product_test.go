package application_test

import (
	"github.com/dumunari/go-ports-and-adapters/application"
	"github.com/stretchr/testify/require"
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


package application_test

import (
	"github.com/dumunari/go-ports-and-adapters/application"
	mock_application "github.com/dumunari/go-ports-and-adapters/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTestID := "testeID"

	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Get(mockTestID).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Get(mockTestID)
	require.Nil(t, err)
	require.Equal(t, product, result)
}



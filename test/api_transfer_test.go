/*
FiveWest Client API

Testing TransferAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package fivewestapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/pbreedt/fivewestapi"
)

func Test_fivewestapi_TransferAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TransferAPIService GetAllTransfersApiV1TransferGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TransferAPI.GetAllTransfersApiV1TransferGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

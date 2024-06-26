/*
FiveWest Client API

Testing ChargeAPIService

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

func Test_fivewestapi_ChargeAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ChargeAPIService CreateApiV1ChargePost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ChargeAPI.CreateApiV1ChargePost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChargeAPIService DeactivateApiV1ChargeIdPatch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id Id1

		httpRes, err := apiClient.ChargeAPI.DeactivateApiV1ChargeIdPatch(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChargeAPIService GetAllApiV1ChargeGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ChargeAPI.GetAllApiV1ChargeGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChargeAPIService GetOneApiV1ChargeIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id Id

		resp, httpRes, err := apiClient.ChargeAPI.GetOneApiV1ChargeIdGet(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

/*
FiveWest Client API

Testing CurrencyAPIService

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

func Test_fivewestapi_CurrencyAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CurrencyAPIService GetAllApiV1CurrencyGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CurrencyAPI.GetAllApiV1CurrencyGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CurrencyAPIService GetAllSupportedCurrencyInformationApiV1CurrencyGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CurrencyAPI.GetAllSupportedCurrencyInformationApiV1CurrencyGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CurrencyAPIService GetAllSupportedSymbolsApiV1CurrencySymbolsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CurrencyAPI.GetAllSupportedSymbolsApiV1CurrencySymbolsGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CurrencyAPIService GetApiV1CurrencyIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.CurrencyAPI.GetApiV1CurrencyIdGet(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CurrencyAPIService SymbolsApiV1CurrencySymbolsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CurrencyAPI.SymbolsApiV1CurrencySymbolsGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
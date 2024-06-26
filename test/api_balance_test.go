/*
FiveWest Client API

Testing BalanceAPIService

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

func Test_fivewestapi_BalanceAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test BalanceAPIService GetAllWalletBalancesApiV1BalanceGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BalanceAPI.GetAllWalletBalancesApiV1BalanceGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BalanceAPIService GetBalanceForSymbolApiV1BalanceSymbolGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var symbol string

		resp, httpRes, err := apiClient.BalanceAPI.GetBalanceForSymbolApiV1BalanceSymbolGet(context.Background(), symbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

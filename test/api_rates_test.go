/*
FiveWest Client API

Testing RatesAPIService

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

func Test_fivewestapi_RatesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RatesAPIService ReadLatestApiV1RatesLatestGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RatesAPI.ReadLatestApiV1RatesLatestGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RatesAPIService StreamLatestApiV1RatesStreamLatestGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RatesAPI.StreamLatestApiV1RatesStreamLatestGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
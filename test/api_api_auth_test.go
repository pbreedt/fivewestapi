/*
FiveWest Client API

Testing APIAuthAPIService

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

func Test_fivewestapi_APIAuthAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test APIAuthAPIService LoginApiV1AuthTokenPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.APIAuthAPI.LoginApiV1AuthTokenPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
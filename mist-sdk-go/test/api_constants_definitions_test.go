/*
Mist API

Testing ConstantsDefinitionsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package mistsdkgo

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/tmunzer/mistsdkgo"
)

func Test_mistsdkgo_ConstantsDefinitionsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ConstantsDefinitionsAPIService ListAlarmDefinitions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ConstantsDefinitionsAPI.ListAlarmDefinitions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConstantsDefinitionsAPIService ListApLedDefinition", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ConstantsDefinitionsAPI.ListApLedDefinition(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConstantsDefinitionsAPIService ListAppCategoryDefinitions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ConstantsDefinitionsAPI.ListAppCategoryDefinitions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConstantsDefinitionsAPIService ListAppSubCategoryDefinitions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ConstantsDefinitionsAPI.ListAppSubCategoryDefinitions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
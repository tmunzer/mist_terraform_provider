/*
Mist API

Testing MSPsAPIService

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

func Test_mistsdkgo_MSPsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MSPsAPIService CreateMsp", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MSPsAPI.CreateMsp(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MSPsAPIService DeleteMsp", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var mspId string

		httpRes, err := apiClient.MSPsAPI.DeleteMsp(context.Background(), mspId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MSPsAPIService GetMspDetails", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var mspId string

		resp, httpRes, err := apiClient.MSPsAPI.GetMspDetails(context.Background(), mspId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MSPsAPIService SearchMspOrgGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var mspId string

		resp, httpRes, err := apiClient.MSPsAPI.SearchMspOrgGroup(context.Background(), mspId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MSPsAPIService UpdateMsp", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var mspId string

		resp, httpRes, err := apiClient.MSPsAPI.UpdateMsp(context.Background(), mspId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
/*
Mist API

Testing OrgsNetworkTemplatesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package mistsdkgo

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_mistsdkgo_OrgsNetworkTemplatesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OrgsNetworkTemplatesAPIService CreateOrgNetworkTemplate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsNetworkTemplatesAPI.CreateOrgNetworkTemplate(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsNetworkTemplatesAPIService DeleteOrgNetworkTemplate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var networktemplateId string

		httpRes, err := apiClient.OrgsNetworkTemplatesAPI.DeleteOrgNetworkTemplate(context.Background(), orgId, networktemplateId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsNetworkTemplatesAPIService GetOrgNetworkTemplate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var networktemplateId string

		resp, httpRes, err := apiClient.OrgsNetworkTemplatesAPI.GetOrgNetworkTemplate(context.Background(), orgId, networktemplateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsNetworkTemplatesAPIService ListOrgNetworkTemplates", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsNetworkTemplatesAPI.ListOrgNetworkTemplates(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsNetworkTemplatesAPIService UpdateOrgNetworkTemplates", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var networktemplateId string

		resp, httpRes, err := apiClient.OrgsNetworkTemplatesAPI.UpdateOrgNetworkTemplates(context.Background(), orgId, networktemplateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
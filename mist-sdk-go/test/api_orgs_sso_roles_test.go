/*
Mist API

Testing OrgsSSORolesAPIService

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

func Test_mistsdkgo_OrgsSSORolesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OrgsSSORolesAPIService CreateOrgSsoRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsSSORolesAPI.CreateOrgSsoRole(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsSSORolesAPIService DeleteOrgSsoRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var ssoroleId string

		httpRes, err := apiClient.OrgsSSORolesAPI.DeleteOrgSsoRole(context.Background(), orgId, ssoroleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsSSORolesAPIService GetOrgSsoRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var ssoroleId string

		resp, httpRes, err := apiClient.OrgsSSORolesAPI.GetOrgSsoRole(context.Background(), orgId, ssoroleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsSSORolesAPIService ListOrgSsoRoles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsSSORolesAPI.ListOrgSsoRoles(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsSSORolesAPIService UpdateOrgSsoRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var ssoroleId string

		resp, httpRes, err := apiClient.OrgsSSORolesAPI.UpdateOrgSsoRole(context.Background(), orgId, ssoroleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
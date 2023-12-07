/*
Swagger Petstore

Testing PetsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package petstore

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_petstore_PetsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PetsAPIService CreatePets", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.PetsAPI.CreatePets(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PetsAPIService ListPets", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PetsAPI.ListPets(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PetsAPIService ShowPetById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var petId string

		resp, httpRes, err := apiClient.PetsAPI.ShowPetById(context.Background(), petId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

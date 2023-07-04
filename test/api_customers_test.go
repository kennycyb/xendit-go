/*
Xendit API

Testing CustomersApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package xendit

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"

	xendit "github.com/kennycyb/xendit-go"
)

func Test_xendit_CustomersApiService(t *testing.T) {

	config := xendit.NewConfiguration()
	apiClient := xendit.NewAPIClient(config)

	t.Run("Test CustomersApiService GetCustomer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.CustomersApi.GetCustomer(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

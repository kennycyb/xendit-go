/*
Xendit API_test

Testing CustomerApiService

*/

// Code generated by OpenAPI Generator

package xendit

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	xendit "github.com/kennycyb/xendit-go"
)

func Test_xendit_CustomerApiService(t *testing.T) {

	apiKey := os.Getenv("XND_APIKEY")
	if apiKey == "" {
		t.Skip("XND_APIKEY not set")
	}
	
	apiClient := xendit.NewClient(apiKey)

	t.Run("Test CustomerApiService GetCustomer", func(t *testing.T) {

		var id string

		resp, httpRes, err := apiClient.CustomerApi.GetCustomer(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

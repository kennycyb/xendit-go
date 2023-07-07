/*
Refund Service_test

Testing RefundApiService

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

func Test_xendit_RefundApiService(t *testing.T) {

	apiKey := os.Getenv("XND_APIKEY")
	if apiKey == "" {
		t.Skip("XND_APIKEY not set")
	}
	
	apiClient := xendit.NewClient(apiKey)

	t.Run("Test RefundApiService RefundsGet", func(t *testing.T) {

		resp, httpRes, err := apiClient.RefundApi.RefundsGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RefundApiService RefundsPost", func(t *testing.T) {

		resp, httpRes, err := apiClient.RefundApi.RefundsPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RefundApiService RefundsRefundIDGet", func(t *testing.T) {

		var refundID string

		resp, httpRes, err := apiClient.RefundApi.RefundsRefundIDGet(context.Background(), refundID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

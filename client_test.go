package xenditsdk_test

import (
	"context"
	"fmt"
	"testing"

	xenditsdk "github.com/kennycyb/xendit-go"
)

func TestGetBalance(t *testing.T) {

	cfg := xenditsdk.NewConfiguration()
	cfg.Debug = true

	client := xenditsdk.NewAPIClient(cfg)

	ctx := context.WithValue(context.Background(), xenditsdk.ContextBasicAuth, xenditsdk.BasicAuth{
		UserName: "xnd_development_YNIsPvRSM30n5z5o49etyMBYemQdZ4bdp7PVBKFZDJdQj0jfGhmiyflN9xBeX",
		Password: "",
	})

	ctxTest := context.WithValue(ctx, xenditsdk.ContextServerIndex, 1)

	balance, resp, err := client.BalancesApi.GetBalance(ctxTest).AccountType("CASH").Execute()

	if balance != nil {
		fmt.Printf("Balance: %.2f\n", *balance.Balance)
	}

	fmt.Printf("Response: %v\n", resp)
	fmt.Printf("Error: %v\n", err)
}

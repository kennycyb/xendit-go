# xendit\BalancesApi

All URIs are relative to *https://api.xendit.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBalance**](BalancesApi.md#GetBalance) | **Get** /balance | Get Balance



## GetBalance

> Balance GetBalance(ctx).AccountType(accountType).ForUserId(forUserId).XIdempotencyKey(xIdempotencyKey).Execute()

Get Balance



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    xendit "github.com/kennycyb/xendit-go"
)

func main() {
    accountType := "CASH" // string | Account Type
    forUserId := "forUserId_example" // string | For User ID (optional)
    xIdempotencyKey := "xIdempotencyKey_example" // string | Idempotency Key (optional)

    xnd := xendit.NewClient("API-KEY")

    resp, r, err := xnd.BalancesApi.GetBalance(context.Background()).AccountType(accountType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BalancesApi.GetBalance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBalance`: Balance
    fmt.Fprintf(os.Stdout, "Response from `BalancesApi.GetBalance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountType** | **string** | Account Type | 
 **forUserId** | **string** | For User ID | 
 **xIdempotencyKey** | **string** | Idempotency Key | 

### Return type

[**Balance**](Balance.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


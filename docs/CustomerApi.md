# xendit\CustomerApi

All URIs are relative to *https://api.xendit.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCustomer**](CustomerApi.md#GetCustomer) | **Get** /customers/{id} | Get Customer



## GetCustomer

> Customer GetCustomer(ctx, id).ForUserId(forUserId).XIdempotencyKey(xIdempotencyKey).Execute()

Get Customer



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
    
    id := "id_example" // string | Xendit-generated Customer ID.  Will start with cust-...



    xnd := xendit.NewClient("API-KEY")

    resp, r, err := xnd.CustomerApi.GetCustomer(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerApi.GetCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomer`: Customer
    fmt.Fprintf(os.Stdout, "Response from `CustomerApi.GetCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Xendit-generated Customer ID.  Will start with cust-... | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forUserId** | **string** | For User ID | 
 **xIdempotencyKey** | **string** | Idempotency Key | 

### Return type

[**Customer**](Customer.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# xendit\RefundApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RefundsGet**](RefundApi.md#RefundsGet) | **Get** /refunds/ | 
[**RefundsPost**](RefundApi.md#RefundsPost) | **Post** /refunds | 
[**RefundsRefundIDGet**](RefundApi.md#RefundsRefundIDGet) | **Get** /refunds/{refundID} | 



## RefundsGet

> RefundList RefundsGet(ctx).XIdempotencyKey(xIdempotencyKey).Execute()



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
    

    xnd := xendit.NewClient("API-KEY")

    resp, r, err := xnd.RefundApi.RefundsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RefundApi.RefundsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RefundsGet`: RefundList
    fmt.Fprintf(os.Stdout, "Response from `RefundApi.RefundsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRefundsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xIdempotencyKey** | **string** | Idempotency Key | 

### Return type

[**RefundList**](RefundList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefundsPost

> Refund RefundsPost(ctx).IdempotencyKey(idempotencyKey).XIdempotencyKey(xIdempotencyKey).CreateRefund(createRefund).Execute()



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
    



    xnd := xendit.NewClient("API-KEY")

    resp, r, err := xnd.RefundApi.RefundsPost(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RefundApi.RefundsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RefundsPost`: Refund
    fmt.Fprintf(os.Stdout, "Response from `RefundApi.RefundsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRefundsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** |  | 
 **xIdempotencyKey** | **string** | Idempotency Key | 
 **createRefund** | [**CreateRefund**](CreateRefund.md) |  | 

### Return type

[**Refund**](Refund.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefundsRefundIDGet

> Refund RefundsRefundIDGet(ctx, refundID).XIdempotencyKey(xIdempotencyKey).Execute()



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
    
    refundID := "rfd-1fdaf346-dd2e-4b6c-b938-124c7167a822" // string | 


    xnd := xendit.NewClient("API-KEY")

    resp, r, err := xnd.RefundApi.RefundsRefundIDGet(context.Background(), refundID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RefundApi.RefundsRefundIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RefundsRefundIDGet`: Refund
    fmt.Fprintf(os.Stdout, "Response from `RefundApi.RefundsRefundIDGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**refundID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefundsRefundIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xIdempotencyKey** | **string** | Idempotency Key | 

### Return type

[**Refund**](Refund.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# xendit\PaymentrequestApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePaymentRequest**](PaymentrequestApi.md#CreatePaymentRequest) | **Post** /payment_requests | Create Payment Request
[**GetAllPaymentRequests**](PaymentrequestApi.md#GetAllPaymentRequests) | **Get** /payment_requests | Get all payment requests by filter
[**GetPaymentRequestByID**](PaymentrequestApi.md#GetPaymentRequestByID) | **Get** /payment_requests/{paymentRequestId} | Get payment request by ID
[**GetPaymentRequestCapture**](PaymentrequestApi.md#GetPaymentRequestCapture) | **Get** /payment_requests/{paymentRequestId}/captures | Get Payment Request Capture
[**PaymentRequestAuthorize**](PaymentrequestApi.md#PaymentRequestAuthorize) | **Post** /payment_requests/{paymentRequestId}/auth | Payment Request Authorize
[**PaymentRequestCapture**](PaymentrequestApi.md#PaymentRequestCapture) | **Post** /payment_requests/{paymentRequestId}/captures | Payment Request Capture
[**PaymentRequestResendAuth**](PaymentrequestApi.md#PaymentRequestResendAuth) | **Post** /payment_requests/{paymentRequestId}/auth/resend | Payment Request Resend Auth



## CreatePaymentRequest

> PublicPaymentRequest CreatePaymentRequest(ctx).XIdempotencyKey(xIdempotencyKey).PublicCreatePaymentRequest(publicCreatePaymentRequest).Execute()

Create Payment Request



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

    resp, r, err := xnd.PaymentrequestApi.CreatePaymentRequest(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentrequestApi.CreatePaymentRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePaymentRequest`: PublicPaymentRequest
    fmt.Fprintf(os.Stdout, "Response from `PaymentrequestApi.CreatePaymentRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePaymentRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xIdempotencyKey** | **string** | Idempotency Key | 
 **publicCreatePaymentRequest** | [**PublicCreatePaymentRequest**](PublicCreatePaymentRequest.md) |  | 

### Return type

[**PublicPaymentRequest**](PublicPaymentRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllPaymentRequests

> PublicPaymentRequestListResponse GetAllPaymentRequests(ctx).ReferenceId(referenceId).Limit(limit).XIdempotencyKey(xIdempotencyKey).Execute()

Get all payment requests by filter



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

    resp, r, err := xnd.PaymentrequestApi.GetAllPaymentRequests(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentrequestApi.GetAllPaymentRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllPaymentRequests`: PublicPaymentRequestListResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentrequestApi.GetAllPaymentRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllPaymentRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **referenceId** | **[]string** |  | 
 **limit** | **int32** |  | 
 **xIdempotencyKey** | **string** | Idempotency Key | 

### Return type

[**PublicPaymentRequestListResponse**](PublicPaymentRequestListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentRequestByID

> PublicPaymentRequest GetPaymentRequestByID(ctx, paymentRequestId).XIdempotencyKey(xIdempotencyKey).Execute()

Get payment request by ID



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
    
    paymentRequestId := "pr-1fdaf346-dd2e-4b6c-b938-124c7167a822" // string | 


    xnd := xendit.NewClient("API-KEY")

    resp, r, err := xnd.PaymentrequestApi.GetPaymentRequestByID(context.Background(), paymentRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentrequestApi.GetPaymentRequestByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPaymentRequestByID`: PublicPaymentRequest
    fmt.Fprintf(os.Stdout, "Response from `PaymentrequestApi.GetPaymentRequestByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentRequestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentRequestByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xIdempotencyKey** | **string** | Idempotency Key | 

### Return type

[**PublicPaymentRequest**](PublicPaymentRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentRequestCapture

> CaptureListResponse GetPaymentRequestCapture(ctx, paymentRequestId).Limit(limit).AfterId(afterId).BeforeId(beforeId).XIdempotencyKey(xIdempotencyKey).Execute()

Get Payment Request Capture



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
    
    paymentRequestId := "pr-1fdaf346-dd2e-4b6c-b938-124c7167a822" // string | 





    xnd := xendit.NewClient("API-KEY")

    resp, r, err := xnd.PaymentrequestApi.GetPaymentRequestCapture(context.Background(), paymentRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentrequestApi.GetPaymentRequestCapture``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPaymentRequestCapture`: CaptureListResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentrequestApi.GetPaymentRequestCapture`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentRequestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentRequestCaptureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **afterId** | **string** |  | 
 **beforeId** | **string** |  | 
 **xIdempotencyKey** | **string** | Idempotency Key | 

### Return type

[**CaptureListResponse**](CaptureListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentRequestAuthorize

> PublicPaymentRequest PaymentRequestAuthorize(ctx, paymentRequestId).XIdempotencyKey(xIdempotencyKey).PublicPaymentRequestAuthorize(publicPaymentRequestAuthorize).Execute()

Payment Request Authorize



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
    
    paymentRequestId := "pr-1fdaf346-dd2e-4b6c-b938-124c7167a822" // string | 



    xnd := xendit.NewClient("API-KEY")

    resp, r, err := xnd.PaymentrequestApi.PaymentRequestAuthorize(context.Background(), paymentRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentrequestApi.PaymentRequestAuthorize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaymentRequestAuthorize`: PublicPaymentRequest
    fmt.Fprintf(os.Stdout, "Response from `PaymentrequestApi.PaymentRequestAuthorize`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentRequestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentRequestAuthorizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xIdempotencyKey** | **string** | Idempotency Key | 
 **publicPaymentRequestAuthorize** | [**PublicPaymentRequestAuthorize**](PublicPaymentRequestAuthorize.md) |  | 

### Return type

[**PublicPaymentRequest**](PublicPaymentRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentRequestCapture

> Capture PaymentRequestCapture(ctx, paymentRequestId).XIdempotencyKey(xIdempotencyKey).CreateCapture(createCapture).Execute()

Payment Request Capture



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
    
    paymentRequestId := "pr-1fdaf346-dd2e-4b6c-b938-124c7167a822" // string | 



    xnd := xendit.NewClient("API-KEY")

    resp, r, err := xnd.PaymentrequestApi.PaymentRequestCapture(context.Background(), paymentRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentrequestApi.PaymentRequestCapture``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaymentRequestCapture`: Capture
    fmt.Fprintf(os.Stdout, "Response from `PaymentrequestApi.PaymentRequestCapture`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentRequestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentRequestCaptureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xIdempotencyKey** | **string** | Idempotency Key | 
 **createCapture** | [**CreateCapture**](CreateCapture.md) |  | 

### Return type

[**Capture**](Capture.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentRequestResendAuth

> PublicPaymentRequest PaymentRequestResendAuth(ctx, paymentRequestId).XIdempotencyKey(xIdempotencyKey).Execute()

Payment Request Resend Auth



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
    
    paymentRequestId := "pr-1fdaf346-dd2e-4b6c-b938-124c7167a822" // string | 


    xnd := xendit.NewClient("API-KEY")

    resp, r, err := xnd.PaymentrequestApi.PaymentRequestResendAuth(context.Background(), paymentRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentrequestApi.PaymentRequestResendAuth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaymentRequestResendAuth`: PublicPaymentRequest
    fmt.Fprintf(os.Stdout, "Response from `PaymentrequestApi.PaymentRequestResendAuth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentRequestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentRequestResendAuthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xIdempotencyKey** | **string** | Idempotency Key | 

### Return type

[**PublicPaymentRequest**](PublicPaymentRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


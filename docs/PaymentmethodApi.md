# xendit\PaymentmethodApi

All URIs are relative to *https://api.xendit.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthPaymentMethod**](PaymentmethodApi.md#AuthPaymentMethod) | **Post** /v2/payment_methods/{paymentMethodId}/auth | Validate a payment method&#39;s linking OTP
[**CreatePaymentMethod**](PaymentmethodApi.md#CreatePaymentMethod) | **Post** /v2/payment_methods | Creates payment method
[**ExpirePaymentMethod**](PaymentmethodApi.md#ExpirePaymentMethod) | **Post** /v2/payment_methods/{paymentMethodId}/expire | Expires a payment method
[**GetAllPaymentChannels**](PaymentmethodApi.md#GetAllPaymentChannels) | **Get** /v2/payment_methods/channels | Get all payment channels
[**GetAllPaymentMethods**](PaymentmethodApi.md#GetAllPaymentMethods) | **Get** /v2/payment_methods | Get all payment methods by filters
[**GetPaymentMethodByID**](PaymentmethodApi.md#GetPaymentMethodByID) | **Get** /v2/payment_methods/{paymentMethodId} | Get payment method by ID
[**PatchPaymentMethods**](PaymentmethodApi.md#PatchPaymentMethods) | **Patch** /v2/payment_methods/{paymentMethodId} | Patch payment methods
[**PublicGetPaymentsByPaymentMethodId**](PaymentmethodApi.md#PublicGetPaymentsByPaymentMethodId) | **Get** /v2/payment_methods/{paymentMethodId}/payments | Returns payments with matching PaymentMethodID.
[**PublicSimulatePaymentByPaymentMethodId**](PaymentmethodApi.md#PublicSimulatePaymentByPaymentMethodId) | **Post** /v2/payment_methods/{paymentMethodId}/payments/simulate | Makes payment with matching PaymentMethodID.
[**PublicSimulatePaymentChannelHealth**](PaymentmethodApi.md#PublicSimulatePaymentChannelHealth) | **Post** /v2/payment_methods/channels/simulate_health | Simulate payment channel health to a given business id.



## AuthPaymentMethod

> PublicPaymentMethod AuthPaymentMethod(ctx, paymentMethodId).XIdempotencyKey(xIdempotencyKey).PublicAuthPaymentMethod(publicAuthPaymentMethod).Execute()

Validate a payment method's linking OTP



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
    
    paymentMethodId := "pm-1fdaf346-dd2e-4b6c-b938-124c7167a822" // string | 



    xnd := xendit.NewClient("API-KEY")

    resp, r, err := xnd.PaymentmethodApi.AuthPaymentMethod(context.Background(), paymentMethodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentmethodApi.AuthPaymentMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthPaymentMethod`: PublicPaymentMethod
    fmt.Fprintf(os.Stdout, "Response from `PaymentmethodApi.AuthPaymentMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentMethodId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthPaymentMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xIdempotencyKey** | **string** | Idempotency Key | 
 **publicAuthPaymentMethod** | [**PublicAuthPaymentMethod**](PublicAuthPaymentMethod.md) |  | 

### Return type

[**PublicPaymentMethod**](PublicPaymentMethod.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePaymentMethod

> PublicPaymentMethod CreatePaymentMethod(ctx).XIdempotencyKey(xIdempotencyKey).PublicCreatePaymentMethod(publicCreatePaymentMethod).Execute()

Creates payment method



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

    resp, r, err := xnd.PaymentmethodApi.CreatePaymentMethod(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentmethodApi.CreatePaymentMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePaymentMethod`: PublicPaymentMethod
    fmt.Fprintf(os.Stdout, "Response from `PaymentmethodApi.CreatePaymentMethod`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePaymentMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xIdempotencyKey** | **string** | Idempotency Key | 
 **publicCreatePaymentMethod** | [**PublicCreatePaymentMethod**](PublicCreatePaymentMethod.md) |  | 

### Return type

[**PublicPaymentMethod**](PublicPaymentMethod.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExpirePaymentMethod

> PublicPaymentMethod ExpirePaymentMethod(ctx, paymentMethodId).XIdempotencyKey(xIdempotencyKey).PublicExpirePaymentMethod(publicExpirePaymentMethod).Execute()

Expires a payment method



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
    
    paymentMethodId := "pm-1fdaf346-dd2e-4b6c-b938-124c7167a822" // string | 



    xnd := xendit.NewClient("API-KEY")

    resp, r, err := xnd.PaymentmethodApi.ExpirePaymentMethod(context.Background(), paymentMethodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentmethodApi.ExpirePaymentMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExpirePaymentMethod`: PublicPaymentMethod
    fmt.Fprintf(os.Stdout, "Response from `PaymentmethodApi.ExpirePaymentMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentMethodId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExpirePaymentMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xIdempotencyKey** | **string** | Idempotency Key | 
 **publicExpirePaymentMethod** | [**PublicExpirePaymentMethod**](PublicExpirePaymentMethod.md) |  | 

### Return type

[**PublicPaymentMethod**](PublicPaymentMethod.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllPaymentChannels

> PublicPaymentChannelList GetAllPaymentChannels(ctx).IsActivated(isActivated).Type_(type_).XIdempotencyKey(xIdempotencyKey).Execute()

Get all payment channels



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

    resp, r, err := xnd.PaymentmethodApi.GetAllPaymentChannels(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentmethodApi.GetAllPaymentChannels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllPaymentChannels`: PublicPaymentChannelList
    fmt.Fprintf(os.Stdout, "Response from `PaymentmethodApi.GetAllPaymentChannels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllPaymentChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isActivated** | **bool** |  | [default to true]
 **type_** | **string** |  | 
 **xIdempotencyKey** | **string** | Idempotency Key | 

### Return type

[**PublicPaymentChannelList**](PublicPaymentChannelList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllPaymentMethods

> PublicPaymentMethodList GetAllPaymentMethods(ctx).Id(id).Type_(type_).Status(status).Reusability(reusability).CustomerId(customerId).ReferenceId(referenceId).AfterId(afterId).BeforeId(beforeId).Limit(limit).XIdempotencyKey(xIdempotencyKey).Execute()

Get all payment methods by filters



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

    resp, r, err := xnd.PaymentmethodApi.GetAllPaymentMethods(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentmethodApi.GetAllPaymentMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllPaymentMethods`: PublicPaymentMethodList
    fmt.Fprintf(os.Stdout, "Response from `PaymentmethodApi.GetAllPaymentMethods`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllPaymentMethodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **[]string** |  | 
 **type_** | **[]string** |  | 
 **status** | [**[]PaymentMethodStatus**](PaymentMethodStatus.md) |  | 
 **reusability** | [**PaymentMethodReusability**](PaymentMethodReusability.md) |  | 
 **customerId** | **string** |  | 
 **referenceId** | **string** |  | 
 **afterId** | **string** |  | 
 **beforeId** | **string** |  | 
 **limit** | **int32** |  | 
 **xIdempotencyKey** | **string** | Idempotency Key | 

### Return type

[**PublicPaymentMethodList**](PublicPaymentMethodList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentMethodByID

> PublicPaymentMethod GetPaymentMethodByID(ctx, paymentMethodId).XIdempotencyKey(xIdempotencyKey).Execute()

Get payment method by ID



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
    
    paymentMethodId := "pm-1fdaf346-dd2e-4b6c-b938-124c7167a822" // string | 


    xnd := xendit.NewClient("API-KEY")

    resp, r, err := xnd.PaymentmethodApi.GetPaymentMethodByID(context.Background(), paymentMethodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentmethodApi.GetPaymentMethodByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPaymentMethodByID`: PublicPaymentMethod
    fmt.Fprintf(os.Stdout, "Response from `PaymentmethodApi.GetPaymentMethodByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentMethodId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentMethodByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xIdempotencyKey** | **string** | Idempotency Key | 

### Return type

[**PublicPaymentMethod**](PublicPaymentMethod.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchPaymentMethods

> PublicPaymentMethod PatchPaymentMethods(ctx, paymentMethodId).XIdempotencyKey(xIdempotencyKey).PublicPatchPaymentMethod(publicPatchPaymentMethod).Execute()

Patch payment methods



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
    
    paymentMethodId := "pm-1fdaf346-dd2e-4b6c-b938-124c7167a822" // string | 



    xnd := xendit.NewClient("API-KEY")

    resp, r, err := xnd.PaymentmethodApi.PatchPaymentMethods(context.Background(), paymentMethodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentmethodApi.PatchPaymentMethods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchPaymentMethods`: PublicPaymentMethod
    fmt.Fprintf(os.Stdout, "Response from `PaymentmethodApi.PatchPaymentMethods`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentMethodId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPaymentMethodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xIdempotencyKey** | **string** | Idempotency Key | 
 **publicPatchPaymentMethod** | [**PublicPatchPaymentMethod**](PublicPatchPaymentMethod.md) |  | 

### Return type

[**PublicPaymentMethod**](PublicPaymentMethod.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublicGetPaymentsByPaymentMethodId

> map[string]interface{} PublicGetPaymentsByPaymentMethodId(ctx, paymentMethodId).PaymentRequestId(paymentRequestId).PaymentMethodId2(paymentMethodId2).ReferenceId(referenceId).PaymentMethodType(paymentMethodType).ChannelCode(channelCode).Status(status).Currency(currency).CreatedGte(createdGte).CreatedLte(createdLte).UpdatedGte(updatedGte).UpdatedLte(updatedLte).Limit(limit).AfterId(afterId).BeforeId(beforeId).XIdempotencyKey(xIdempotencyKey).Execute()

Returns payments with matching PaymentMethodID.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    xendit "github.com/kennycyb/xendit-go"
)

func main() {
    
    paymentMethodId := "pm-1fdaf346-dd2e-4b6c-b938-124c7167a822" // string | 
















    xnd := xendit.NewClient("API-KEY")

    resp, r, err := xnd.PaymentmethodApi.PublicGetPaymentsByPaymentMethodId(context.Background(), paymentMethodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentmethodApi.PublicGetPaymentsByPaymentMethodId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PublicGetPaymentsByPaymentMethodId`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PaymentmethodApi.PublicGetPaymentsByPaymentMethodId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentMethodId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublicGetPaymentsByPaymentMethodIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paymentRequestId** | **[]string** |  | 
 **paymentMethodId2** | **[]string** |  | 
 **referenceId** | **[]string** |  | 
 **paymentMethodType** | [**[]PaymentMethodType**](PaymentMethodType.md) |  | 
 **channelCode** | **[]string** |  | 
 **status** | **[]string** |  | 
 **currency** | **[]string** |  | 
 **createdGte** | **time.Time** |  | 
 **createdLte** | **time.Time** |  | 
 **updatedGte** | **time.Time** |  | 
 **updatedLte** | **time.Time** |  | 
 **limit** | **int32** |  | 
 **afterId** | **string** |  | 
 **beforeId** | **string** |  | 
 **xIdempotencyKey** | **string** | Idempotency Key | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublicSimulatePaymentByPaymentMethodId

> PublicSimulatePayment PublicSimulatePaymentByPaymentMethodId(ctx, paymentMethodId).XIdempotencyKey(xIdempotencyKey).PublicSimulatePaymentByPaymentMethodIdRequest(publicSimulatePaymentByPaymentMethodIdRequest).Execute()

Makes payment with matching PaymentMethodID.



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
    
    paymentMethodId := "pm-1fdaf346-dd2e-4b6c-b938-124c7167a822" // string | 



    xnd := xendit.NewClient("API-KEY")

    resp, r, err := xnd.PaymentmethodApi.PublicSimulatePaymentByPaymentMethodId(context.Background(), paymentMethodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentmethodApi.PublicSimulatePaymentByPaymentMethodId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PublicSimulatePaymentByPaymentMethodId`: PublicSimulatePayment
    fmt.Fprintf(os.Stdout, "Response from `PaymentmethodApi.PublicSimulatePaymentByPaymentMethodId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentMethodId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublicSimulatePaymentByPaymentMethodIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xIdempotencyKey** | **string** | Idempotency Key | 
 **publicSimulatePaymentByPaymentMethodIdRequest** | [**PublicSimulatePaymentByPaymentMethodIdRequest**](PublicSimulatePaymentByPaymentMethodIdRequest.md) |  | 

### Return type

[**PublicSimulatePayment**](PublicSimulatePayment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublicSimulatePaymentChannelHealth

> PublicSimulatePaymentChannelHealth PublicSimulatePaymentChannelHealth(ctx).PublicPostSimulatePaymentChannelHealth(publicPostSimulatePaymentChannelHealth).XIdempotencyKey(xIdempotencyKey).Execute()

Simulate payment channel health to a given business id.



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
    
    publicPostSimulatePaymentChannelHealth := *openapiclient.NewPublicPostSimulatePaymentChannelHealth("ChannelCode_example", "Status_example") // PublicPostSimulatePaymentChannelHealth | 


    xnd := xendit.NewClient("API-KEY")

    resp, r, err := xnd.PaymentmethodApi.PublicSimulatePaymentChannelHealth(context.Background()).PublicPostSimulatePaymentChannelHealth(publicPostSimulatePaymentChannelHealth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentmethodApi.PublicSimulatePaymentChannelHealth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PublicSimulatePaymentChannelHealth`: PublicSimulatePaymentChannelHealth
    fmt.Fprintf(os.Stdout, "Response from `PaymentmethodApi.PublicSimulatePaymentChannelHealth`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPublicSimulatePaymentChannelHealthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publicPostSimulatePaymentChannelHealth** | [**PublicPostSimulatePaymentChannelHealth**](PublicPostSimulatePaymentChannelHealth.md) |  | 
 **xIdempotencyKey** | **string** | Idempotency Key | 

### Return type

[**PublicSimulatePaymentChannelHealth**](PublicSimulatePaymentChannelHealth.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


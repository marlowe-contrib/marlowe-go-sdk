# \DefaultAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContractsContractIdGet**](DefaultAPI.md#ContractsContractIdGet) | **Get** /contracts/{contractId} | 
[**ContractsContractIdNextGet**](DefaultAPI.md#ContractsContractIdNextGet) | **Get** /contracts/{contractId}/next | 
[**ContractsContractIdPut**](DefaultAPI.md#ContractsContractIdPut) | **Put** /contracts/{contractId} | 
[**ContractsContractIdTransactionsGet**](DefaultAPI.md#ContractsContractIdTransactionsGet) | **Get** /contracts/{contractId}/transactions | 
[**ContractsContractIdTransactionsPost**](DefaultAPI.md#ContractsContractIdTransactionsPost) | **Post** /contracts/{contractId}/transactions | 
[**ContractsContractIdTransactionsTransactionIdGet**](DefaultAPI.md#ContractsContractIdTransactionsTransactionIdGet) | **Get** /contracts/{contractId}/transactions/{transactionId} | 
[**ContractsContractIdTransactionsTransactionIdPut**](DefaultAPI.md#ContractsContractIdTransactionsTransactionIdPut) | **Put** /contracts/{contractId}/transactions/{transactionId} | 
[**ContractsGet**](DefaultAPI.md#ContractsGet) | **Get** /contracts | 
[**ContractsPost**](DefaultAPI.md#ContractsPost) | **Post** /contracts | 
[**ContractsSourcesContractSourceIdAdjacencyGet**](DefaultAPI.md#ContractsSourcesContractSourceIdAdjacencyGet) | **Get** /contracts/sources/{contractSourceId}/adjacency | 
[**ContractsSourcesContractSourceIdClosureGet**](DefaultAPI.md#ContractsSourcesContractSourceIdClosureGet) | **Get** /contracts/sources/{contractSourceId}/closure | 
[**ContractsSourcesContractSourceIdGet**](DefaultAPI.md#ContractsSourcesContractSourceIdGet) | **Get** /contracts/sources/{contractSourceId} | 
[**ContractsSourcesPost**](DefaultAPI.md#ContractsSourcesPost) | **Post** /contracts/sources | 
[**HealthcheckGet**](DefaultAPI.md#HealthcheckGet) | **Get** /healthcheck | 
[**PayoutsGet**](DefaultAPI.md#PayoutsGet) | **Get** /payouts | 
[**PayoutsPayoutIdGet**](DefaultAPI.md#PayoutsPayoutIdGet) | **Get** /payouts/{payoutId} | 
[**WithdrawalsGet**](DefaultAPI.md#WithdrawalsGet) | **Get** /withdrawals | 
[**WithdrawalsPost**](DefaultAPI.md#WithdrawalsPost) | **Post** /withdrawals | 
[**WithdrawalsWithdrawalIdGet**](DefaultAPI.md#WithdrawalsWithdrawalIdGet) | **Get** /withdrawals/{withdrawalId} | 
[**WithdrawalsWithdrawalIdPut**](DefaultAPI.md#WithdrawalsWithdrawalIdPut) | **Put** /withdrawals/{withdrawalId} | 



## ContractsContractIdGet

> ContractsContractIdGet200Response ContractsContractIdGet(ctx, contractId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapi"
)

func main() {
    contractId := "98d601c9307dd43307cf68a03aad0086d4e07a789b66919ccf9f7f7676577eb7%231" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.ContractsContractIdGet(context.Background(), contractId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ContractsContractIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContractsContractIdGet`: ContractsContractIdGet200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ContractsContractIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContractsContractIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContractsContractIdGet200Response**](ContractsContractIdGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContractsContractIdNextGet

> Next ContractsContractIdNextGet(ctx, contractId).ValidityStart(validityStart).ValidityEnd(validityEnd).Party(party).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapi"
)

func main() {
    contractId := "98d601c9307dd43307cf68a03aad0086d4e07a789b66919ccf9f7f7676577eb7%231" // string | 
    validityStart := "validityStart_example" // string | 
    validityEnd := "validityEnd_example" // string | 
    party := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.ContractsContractIdNextGet(context.Background(), contractId).ValidityStart(validityStart).ValidityEnd(validityEnd).Party(party).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ContractsContractIdNextGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContractsContractIdNextGet`: Next
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ContractsContractIdNextGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContractsContractIdNextGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **validityStart** | **string** |  | 
 **validityEnd** | **string** |  | 
 **party** | **[]string** |  | 

### Return type

[**Next**](Next.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContractsContractIdPut

> ContractsContractIdPut(ctx, contractId).TextEnvelope(textEnvelope).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapi"
)

func main() {
    contractId := "98d601c9307dd43307cf68a03aad0086d4e07a789b66919ccf9f7f7676577eb7%231" // string | 
    textEnvelope := *openapiclient.NewTextEnvelope("CborHex_example", "Description_example", "Type_example") // TextEnvelope |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.ContractsContractIdPut(context.Background(), contractId).TextEnvelope(textEnvelope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ContractsContractIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContractsContractIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **textEnvelope** | [**TextEnvelope**](TextEnvelope.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json;charset=utf-8
- **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContractsContractIdTransactionsGet

> ListObjectTxHeader ContractsContractIdTransactionsGet(ctx, contractId).Range_(range_).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapi"
)

func main() {
    contractId := "98d601c9307dd43307cf68a03aad0086d4e07a789b66919ccf9f7f7676577eb7%231" // string | 
    range_ := "range__example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.ContractsContractIdTransactionsGet(context.Background(), contractId).Range_(range_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ContractsContractIdTransactionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContractsContractIdTransactionsGet`: ListObjectTxHeader
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ContractsContractIdTransactionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContractsContractIdTransactionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **range_** | **string** |  | 

### Return type

[**ListObjectTxHeader**](ListObjectTxHeader.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContractsContractIdTransactionsPost

> ContractsContractIdTransactionsPost201Response ContractsContractIdTransactionsPost(ctx, contractId).XChangeAddress(xChangeAddress).XAddress(xAddress).XCollateralUTxO(xCollateralUTxO).PostTransactionsRequest(postTransactionsRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapi"
)

func main() {
    contractId := "98d601c9307dd43307cf68a03aad0086d4e07a789b66919ccf9f7f7676577eb7%231" // string | 
    xChangeAddress := "addr1w94f8ywk4fg672xasahtk4t9k6w3aql943uxz5rt62d4dvq8evxaf" // string | 
    xAddress := "xAddress_example" // string |  (optional)
    xCollateralUTxO := "xCollateralUTxO_example" // string |  (optional)
    postTransactionsRequest := *openapiclient.NewPostTransactionsRequest([]openapiclient.Input{openapiclient.Input{InputOneOf: openapiclient.NewInputOneOf("ContinuationHash_example", openapiclient.Contract{ContractOneOf: openapiclient.NewContractOneOf(openapiclient.Party{PartyOneOf: openapiclient.NewPartyOneOf("RoleToken_example")}, openapiclient.Value{ValueOneOf: openapiclient.NewValueOneOf(*openapiclient.NewToken("CurrencySymbol_example", "TokenName_example"), openapiclient.Party{PartyOneOf: openapiclient.NewPartyOneOf("RoleToken_example")})}, openapiclient.Contract{ContractOneOf: openapiclient.NewContractOneOf(openapiclient.Party{PartyOneOf: }, openapiclient.Value{ValueOneOf: openapiclient.NewValueOneOf(*openapiclient.NewToken("CurrencySymbol_example", "TokenName_example"), openapiclient.Party{PartyOneOf: })}, openapiclient.Contract{ContractOneOf: }, openapiclient.Payee{PayeeOneOf: openapiclient.NewPayeeOneOf(openapiclient.Party{PartyOneOf: })}, )}, openapiclient.Payee{PayeeOneOf: openapiclient.NewPayeeOneOf(openapiclient.Party{PartyOneOf: })}, )})}}, map[string]interface{}{"key": interface{}(123)}, map[string]interface{}{"key": interface{}(123)}, openapiclient.MarloweVersion("v1")) // PostTransactionsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.ContractsContractIdTransactionsPost(context.Background(), contractId).XChangeAddress(xChangeAddress).XAddress(xAddress).XCollateralUTxO(xCollateralUTxO).PostTransactionsRequest(postTransactionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ContractsContractIdTransactionsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContractsContractIdTransactionsPost`: ContractsContractIdTransactionsPost201Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ContractsContractIdTransactionsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContractsContractIdTransactionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xChangeAddress** | **string** |  | 
 **xAddress** | **string** |  | 
 **xCollateralUTxO** | **string** |  | 
 **postTransactionsRequest** | [**PostTransactionsRequest**](PostTransactionsRequest.md) |  | 

### Return type

[**ContractsContractIdTransactionsPost201Response**](ContractsContractIdTransactionsPost201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json;charset=utf-8
- **Accept**: application/vendor.iog.marlowe-runtime.apply-inputs-tx-json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContractsContractIdTransactionsTransactionIdGet

> ContractsContractIdTransactionsTransactionIdGet200Response ContractsContractIdTransactionsTransactionIdGet(ctx, contractId, transactionId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapi"
)

func main() {
    contractId := "98d601c9307dd43307cf68a03aad0086d4e07a789b66919ccf9f7f7676577eb7%231" // string | 
    transactionId := "transactionId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.ContractsContractIdTransactionsTransactionIdGet(context.Background(), contractId, transactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ContractsContractIdTransactionsTransactionIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContractsContractIdTransactionsTransactionIdGet`: ContractsContractIdTransactionsTransactionIdGet200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ContractsContractIdTransactionsTransactionIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** |  | 
**transactionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContractsContractIdTransactionsTransactionIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ContractsContractIdTransactionsTransactionIdGet200Response**](ContractsContractIdTransactionsTransactionIdGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContractsContractIdTransactionsTransactionIdPut

> ContractsContractIdTransactionsTransactionIdPut(ctx, contractId, transactionId).TextEnvelope(textEnvelope).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapi"
)

func main() {
    contractId := "98d601c9307dd43307cf68a03aad0086d4e07a789b66919ccf9f7f7676577eb7%231" // string | 
    transactionId := "transactionId_example" // string | 
    textEnvelope := *openapiclient.NewTextEnvelope("CborHex_example", "Description_example", "Type_example") // TextEnvelope |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.ContractsContractIdTransactionsTransactionIdPut(context.Background(), contractId, transactionId).TextEnvelope(textEnvelope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ContractsContractIdTransactionsTransactionIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** |  | 
**transactionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContractsContractIdTransactionsTransactionIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **textEnvelope** | [**TextEnvelope**](TextEnvelope.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json;charset=utf-8
- **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContractsGet

> ListObjectContractHeader ContractsGet(ctx).RoleCurrency(roleCurrency).Tag(tag).PartyAddress(partyAddress).PartyRole(partyRole).Range_(range_).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapi"
)

func main() {
    roleCurrency := []string{"Inner_example"} // []string |  (optional)
    tag := []string{"Inner_example"} // []string |  (optional)
    partyAddress := []string{"addr1w94f8ywk4fg672xasahtk4t9k6w3aql943uxz5rt62d4dvq8evxaf"} // []string |  (optional)
    partyRole := []string{"Inner_example"} // []string |  (optional)
    range_ := "range__example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.ContractsGet(context.Background()).RoleCurrency(roleCurrency).Tag(tag).PartyAddress(partyAddress).PartyRole(partyRole).Range_(range_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ContractsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContractsGet`: ListObjectContractHeader
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ContractsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContractsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roleCurrency** | **[]string** |  | 
 **tag** | **[]string** |  | 
 **partyAddress** | **[]string** |  | 
 **partyRole** | **[]string** |  | 
 **range_** | **string** |  | 

### Return type

[**ListObjectContractHeader**](ListObjectContractHeader.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContractsPost

> ContractsPost201Response ContractsPost(ctx).XChangeAddress(xChangeAddress).XStakeAddress(xStakeAddress).XAddress(xAddress).XCollateralUTxO(xCollateralUTxO).PostContractsRequest(postContractsRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapi"
)

func main() {
    xChangeAddress := "addr1w94f8ywk4fg672xasahtk4t9k6w3aql943uxz5rt62d4dvq8evxaf" // string | 
    xStakeAddress := "stake1ux7lyy9nhecm033qsmel9awnr22up6jadlzkrxufr78w82gsfsn0d" // string |  (optional)
    xAddress := "xAddress_example" // string |  (optional)
    xCollateralUTxO := "xCollateralUTxO_example" // string |  (optional)
    postContractsRequest := *openapiclient.NewPostContractsRequest(openapiclient.PostContractsRequest_contract{Contract: openapiclient.Contract{ContractOneOf: openapiclient.NewContractOneOf(openapiclient.Party{PartyOneOf: openapiclient.NewPartyOneOf("RoleToken_example")}, openapiclient.Value{ValueOneOf: openapiclient.NewValueOneOf(*openapiclient.NewToken("CurrencySymbol_example", "TokenName_example"), openapiclient.Party{PartyOneOf: openapiclient.NewPartyOneOf("RoleToken_example")})}, openapiclient.Contract{ContractOneOf: openapiclient.NewContractOneOf(openapiclient.Party{PartyOneOf: }, openapiclient.Value{ValueOneOf: openapiclient.NewValueOneOf(*openapiclient.NewToken("CurrencySymbol_example", "TokenName_example"), openapiclient.Party{PartyOneOf: })}, openapiclient.Contract{ContractOneOf: }, openapiclient.Payee{PayeeOneOf: openapiclient.NewPayeeOneOf(openapiclient.Party{PartyOneOf: })}, )}, openapiclient.Payee{PayeeOneOf: openapiclient.NewPayeeOneOf(openapiclient.Party{PartyOneOf: })}, )}}, map[string]interface{}{"key": interface{}(123)}, int64(123), map[string]interface{}{"key": interface{}(123)}, openapiclient.MarloweVersion("v1")) // PostContractsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.ContractsPost(context.Background()).XChangeAddress(xChangeAddress).XStakeAddress(xStakeAddress).XAddress(xAddress).XCollateralUTxO(xCollateralUTxO).PostContractsRequest(postContractsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ContractsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContractsPost`: ContractsPost201Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ContractsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContractsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xChangeAddress** | **string** |  | 
 **xStakeAddress** | **string** |  | 
 **xAddress** | **string** |  | 
 **xCollateralUTxO** | **string** |  | 
 **postContractsRequest** | [**PostContractsRequest**](PostContractsRequest.md) |  | 

### Return type

[**ContractsPost201Response**](ContractsPost201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json;charset=utf-8
- **Accept**: application/vendor.iog.marlowe-runtime.contract-tx-json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContractsSourcesContractSourceIdAdjacencyGet

> ListObjectContractSourceId ContractsSourcesContractSourceIdAdjacencyGet(ctx, contractSourceId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapi"
)

func main() {
    contractSourceId := "contractSourceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.ContractsSourcesContractSourceIdAdjacencyGet(context.Background(), contractSourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ContractsSourcesContractSourceIdAdjacencyGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContractsSourcesContractSourceIdAdjacencyGet`: ListObjectContractSourceId
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ContractsSourcesContractSourceIdAdjacencyGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractSourceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContractsSourcesContractSourceIdAdjacencyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListObjectContractSourceId**](ListObjectContractSourceId.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContractsSourcesContractSourceIdClosureGet

> ListObjectContractSourceId ContractsSourcesContractSourceIdClosureGet(ctx, contractSourceId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapi"
)

func main() {
    contractSourceId := "contractSourceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.ContractsSourcesContractSourceIdClosureGet(context.Background(), contractSourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ContractsSourcesContractSourceIdClosureGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContractsSourcesContractSourceIdClosureGet`: ListObjectContractSourceId
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ContractsSourcesContractSourceIdClosureGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractSourceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContractsSourcesContractSourceIdClosureGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListObjectContractSourceId**](ListObjectContractSourceId.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContractsSourcesContractSourceIdGet

> Contract ContractsSourcesContractSourceIdGet(ctx, contractSourceId).Expand(expand).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapi"
)

func main() {
    contractSourceId := "contractSourceId_example" // string | 
    expand := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.ContractsSourcesContractSourceIdGet(context.Background(), contractSourceId).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ContractsSourcesContractSourceIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContractsSourcesContractSourceIdGet`: Contract
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ContractsSourcesContractSourceIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractSourceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContractsSourcesContractSourceIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **bool** |  | [default to false]

### Return type

[**Contract**](Contract.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContractsSourcesPost

> PostContractSourceResponse ContractsSourcesPost(ctx).Main(main).LabelledObject(labelledObject).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapi"
)

func main() {
    main := "main_example" // string | 
    labelledObject := []openapiclient.LabelledObject{*openapiclient.NewLabelledObject("Label_example", "Type_example", openapiclient.LabelledObject_value{ActionObject: openapiclient.ActionObject{ActionOneOf: openapiclient.NewActionOneOf(openapiclient.Value{ValueOneOf: openapiclient.NewValueOneOf(*openapiclient.NewToken("CurrencySymbol_example", "TokenName_example"), openapiclient.Party{PartyOneOf: openapiclient.NewPartyOneOf("RoleToken_example")})}, openapiclient.Party{PartyOneOf: openapiclient.NewPartyOneOf("RoleToken_example")}, *openapiclient.NewToken("CurrencySymbol_example", "TokenName_example"), openapiclient.Party{PartyOneOf: })}})} // []LabelledObject |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.ContractsSourcesPost(context.Background()).Main(main).LabelledObject(labelledObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ContractsSourcesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContractsSourcesPost`: PostContractSourceResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ContractsSourcesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContractsSourcesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **main** | **string** |  | 
 **labelledObject** | [**[]LabelledObject**](LabelledObject.md) |  | 

### Return type

[**PostContractSourceResponse**](PostContractSourceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/json;charset=utf-8
- **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HealthcheckGet

> HealthcheckGet(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.HealthcheckGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.HealthcheckGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHealthcheckGetRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PayoutsGet

> ListObjectPayoutHeader PayoutsGet(ctx).ContractId(contractId).RoleToken(roleToken).Status(status).Range_(range_).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapi"
)

func main() {
    contractId := []string{"98d601c9307dd43307cf68a03aad0086d4e07a789b66919ccf9f7f7676577eb7%231"} // []string |  (optional)
    roleToken := []string{"Inner_example"} // []string |  (optional)
    status := "status_example" // string |  (optional)
    range_ := "range__example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.PayoutsGet(context.Background()).ContractId(contractId).RoleToken(roleToken).Status(status).Range_(range_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PayoutsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PayoutsGet`: ListObjectPayoutHeader
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PayoutsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPayoutsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contractId** | **[]string** |  | 
 **roleToken** | **[]string** |  | 
 **status** | **string** |  | 
 **range_** | **string** |  | 

### Return type

[**ListObjectPayoutHeader**](ListObjectPayoutHeader.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PayoutsPayoutIdGet

> PayoutsPayoutIdGet200Response PayoutsPayoutIdGet(ctx, payoutId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapi"
)

func main() {
    payoutId := "98d601c9307dd43307cf68a03aad0086d4e07a789b66919ccf9f7f7676577eb7%231" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.PayoutsPayoutIdGet(context.Background(), payoutId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PayoutsPayoutIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PayoutsPayoutIdGet`: PayoutsPayoutIdGet200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PayoutsPayoutIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payoutId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPayoutsPayoutIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PayoutsPayoutIdGet200Response**](PayoutsPayoutIdGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WithdrawalsGet

> ListObjectWithdrawalHeader WithdrawalsGet(ctx).RoleCurrency(roleCurrency).Range_(range_).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapi"
)

func main() {
    roleCurrency := []string{"Inner_example"} // []string |  (optional)
    range_ := "range__example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.WithdrawalsGet(context.Background()).RoleCurrency(roleCurrency).Range_(range_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.WithdrawalsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WithdrawalsGet`: ListObjectWithdrawalHeader
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.WithdrawalsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWithdrawalsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roleCurrency** | **[]string** |  | 
 **range_** | **string** |  | 

### Return type

[**ListObjectWithdrawalHeader**](ListObjectWithdrawalHeader.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WithdrawalsPost

> WithdrawalsPost201Response WithdrawalsPost(ctx).XChangeAddress(xChangeAddress).XAddress(xAddress).XCollateralUTxO(xCollateralUTxO).PostWithdrawalsRequest(postWithdrawalsRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapi"
)

func main() {
    xChangeAddress := "addr1w94f8ywk4fg672xasahtk4t9k6w3aql943uxz5rt62d4dvq8evxaf" // string | 
    xAddress := "xAddress_example" // string |  (optional)
    xCollateralUTxO := "xCollateralUTxO_example" // string |  (optional)
    postWithdrawalsRequest := *openapiclient.NewPostWithdrawalsRequest([]string{"98d601c9307dd43307cf68a03aad0086d4e07a789b66919ccf9f7f7676577eb7#1"}) // PostWithdrawalsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.WithdrawalsPost(context.Background()).XChangeAddress(xChangeAddress).XAddress(xAddress).XCollateralUTxO(xCollateralUTxO).PostWithdrawalsRequest(postWithdrawalsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.WithdrawalsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WithdrawalsPost`: WithdrawalsPost201Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.WithdrawalsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWithdrawalsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xChangeAddress** | **string** |  | 
 **xAddress** | **string** |  | 
 **xCollateralUTxO** | **string** |  | 
 **postWithdrawalsRequest** | [**PostWithdrawalsRequest**](PostWithdrawalsRequest.md) |  | 

### Return type

[**WithdrawalsPost201Response**](WithdrawalsPost201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json;charset=utf-8
- **Accept**: application/vendor.iog.marlowe-runtime.withdraw-tx-json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WithdrawalsWithdrawalIdGet

> Withdrawal WithdrawalsWithdrawalIdGet(ctx, withdrawalId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapi"
)

func main() {
    withdrawalId := "withdrawalId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.WithdrawalsWithdrawalIdGet(context.Background(), withdrawalId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.WithdrawalsWithdrawalIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WithdrawalsWithdrawalIdGet`: Withdrawal
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.WithdrawalsWithdrawalIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**withdrawalId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWithdrawalsWithdrawalIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Withdrawal**](Withdrawal.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WithdrawalsWithdrawalIdPut

> WithdrawalsWithdrawalIdPut(ctx, withdrawalId).TextEnvelope(textEnvelope).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapi"
)

func main() {
    withdrawalId := "withdrawalId_example" // string | 
    textEnvelope := *openapiclient.NewTextEnvelope("CborHex_example", "Description_example", "Type_example") // TextEnvelope |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.WithdrawalsWithdrawalIdPut(context.Background(), withdrawalId).TextEnvelope(textEnvelope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.WithdrawalsWithdrawalIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**withdrawalId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWithdrawalsWithdrawalIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **textEnvelope** | [**TextEnvelope**](TextEnvelope.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json;charset=utf-8
- **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


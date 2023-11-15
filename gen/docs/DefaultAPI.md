# \DefaultAPI

All URIs are relative to *https://marlowe-runtime-preprod-web.scdev.aws.iohkdev.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyInputsToContract**](DefaultAPI.md#ApplyInputsToContract) | **Post** /contracts/{contractId}/transactions | Apply inputs to contract
[**CreateContract**](DefaultAPI.md#CreateContract) | **Post** /contracts | Create a new contract
[**CreateContractSources**](DefaultAPI.md#CreateContractSources) | **Post** /contracts/sources | Upload contract sources
[**GetContractById**](DefaultAPI.md#GetContractById) | **Get** /contracts/{contractId} | Get contract by ID
[**GetContractSourceAdjacency**](DefaultAPI.md#GetContractSourceAdjacency) | **Get** /contracts/sources/{contractSourceId}/adjacency | Get adjacent contract source IDs by ID
[**GetContractSourceById**](DefaultAPI.md#GetContractSourceById) | **Get** /contracts/sources/{contractSourceId} | Get contract source by ID
[**GetContractSourceClosure**](DefaultAPI.md#GetContractSourceClosure) | **Get** /contracts/sources/{contractSourceId}/closure | Get contract source closure by ID
[**GetContractTransactionById**](DefaultAPI.md#GetContractTransactionById) | **Get** /contracts/{contractId}/transactions/{transactionId} | Get contract transaction by ID
[**GetContracts**](DefaultAPI.md#GetContracts) | **Get** /contracts | Get contracts
[**GetNextStepsForContract**](DefaultAPI.md#GetNextStepsForContract) | **Get** /contracts/{contractId}/next | Get next contract steps
[**GetPayoutById**](DefaultAPI.md#GetPayoutById) | **Get** /payouts/{payoutId} | Get payout by ID
[**GetPayouts**](DefaultAPI.md#GetPayouts) | **Get** /payouts | Get role payouts
[**GetTransactionsForContract**](DefaultAPI.md#GetTransactionsForContract) | **Get** /contracts/{contractId}/transactions | Get transactions for contract
[**GetWithdrawalById**](DefaultAPI.md#GetWithdrawalById) | **Get** /withdrawals/{withdrawalId} | Get withdrawal by ID
[**GetWithdrawals**](DefaultAPI.md#GetWithdrawals) | **Get** /withdrawals | Get withdrawals
[**Healthcheck**](DefaultAPI.md#Healthcheck) | **Get** /healthcheck | Test server status
[**SubmitContract**](DefaultAPI.md#SubmitContract) | **Put** /contracts/{contractId} | Submit contract to chain
[**SubmitContractTransaction**](DefaultAPI.md#SubmitContractTransaction) | **Put** /contracts/{contractId}/transactions/{transactionId} | Submit contract input application
[**SubmitWithdrawal**](DefaultAPI.md#SubmitWithdrawal) | **Put** /withdrawals/{withdrawalId} | Submit payout withdrawal
[**WithdrawPayouts**](DefaultAPI.md#WithdrawPayouts) | **Post** /withdrawals | Withdraw payouts



## ApplyInputsToContract

> ApplyInputsResponse ApplyInputsToContract(ctx, contractId).XChangeAddress(xChangeAddress).XAddress(xAddress).XCollateralUTxO(xCollateralUTxO).PostTransactionsRequest(postTransactionsRequest).Execute()

Apply inputs to contract



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
    resp, r, err := apiClient.DefaultAPI.ApplyInputsToContract(context.Background(), contractId).XChangeAddress(xChangeAddress).XAddress(xAddress).XCollateralUTxO(xCollateralUTxO).PostTransactionsRequest(postTransactionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApplyInputsToContract``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplyInputsToContract`: ApplyInputsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApplyInputsToContract`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplyInputsToContractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xChangeAddress** | **string** |  | 
 **xAddress** | **string** |  | 
 **xCollateralUTxO** | **string** |  | 
 **postTransactionsRequest** | [**PostTransactionsRequest**](PostTransactionsRequest.md) |  | 

### Return type

[**ApplyInputsResponse**](ApplyInputsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json;charset=utf-8
- **Accept**: application/vendor.iog.marlowe-runtime.apply-inputs-tx-json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContract

> CreateContractResponse CreateContract(ctx).XChangeAddress(xChangeAddress).XStakeAddress(xStakeAddress).XAddress(xAddress).XCollateralUTxO(xCollateralUTxO).PostContractsRequest(postContractsRequest).Execute()

Create a new contract



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
    xStakeAddress := "stake1ux7lyy9nhecm033qsmel9awnr22up6jadlzkrxufr78w82gsfsn0d" // string | Where to send staking rewards for the Marlowe script outputs of this contract. (optional)
    xAddress := "xAddress_example" // string |  (optional)
    xCollateralUTxO := "xCollateralUTxO_example" // string |  (optional)
    postContractsRequest := *openapiclient.NewPostContractsRequest(openapiclient.PostContractsRequest_contract{Contract: openapiclient.Contract{ContractOneOf: openapiclient.NewContractOneOf(openapiclient.Party{PartyOneOf: openapiclient.NewPartyOneOf("RoleToken_example")}, openapiclient.Value{ValueOneOf: openapiclient.NewValueOneOf(*openapiclient.NewToken("CurrencySymbol_example", "TokenName_example"), openapiclient.Party{PartyOneOf: openapiclient.NewPartyOneOf("RoleToken_example")})}, openapiclient.Contract{ContractOneOf: openapiclient.NewContractOneOf(openapiclient.Party{PartyOneOf: }, openapiclient.Value{ValueOneOf: openapiclient.NewValueOneOf(*openapiclient.NewToken("CurrencySymbol_example", "TokenName_example"), openapiclient.Party{PartyOneOf: })}, openapiclient.Contract{ContractOneOf: }, openapiclient.Payee{PayeeOneOf: openapiclient.NewPayeeOneOf(openapiclient.Party{PartyOneOf: })}, )}, openapiclient.Payee{PayeeOneOf: openapiclient.NewPayeeOneOf(openapiclient.Party{PartyOneOf: })}, )}}, map[string]interface{}{"key": interface{}(123)}, map[string]interface{}{"key": interface{}(123)}, openapiclient.MarloweVersion("v1")) // PostContractsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.CreateContract(context.Background()).XChangeAddress(xChangeAddress).XStakeAddress(xStakeAddress).XAddress(xAddress).XCollateralUTxO(xCollateralUTxO).PostContractsRequest(postContractsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateContract``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateContract`: CreateContractResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateContract`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateContractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xChangeAddress** | **string** |  | 
 **xStakeAddress** | **string** | Where to send staking rewards for the Marlowe script outputs of this contract. | 
 **xAddress** | **string** |  | 
 **xCollateralUTxO** | **string** |  | 
 **postContractsRequest** | [**PostContractsRequest**](PostContractsRequest.md) |  | 

### Return type

[**CreateContractResponse**](CreateContractResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json;charset=utf-8
- **Accept**: application/vendor.iog.marlowe-runtime.contract-tx-json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContractSources

> PostContractSourceResponse CreateContractSources(ctx).Main(main).LabelledObject(labelledObject).Execute()

Upload contract sources



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
    main := "main_example" // string | The label of the top-level contract object in the bundle(s).
    labelledObject := []openapiclient.LabelledObject{*openapiclient.NewLabelledObject("Label_example", "Type_example", openapiclient.LabelledObject_value{ActionObject: openapiclient.ActionObject{ActionObjectOneOf: openapiclient.NewActionObjectOneOf("Ref_example")}})} // []LabelledObject |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.CreateContractSources(context.Background()).Main(main).LabelledObject(labelledObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateContractSources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateContractSources`: PostContractSourceResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateContractSources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateContractSourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **main** | **string** | The label of the top-level contract object in the bundle(s). | 
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


## GetContractById

> GetContractResponse GetContractById(ctx, contractId).Execute()

Get contract by ID

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
    resp, r, err := apiClient.DefaultAPI.GetContractById(context.Background(), contractId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetContractById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContractById`: GetContractResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetContractById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContractByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetContractResponse**](GetContractResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContractSourceAdjacency

> ContractSourceIds GetContractSourceAdjacency(ctx, contractSourceId).Execute()

Get adjacent contract source IDs by ID



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
    resp, r, err := apiClient.DefaultAPI.GetContractSourceAdjacency(context.Background(), contractSourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetContractSourceAdjacency``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContractSourceAdjacency`: ContractSourceIds
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetContractSourceAdjacency`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractSourceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContractSourceAdjacencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContractSourceIds**](ContractSourceIds.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContractSourceById

> Contract GetContractSourceById(ctx, contractSourceId).Expand(expand).Execute()

Get contract source by ID

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
    resp, r, err := apiClient.DefaultAPI.GetContractSourceById(context.Background(), contractSourceId).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetContractSourceById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContractSourceById`: Contract
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetContractSourceById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractSourceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContractSourceByIdRequest struct via the builder pattern


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


## GetContractSourceClosure

> ContractSourceIds GetContractSourceClosure(ctx, contractSourceId).Execute()

Get contract source closure by ID



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
    resp, r, err := apiClient.DefaultAPI.GetContractSourceClosure(context.Background(), contractSourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetContractSourceClosure``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContractSourceClosure`: ContractSourceIds
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetContractSourceClosure`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractSourceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContractSourceClosureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContractSourceIds**](ContractSourceIds.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContractTransactionById

> GetTransactionResponse GetContractTransactionById(ctx, contractId, transactionId).Execute()

Get contract transaction by ID

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
    resp, r, err := apiClient.DefaultAPI.GetContractTransactionById(context.Background(), contractId, transactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetContractTransactionById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContractTransactionById`: GetTransactionResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetContractTransactionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** |  | 
**transactionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContractTransactionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetTransactionResponse**](GetTransactionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContracts

> GetContractsResponse GetContracts(ctx).RoleCurrency(roleCurrency).Tag(tag).PartyAddress(partyAddress).PartyRole(partyRole).Range_(range_).Execute()

Get contracts



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
    resp, r, err := apiClient.DefaultAPI.GetContracts(context.Background()).RoleCurrency(roleCurrency).Tag(tag).PartyAddress(partyAddress).PartyRole(partyRole).Range_(range_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetContracts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContracts`: GetContractsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetContracts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetContractsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roleCurrency** | **[]string** |  | 
 **tag** | **[]string** |  | 
 **partyAddress** | **[]string** |  | 
 **partyRole** | **[]string** |  | 
 **range_** | **string** |  | 

### Return type

[**GetContractsResponse**](GetContractsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNextStepsForContract

> Next GetNextStepsForContract(ctx, contractId).ValidityStart(validityStart).ValidityEnd(validityEnd).Party(party).Execute()

Get next contract steps



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
    validityStart := "validityStart_example" // string | The beginning of the validity range.
    validityEnd := "validityEnd_example" // string | The end of the validity range.
    party := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetNextStepsForContract(context.Background(), contractId).ValidityStart(validityStart).ValidityEnd(validityEnd).Party(party).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetNextStepsForContract``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNextStepsForContract`: Next
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetNextStepsForContract`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNextStepsForContractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **validityStart** | **string** | The beginning of the validity range. | 
 **validityEnd** | **string** | The end of the validity range. | 
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


## GetPayoutById

> GetPayoutResponse GetPayoutById(ctx, payoutId).Execute()

Get payout by ID

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
    resp, r, err := apiClient.DefaultAPI.GetPayoutById(context.Background(), payoutId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetPayoutById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPayoutById`: GetPayoutResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetPayoutById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payoutId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPayoutByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetPayoutResponse**](GetPayoutResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPayouts

> GetPayoutsResponse GetPayouts(ctx).ContractId(contractId).RoleToken(roleToken).Status(status).Range_(range_).Execute()

Get role payouts



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
    status := "status_example" // string | Whether to include available or withdrawn payouts in the results. (optional)
    range_ := "range__example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetPayouts(context.Background()).ContractId(contractId).RoleToken(roleToken).Status(status).Range_(range_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetPayouts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPayouts`: GetPayoutsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetPayouts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPayoutsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contractId** | **[]string** |  | 
 **roleToken** | **[]string** |  | 
 **status** | **string** | Whether to include available or withdrawn payouts in the results. | 
 **range_** | **string** |  | 

### Return type

[**GetPayoutsResponse**](GetPayoutsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactionsForContract

> GetTransactionsResponse GetTransactionsForContract(ctx, contractId).Range_(range_).Execute()

Get transactions for contract



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
    resp, r, err := apiClient.DefaultAPI.GetTransactionsForContract(context.Background(), contractId).Range_(range_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetTransactionsForContract``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransactionsForContract`: GetTransactionsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetTransactionsForContract`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionsForContractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **range_** | **string** |  | 

### Return type

[**GetTransactionsResponse**](GetTransactionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWithdrawalById

> Withdrawal GetWithdrawalById(ctx, withdrawalId).Execute()

Get withdrawal by ID

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
    resp, r, err := apiClient.DefaultAPI.GetWithdrawalById(context.Background(), withdrawalId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetWithdrawalById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWithdrawalById`: Withdrawal
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetWithdrawalById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**withdrawalId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWithdrawalByIdRequest struct via the builder pattern


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


## GetWithdrawals

> GetWithdrawalsResponse GetWithdrawals(ctx).RoleCurrency(roleCurrency).Range_(range_).Execute()

Get withdrawals



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
    resp, r, err := apiClient.DefaultAPI.GetWithdrawals(context.Background()).RoleCurrency(roleCurrency).Range_(range_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetWithdrawals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWithdrawals`: GetWithdrawalsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetWithdrawals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWithdrawalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roleCurrency** | **[]string** |  | 
 **range_** | **string** |  | 

### Return type

[**GetWithdrawalsResponse**](GetWithdrawalsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Healthcheck

> Healthcheck(ctx).Execute()

Test server status



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
    r, err := apiClient.DefaultAPI.Healthcheck(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.Healthcheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHealthcheckRequest struct via the builder pattern


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


## SubmitContract

> SubmitContract(ctx, contractId).TextEnvelope(textEnvelope).Execute()

Submit contract to chain



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
    r, err := apiClient.DefaultAPI.SubmitContract(context.Background(), contractId).TextEnvelope(textEnvelope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SubmitContract``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSubmitContractRequest struct via the builder pattern


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


## SubmitContractTransaction

> SubmitContractTransaction(ctx, contractId, transactionId).TextEnvelope(textEnvelope).Execute()

Submit contract input application



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
    r, err := apiClient.DefaultAPI.SubmitContractTransaction(context.Background(), contractId, transactionId).TextEnvelope(textEnvelope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SubmitContractTransaction``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSubmitContractTransactionRequest struct via the builder pattern


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


## SubmitWithdrawal

> SubmitWithdrawal(ctx, withdrawalId).TextEnvelope(textEnvelope).Execute()

Submit payout withdrawal



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
    r, err := apiClient.DefaultAPI.SubmitWithdrawal(context.Background(), withdrawalId).TextEnvelope(textEnvelope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SubmitWithdrawal``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSubmitWithdrawalRequest struct via the builder pattern


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


## WithdrawPayouts

> WithdrawPayoutsResponse WithdrawPayouts(ctx).XChangeAddress(xChangeAddress).XAddress(xAddress).XCollateralUTxO(xCollateralUTxO).PostWithdrawalsRequest(postWithdrawalsRequest).Execute()

Withdraw payouts



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
    resp, r, err := apiClient.DefaultAPI.WithdrawPayouts(context.Background()).XChangeAddress(xChangeAddress).XAddress(xAddress).XCollateralUTxO(xCollateralUTxO).PostWithdrawalsRequest(postWithdrawalsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.WithdrawPayouts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WithdrawPayouts`: WithdrawPayoutsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.WithdrawPayouts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWithdrawPayoutsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xChangeAddress** | **string** |  | 
 **xAddress** | **string** |  | 
 **xCollateralUTxO** | **string** |  | 
 **postWithdrawalsRequest** | [**PostWithdrawalsRequest**](PostWithdrawalsRequest.md) |  | 

### Return type

[**WithdrawPayoutsResponse**](WithdrawPayoutsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json;charset=utf-8
- **Accept**: application/vendor.iog.marlowe-runtime.withdraw-tx-json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


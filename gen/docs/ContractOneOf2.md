# ContractOneOf2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timeout** | **int32** |  | 
**TimeoutContinuation** | [**Contract**](Contract.md) |  | 
**When** | [**[]Case**](Case.md) |  | 

## Methods

### NewContractOneOf2

`func NewContractOneOf2(timeout int32, timeoutContinuation Contract, when []Case, ) *ContractOneOf2`

NewContractOneOf2 instantiates a new ContractOneOf2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractOneOf2WithDefaults

`func NewContractOneOf2WithDefaults() *ContractOneOf2`

NewContractOneOf2WithDefaults instantiates a new ContractOneOf2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeout

`func (o *ContractOneOf2) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ContractOneOf2) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ContractOneOf2) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.


### GetTimeoutContinuation

`func (o *ContractOneOf2) GetTimeoutContinuation() Contract`

GetTimeoutContinuation returns the TimeoutContinuation field if non-nil, zero value otherwise.

### GetTimeoutContinuationOk

`func (o *ContractOneOf2) GetTimeoutContinuationOk() (*Contract, bool)`

GetTimeoutContinuationOk returns a tuple with the TimeoutContinuation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutContinuation

`func (o *ContractOneOf2) SetTimeoutContinuation(v Contract)`

SetTimeoutContinuation sets TimeoutContinuation field to given value.


### GetWhen

`func (o *ContractOneOf2) GetWhen() []Case`

GetWhen returns the When field if non-nil, zero value otherwise.

### GetWhenOk

`func (o *ContractOneOf2) GetWhenOk() (*[]Case, bool)`

GetWhenOk returns a tuple with the When field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhen

`func (o *ContractOneOf2) SetWhen(v []Case)`

SetWhen sets When field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ContractObjectOneOf2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timeout** | **int32** |  | 
**TimeoutContinuation** | [**ContractObject**](ContractObject.md) |  | 
**When** | [**[]CaseObject**](CaseObject.md) |  | 

## Methods

### NewContractObjectOneOf2

`func NewContractObjectOneOf2(timeout int32, timeoutContinuation ContractObject, when []CaseObject, ) *ContractObjectOneOf2`

NewContractObjectOneOf2 instantiates a new ContractObjectOneOf2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractObjectOneOf2WithDefaults

`func NewContractObjectOneOf2WithDefaults() *ContractObjectOneOf2`

NewContractObjectOneOf2WithDefaults instantiates a new ContractObjectOneOf2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeout

`func (o *ContractObjectOneOf2) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ContractObjectOneOf2) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ContractObjectOneOf2) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.


### GetTimeoutContinuation

`func (o *ContractObjectOneOf2) GetTimeoutContinuation() ContractObject`

GetTimeoutContinuation returns the TimeoutContinuation field if non-nil, zero value otherwise.

### GetTimeoutContinuationOk

`func (o *ContractObjectOneOf2) GetTimeoutContinuationOk() (*ContractObject, bool)`

GetTimeoutContinuationOk returns a tuple with the TimeoutContinuation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutContinuation

`func (o *ContractObjectOneOf2) SetTimeoutContinuation(v ContractObject)`

SetTimeoutContinuation sets TimeoutContinuation field to given value.


### GetWhen

`func (o *ContractObjectOneOf2) GetWhen() []CaseObject`

GetWhen returns the When field if non-nil, zero value otherwise.

### GetWhenOk

`func (o *ContractObjectOneOf2) GetWhenOk() (*[]CaseObject, bool)`

GetWhenOk returns a tuple with the When field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhen

`func (o *ContractObjectOneOf2) SetWhen(v []CaseObject)`

SetWhen sets When field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



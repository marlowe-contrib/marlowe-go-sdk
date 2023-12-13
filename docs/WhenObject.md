# WhenObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timeout** | **int64** |  | 
**TimeoutContinuation** | [**ContractObject**](ContractObject.md) |  | 
**When** | [**[]CaseObject**](CaseObject.md) |  | 

## Methods

### NewWhenObject

`func NewWhenObject(timeout int64, timeoutContinuation ContractObject, when []CaseObject, ) *WhenObject`

NewWhenObject instantiates a new WhenObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhenObjectWithDefaults

`func NewWhenObjectWithDefaults() *WhenObject`

NewWhenObjectWithDefaults instantiates a new WhenObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeout

`func (o *WhenObject) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *WhenObject) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *WhenObject) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.


### GetTimeoutContinuation

`func (o *WhenObject) GetTimeoutContinuation() ContractObject`

GetTimeoutContinuation returns the TimeoutContinuation field if non-nil, zero value otherwise.

### GetTimeoutContinuationOk

`func (o *WhenObject) GetTimeoutContinuationOk() (*ContractObject, bool)`

GetTimeoutContinuationOk returns a tuple with the TimeoutContinuation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutContinuation

`func (o *WhenObject) SetTimeoutContinuation(v ContractObject)`

SetTimeoutContinuation sets TimeoutContinuation field to given value.


### GetWhen

`func (o *WhenObject) GetWhen() []CaseObject`

GetWhen returns the When field if non-nil, zero value otherwise.

### GetWhenOk

`func (o *WhenObject) GetWhenOk() (*[]CaseObject, bool)`

GetWhenOk returns a tuple with the When field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhen

`func (o *WhenObject) SetWhen(v []CaseObject)`

SetWhen sets When field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# When

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timeout** | **int64** |  | 
**TimeoutContinuation** | [**Contract**](Contract.md) |  | 
**When** | [**[]Case**](Case.md) |  | 

## Methods

### NewWhen

`func NewWhen(timeout int64, timeoutContinuation Contract, when []Case, ) *When`

NewWhen instantiates a new When object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhenWithDefaults

`func NewWhenWithDefaults() *When`

NewWhenWithDefaults instantiates a new When object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeout

`func (o *When) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *When) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *When) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.


### GetTimeoutContinuation

`func (o *When) GetTimeoutContinuation() Contract`

GetTimeoutContinuation returns the TimeoutContinuation field if non-nil, zero value otherwise.

### GetTimeoutContinuationOk

`func (o *When) GetTimeoutContinuationOk() (*Contract, bool)`

GetTimeoutContinuationOk returns a tuple with the TimeoutContinuation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutContinuation

`func (o *When) SetTimeoutContinuation(v Contract)`

SetTimeoutContinuation sets TimeoutContinuation field to given value.


### GetWhen

`func (o *When) GetWhen() []Case`

GetWhen returns the When field if non-nil, zero value otherwise.

### GetWhenOk

`func (o *When) GetWhenOk() (*[]Case, bool)`

GetWhenOk returns a tuple with the When field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhen

`func (o *When) SetWhen(v []Case)`

SetWhen sets When field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



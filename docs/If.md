# If

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Else** | [**Contract**](Contract.md) |  | 
**If** | [**Observation**](Observation.md) |  | 
**Then** | [**Contract**](Contract.md) |  | 

## Methods

### NewIf

`func NewIf(else_ Contract, if_ Observation, then Contract, ) *If`

NewIf instantiates a new If object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIfWithDefaults

`func NewIfWithDefaults() *If`

NewIfWithDefaults instantiates a new If object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElse

`func (o *If) GetElse() Contract`

GetElse returns the Else field if non-nil, zero value otherwise.

### GetElseOk

`func (o *If) GetElseOk() (*Contract, bool)`

GetElseOk returns a tuple with the Else field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElse

`func (o *If) SetElse(v Contract)`

SetElse sets Else field to given value.


### GetIf

`func (o *If) GetIf() Observation`

GetIf returns the If field if non-nil, zero value otherwise.

### GetIfOk

`func (o *If) GetIfOk() (*Observation, bool)`

GetIfOk returns a tuple with the If field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIf

`func (o *If) SetIf(v Observation)`

SetIf sets If field to given value.


### GetThen

`func (o *If) GetThen() Contract`

GetThen returns the Then field if non-nil, zero value otherwise.

### GetThenOk

`func (o *If) GetThenOk() (*Contract, bool)`

GetThenOk returns a tuple with the Then field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThen

`func (o *If) SetThen(v Contract)`

SetThen sets Then field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



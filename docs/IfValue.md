# IfValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Else** | [**Value**](Value.md) |  | 
**If** | [**Observation**](Observation.md) |  | 
**Then** | [**Value**](Value.md) |  | 

## Methods

### NewIfValue

`func NewIfValue(else_ Value, if_ Observation, then Value, ) *IfValue`

NewIfValue instantiates a new IfValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIfValueWithDefaults

`func NewIfValueWithDefaults() *IfValue`

NewIfValueWithDefaults instantiates a new IfValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElse

`func (o *IfValue) GetElse() Value`

GetElse returns the Else field if non-nil, zero value otherwise.

### GetElseOk

`func (o *IfValue) GetElseOk() (*Value, bool)`

GetElseOk returns a tuple with the Else field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElse

`func (o *IfValue) SetElse(v Value)`

SetElse sets Else field to given value.


### GetIf

`func (o *IfValue) GetIf() Observation`

GetIf returns the If field if non-nil, zero value otherwise.

### GetIfOk

`func (o *IfValue) GetIfOk() (*Observation, bool)`

GetIfOk returns a tuple with the If field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIf

`func (o *IfValue) SetIf(v Observation)`

SetIf sets If field to given value.


### GetThen

`func (o *IfValue) GetThen() Value`

GetThen returns the Then field if non-nil, zero value otherwise.

### GetThenOk

`func (o *IfValue) GetThenOk() (*Value, bool)`

GetThenOk returns a tuple with the Then field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThen

`func (o *IfValue) SetThen(v Value)`

SetThen sets Then field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



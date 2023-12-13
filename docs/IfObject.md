# IfObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Else** | [**ContractObject**](ContractObject.md) |  | 
**If** | [**ObservationObject**](ObservationObject.md) |  | 
**Then** | [**ContractObject**](ContractObject.md) |  | 

## Methods

### NewIfObject

`func NewIfObject(else_ ContractObject, if_ ObservationObject, then ContractObject, ) *IfObject`

NewIfObject instantiates a new IfObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIfObjectWithDefaults

`func NewIfObjectWithDefaults() *IfObject`

NewIfObjectWithDefaults instantiates a new IfObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElse

`func (o *IfObject) GetElse() ContractObject`

GetElse returns the Else field if non-nil, zero value otherwise.

### GetElseOk

`func (o *IfObject) GetElseOk() (*ContractObject, bool)`

GetElseOk returns a tuple with the Else field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElse

`func (o *IfObject) SetElse(v ContractObject)`

SetElse sets Else field to given value.


### GetIf

`func (o *IfObject) GetIf() ObservationObject`

GetIf returns the If field if non-nil, zero value otherwise.

### GetIfOk

`func (o *IfObject) GetIfOk() (*ObservationObject, bool)`

GetIfOk returns a tuple with the If field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIf

`func (o *IfObject) SetIf(v ObservationObject)`

SetIf sets If field to given value.


### GetThen

`func (o *IfObject) GetThen() ContractObject`

GetThen returns the Then field if non-nil, zero value otherwise.

### GetThenOk

`func (o *IfObject) GetThenOk() (*ContractObject, bool)`

GetThenOk returns a tuple with the Then field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThen

`func (o *IfObject) SetThen(v ContractObject)`

SetThen sets Then field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



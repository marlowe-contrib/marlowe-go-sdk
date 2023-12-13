# Assert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assert** | [**Observation**](Observation.md) |  | 
**Then** | [**Contract**](Contract.md) |  | 

## Methods

### NewAssert

`func NewAssert(assert Observation, then Contract, ) *Assert`

NewAssert instantiates a new Assert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssertWithDefaults

`func NewAssertWithDefaults() *Assert`

NewAssertWithDefaults instantiates a new Assert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssert

`func (o *Assert) GetAssert() Observation`

GetAssert returns the Assert field if non-nil, zero value otherwise.

### GetAssertOk

`func (o *Assert) GetAssertOk() (*Observation, bool)`

GetAssertOk returns a tuple with the Assert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssert

`func (o *Assert) SetAssert(v Observation)`

SetAssert sets Assert field to given value.


### GetThen

`func (o *Assert) GetThen() Contract`

GetThen returns the Then field if non-nil, zero value otherwise.

### GetThenOk

`func (o *Assert) GetThenOk() (*Contract, bool)`

GetThenOk returns a tuple with the Then field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThen

`func (o *Assert) SetThen(v Contract)`

SetThen sets Then field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



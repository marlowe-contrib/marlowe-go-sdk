# AssertObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assert** | [**ObservationObject**](ObservationObject.md) |  | 
**Then** | [**ContractObject**](ContractObject.md) |  | 

## Methods

### NewAssertObject

`func NewAssertObject(assert ObservationObject, then ContractObject, ) *AssertObject`

NewAssertObject instantiates a new AssertObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssertObjectWithDefaults

`func NewAssertObjectWithDefaults() *AssertObject`

NewAssertObjectWithDefaults instantiates a new AssertObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssert

`func (o *AssertObject) GetAssert() ObservationObject`

GetAssert returns the Assert field if non-nil, zero value otherwise.

### GetAssertOk

`func (o *AssertObject) GetAssertOk() (*ObservationObject, bool)`

GetAssertOk returns a tuple with the Assert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssert

`func (o *AssertObject) SetAssert(v ObservationObject)`

SetAssert sets Assert field to given value.


### GetThen

`func (o *AssertObject) GetThen() ContractObject`

GetThen returns the Then field if non-nil, zero value otherwise.

### GetThenOk

`func (o *AssertObject) GetThenOk() (*ContractObject, bool)`

GetThenOk returns a tuple with the Then field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThen

`func (o *AssertObject) SetThen(v ContractObject)`

SetThen sets Then field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



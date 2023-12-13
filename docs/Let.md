# Let

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Be** | [**Value**](Value.md) |  | 
**Let** | **string** |  | 
**Then** | [**Contract**](Contract.md) |  | 

## Methods

### NewLet

`func NewLet(be Value, let string, then Contract, ) *Let`

NewLet instantiates a new Let object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLetWithDefaults

`func NewLetWithDefaults() *Let`

NewLetWithDefaults instantiates a new Let object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBe

`func (o *Let) GetBe() Value`

GetBe returns the Be field if non-nil, zero value otherwise.

### GetBeOk

`func (o *Let) GetBeOk() (*Value, bool)`

GetBeOk returns a tuple with the Be field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBe

`func (o *Let) SetBe(v Value)`

SetBe sets Be field to given value.


### GetLet

`func (o *Let) GetLet() string`

GetLet returns the Let field if non-nil, zero value otherwise.

### GetLetOk

`func (o *Let) GetLetOk() (*string, bool)`

GetLetOk returns a tuple with the Let field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLet

`func (o *Let) SetLet(v string)`

SetLet sets Let field to given value.


### GetThen

`func (o *Let) GetThen() Contract`

GetThen returns the Then field if non-nil, zero value otherwise.

### GetThenOk

`func (o *Let) GetThenOk() (*Contract, bool)`

GetThenOk returns a tuple with the Then field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThen

`func (o *Let) SetThen(v Contract)`

SetThen sets Then field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



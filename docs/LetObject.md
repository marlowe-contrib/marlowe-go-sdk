# LetObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Be** | [**ValueObject**](ValueObject.md) |  | 
**Let** | **string** |  | 
**Then** | [**ContractObject**](ContractObject.md) |  | 

## Methods

### NewLetObject

`func NewLetObject(be ValueObject, let string, then ContractObject, ) *LetObject`

NewLetObject instantiates a new LetObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLetObjectWithDefaults

`func NewLetObjectWithDefaults() *LetObject`

NewLetObjectWithDefaults instantiates a new LetObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBe

`func (o *LetObject) GetBe() ValueObject`

GetBe returns the Be field if non-nil, zero value otherwise.

### GetBeOk

`func (o *LetObject) GetBeOk() (*ValueObject, bool)`

GetBeOk returns a tuple with the Be field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBe

`func (o *LetObject) SetBe(v ValueObject)`

SetBe sets Be field to given value.


### GetLet

`func (o *LetObject) GetLet() string`

GetLet returns the Let field if non-nil, zero value otherwise.

### GetLetOk

`func (o *LetObject) GetLetOk() (*string, bool)`

GetLetOk returns a tuple with the Let field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLet

`func (o *LetObject) SetLet(v string)`

SetLet sets Let field to given value.


### GetThen

`func (o *LetObject) GetThen() ContractObject`

GetThen returns the Then field if non-nil, zero value otherwise.

### GetThenOk

`func (o *LetObject) GetThenOk() (*ContractObject, bool)`

GetThenOk returns a tuple with the Then field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThen

`func (o *LetObject) SetThen(v ContractObject)`

SetThen sets Then field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



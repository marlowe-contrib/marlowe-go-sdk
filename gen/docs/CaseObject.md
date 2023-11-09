# CaseObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Case** | [**ActionObject**](ActionObject.md) |  | 
**Then** | [**ContractObject**](ContractObject.md) |  | 
**MerkleizedThen** | **string** |  | 

## Methods

### NewCaseObject

`func NewCaseObject(case_ ActionObject, then ContractObject, merkleizedThen string, ) *CaseObject`

NewCaseObject instantiates a new CaseObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaseObjectWithDefaults

`func NewCaseObjectWithDefaults() *CaseObject`

NewCaseObjectWithDefaults instantiates a new CaseObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCase

`func (o *CaseObject) GetCase() ActionObject`

GetCase returns the Case field if non-nil, zero value otherwise.

### GetCaseOk

`func (o *CaseObject) GetCaseOk() (*ActionObject, bool)`

GetCaseOk returns a tuple with the Case field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCase

`func (o *CaseObject) SetCase(v ActionObject)`

SetCase sets Case field to given value.


### GetThen

`func (o *CaseObject) GetThen() ContractObject`

GetThen returns the Then field if non-nil, zero value otherwise.

### GetThenOk

`func (o *CaseObject) GetThenOk() (*ContractObject, bool)`

GetThenOk returns a tuple with the Then field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThen

`func (o *CaseObject) SetThen(v ContractObject)`

SetThen sets Then field to given value.


### GetMerkleizedThen

`func (o *CaseObject) GetMerkleizedThen() string`

GetMerkleizedThen returns the MerkleizedThen field if non-nil, zero value otherwise.

### GetMerkleizedThenOk

`func (o *CaseObject) GetMerkleizedThenOk() (*string, bool)`

GetMerkleizedThenOk returns a tuple with the MerkleizedThen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerkleizedThen

`func (o *CaseObject) SetMerkleizedThen(v string)`

SetMerkleizedThen sets MerkleizedThen field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



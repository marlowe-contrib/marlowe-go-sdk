# Case

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Case** | [**Action**](Action.md) |  | 
**Then** | [**Contract**](Contract.md) |  | 
**MerkleizedThen** | **string** |  | 

## Methods

### NewCase

`func NewCase(case_ Action, then Contract, merkleizedThen string, ) *Case`

NewCase instantiates a new Case object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaseWithDefaults

`func NewCaseWithDefaults() *Case`

NewCaseWithDefaults instantiates a new Case object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCase

`func (o *Case) GetCase() Action`

GetCase returns the Case field if non-nil, zero value otherwise.

### GetCaseOk

`func (o *Case) GetCaseOk() (*Action, bool)`

GetCaseOk returns a tuple with the Case field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCase

`func (o *Case) SetCase(v Action)`

SetCase sets Case field to given value.


### GetThen

`func (o *Case) GetThen() Contract`

GetThen returns the Then field if non-nil, zero value otherwise.

### GetThenOk

`func (o *Case) GetThenOk() (*Contract, bool)`

GetThenOk returns a tuple with the Then field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThen

`func (o *Case) SetThen(v Contract)`

SetThen sets Then field to given value.


### GetMerkleizedThen

`func (o *Case) GetMerkleizedThen() string`

GetMerkleizedThen returns the MerkleizedThen field if non-nil, zero value otherwise.

### GetMerkleizedThenOk

`func (o *Case) GetMerkleizedThenOk() (*string, bool)`

GetMerkleizedThenOk returns a tuple with the MerkleizedThen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerkleizedThen

`func (o *Case) SetMerkleizedThen(v string)`

SetMerkleizedThen sets MerkleizedThen field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



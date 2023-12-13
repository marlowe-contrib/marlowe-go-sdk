# ContinuationInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinuationHash** | **string** |  | 
**MerkleizedContinuation** | [**Contract**](Contract.md) |  | 

## Methods

### NewContinuationInput

`func NewContinuationInput(continuationHash string, merkleizedContinuation Contract, ) *ContinuationInput`

NewContinuationInput instantiates a new ContinuationInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContinuationInputWithDefaults

`func NewContinuationInputWithDefaults() *ContinuationInput`

NewContinuationInputWithDefaults instantiates a new ContinuationInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinuationHash

`func (o *ContinuationInput) GetContinuationHash() string`

GetContinuationHash returns the ContinuationHash field if non-nil, zero value otherwise.

### GetContinuationHashOk

`func (o *ContinuationInput) GetContinuationHashOk() (*string, bool)`

GetContinuationHashOk returns a tuple with the ContinuationHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuationHash

`func (o *ContinuationInput) SetContinuationHash(v string)`

SetContinuationHash sets ContinuationHash field to given value.


### GetMerkleizedContinuation

`func (o *ContinuationInput) GetMerkleizedContinuation() Contract`

GetMerkleizedContinuation returns the MerkleizedContinuation field if non-nil, zero value otherwise.

### GetMerkleizedContinuationOk

`func (o *ContinuationInput) GetMerkleizedContinuationOk() (*Contract, bool)`

GetMerkleizedContinuationOk returns a tuple with the MerkleizedContinuation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerkleizedContinuation

`func (o *ContinuationInput) SetMerkleizedContinuation(v Contract)`

SetMerkleizedContinuation sets MerkleizedContinuation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



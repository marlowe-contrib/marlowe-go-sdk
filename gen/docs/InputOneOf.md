# InputOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinuationHash** | **string** |  | 
**MerkleizedContinuation** | [**Contract**](Contract.md) |  | 

## Methods

### NewInputOneOf

`func NewInputOneOf(continuationHash string, merkleizedContinuation Contract, ) *InputOneOf`

NewInputOneOf instantiates a new InputOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputOneOfWithDefaults

`func NewInputOneOfWithDefaults() *InputOneOf`

NewInputOneOfWithDefaults instantiates a new InputOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinuationHash

`func (o *InputOneOf) GetContinuationHash() string`

GetContinuationHash returns the ContinuationHash field if non-nil, zero value otherwise.

### GetContinuationHashOk

`func (o *InputOneOf) GetContinuationHashOk() (*string, bool)`

GetContinuationHashOk returns a tuple with the ContinuationHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuationHash

`func (o *InputOneOf) SetContinuationHash(v string)`

SetContinuationHash sets ContinuationHash field to given value.


### GetMerkleizedContinuation

`func (o *InputOneOf) GetMerkleizedContinuation() Contract`

GetMerkleizedContinuation returns the MerkleizedContinuation field if non-nil, zero value otherwise.

### GetMerkleizedContinuationOk

`func (o *InputOneOf) GetMerkleizedContinuationOk() (*Contract, bool)`

GetMerkleizedContinuationOk returns a tuple with the MerkleizedContinuation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerkleizedContinuation

`func (o *InputOneOf) SetMerkleizedContinuation(v Contract)`

SetMerkleizedContinuation sets MerkleizedContinuation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# InputOneOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinuationHash** | **string** |  | 
**ForChoiceId** | [**ChoiceId**](ChoiceId.md) |  | 
**InputThatChoosesNum** | **int32** |  | 
**MerkleizedContinuation** | [**Contract**](Contract.md) |  | 

## Methods

### NewInputOneOf1

`func NewInputOneOf1(continuationHash string, forChoiceId ChoiceId, inputThatChoosesNum int32, merkleizedContinuation Contract, ) *InputOneOf1`

NewInputOneOf1 instantiates a new InputOneOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputOneOf1WithDefaults

`func NewInputOneOf1WithDefaults() *InputOneOf1`

NewInputOneOf1WithDefaults instantiates a new InputOneOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinuationHash

`func (o *InputOneOf1) GetContinuationHash() string`

GetContinuationHash returns the ContinuationHash field if non-nil, zero value otherwise.

### GetContinuationHashOk

`func (o *InputOneOf1) GetContinuationHashOk() (*string, bool)`

GetContinuationHashOk returns a tuple with the ContinuationHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuationHash

`func (o *InputOneOf1) SetContinuationHash(v string)`

SetContinuationHash sets ContinuationHash field to given value.


### GetForChoiceId

`func (o *InputOneOf1) GetForChoiceId() ChoiceId`

GetForChoiceId returns the ForChoiceId field if non-nil, zero value otherwise.

### GetForChoiceIdOk

`func (o *InputOneOf1) GetForChoiceIdOk() (*ChoiceId, bool)`

GetForChoiceIdOk returns a tuple with the ForChoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForChoiceId

`func (o *InputOneOf1) SetForChoiceId(v ChoiceId)`

SetForChoiceId sets ForChoiceId field to given value.


### GetInputThatChoosesNum

`func (o *InputOneOf1) GetInputThatChoosesNum() int32`

GetInputThatChoosesNum returns the InputThatChoosesNum field if non-nil, zero value otherwise.

### GetInputThatChoosesNumOk

`func (o *InputOneOf1) GetInputThatChoosesNumOk() (*int32, bool)`

GetInputThatChoosesNumOk returns a tuple with the InputThatChoosesNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputThatChoosesNum

`func (o *InputOneOf1) SetInputThatChoosesNum(v int32)`

SetInputThatChoosesNum sets InputThatChoosesNum field to given value.


### GetMerkleizedContinuation

`func (o *InputOneOf1) GetMerkleizedContinuation() Contract`

GetMerkleizedContinuation returns the MerkleizedContinuation field if non-nil, zero value otherwise.

### GetMerkleizedContinuationOk

`func (o *InputOneOf1) GetMerkleizedContinuationOk() (*Contract, bool)`

GetMerkleizedContinuationOk returns a tuple with the MerkleizedContinuation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerkleizedContinuation

`func (o *InputOneOf1) SetMerkleizedContinuation(v Contract)`

SetMerkleizedContinuation sets MerkleizedContinuation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



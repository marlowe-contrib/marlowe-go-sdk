# ChoiceContinuationInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinuationHash** | **string** |  | 
**ForChoiceId** | [**ChoiceId**](ChoiceId.md) |  | 
**InputThatChoosesNum** | **int32** |  | 
**MerkleizedContinuation** | [**Contract**](Contract.md) |  | 

## Methods

### NewChoiceContinuationInput

`func NewChoiceContinuationInput(continuationHash string, forChoiceId ChoiceId, inputThatChoosesNum int32, merkleizedContinuation Contract, ) *ChoiceContinuationInput`

NewChoiceContinuationInput instantiates a new ChoiceContinuationInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChoiceContinuationInputWithDefaults

`func NewChoiceContinuationInputWithDefaults() *ChoiceContinuationInput`

NewChoiceContinuationInputWithDefaults instantiates a new ChoiceContinuationInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinuationHash

`func (o *ChoiceContinuationInput) GetContinuationHash() string`

GetContinuationHash returns the ContinuationHash field if non-nil, zero value otherwise.

### GetContinuationHashOk

`func (o *ChoiceContinuationInput) GetContinuationHashOk() (*string, bool)`

GetContinuationHashOk returns a tuple with the ContinuationHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuationHash

`func (o *ChoiceContinuationInput) SetContinuationHash(v string)`

SetContinuationHash sets ContinuationHash field to given value.


### GetForChoiceId

`func (o *ChoiceContinuationInput) GetForChoiceId() ChoiceId`

GetForChoiceId returns the ForChoiceId field if non-nil, zero value otherwise.

### GetForChoiceIdOk

`func (o *ChoiceContinuationInput) GetForChoiceIdOk() (*ChoiceId, bool)`

GetForChoiceIdOk returns a tuple with the ForChoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForChoiceId

`func (o *ChoiceContinuationInput) SetForChoiceId(v ChoiceId)`

SetForChoiceId sets ForChoiceId field to given value.


### GetInputThatChoosesNum

`func (o *ChoiceContinuationInput) GetInputThatChoosesNum() int32`

GetInputThatChoosesNum returns the InputThatChoosesNum field if non-nil, zero value otherwise.

### GetInputThatChoosesNumOk

`func (o *ChoiceContinuationInput) GetInputThatChoosesNumOk() (*int32, bool)`

GetInputThatChoosesNumOk returns a tuple with the InputThatChoosesNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputThatChoosesNum

`func (o *ChoiceContinuationInput) SetInputThatChoosesNum(v int32)`

SetInputThatChoosesNum sets InputThatChoosesNum field to given value.


### GetMerkleizedContinuation

`func (o *ChoiceContinuationInput) GetMerkleizedContinuation() Contract`

GetMerkleizedContinuation returns the MerkleizedContinuation field if non-nil, zero value otherwise.

### GetMerkleizedContinuationOk

`func (o *ChoiceContinuationInput) GetMerkleizedContinuationOk() (*Contract, bool)`

GetMerkleizedContinuationOk returns a tuple with the MerkleizedContinuation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerkleizedContinuation

`func (o *ChoiceContinuationInput) SetMerkleizedContinuation(v Contract)`

SetMerkleizedContinuation sets MerkleizedContinuation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



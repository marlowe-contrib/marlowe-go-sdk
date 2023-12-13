# Input

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinuationHash** | **string** |  | 
**MerkleizedContinuation** | [**Contract**](Contract.md) |  | 
**ForChoiceId** | [**ChoiceId**](ChoiceId.md) |  | 
**InputThatChoosesNum** | **int32** |  | 
**InputFromParty** | [**Party**](Party.md) |  | 
**IntoAccount** | [**Party**](Party.md) |  | 
**OfToken** | [**Token**](Token.md) |  | 
**ThatDeposits** | **int32** |  | 

## Methods

### NewInput

`func NewInput(continuationHash string, merkleizedContinuation Contract, forChoiceId ChoiceId, inputThatChoosesNum int32, inputFromParty Party, intoAccount Party, ofToken Token, thatDeposits int32, ) *Input`

NewInput instantiates a new Input object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputWithDefaults

`func NewInputWithDefaults() *Input`

NewInputWithDefaults instantiates a new Input object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinuationHash

`func (o *Input) GetContinuationHash() string`

GetContinuationHash returns the ContinuationHash field if non-nil, zero value otherwise.

### GetContinuationHashOk

`func (o *Input) GetContinuationHashOk() (*string, bool)`

GetContinuationHashOk returns a tuple with the ContinuationHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuationHash

`func (o *Input) SetContinuationHash(v string)`

SetContinuationHash sets ContinuationHash field to given value.


### GetMerkleizedContinuation

`func (o *Input) GetMerkleizedContinuation() Contract`

GetMerkleizedContinuation returns the MerkleizedContinuation field if non-nil, zero value otherwise.

### GetMerkleizedContinuationOk

`func (o *Input) GetMerkleizedContinuationOk() (*Contract, bool)`

GetMerkleizedContinuationOk returns a tuple with the MerkleizedContinuation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerkleizedContinuation

`func (o *Input) SetMerkleizedContinuation(v Contract)`

SetMerkleizedContinuation sets MerkleizedContinuation field to given value.


### GetForChoiceId

`func (o *Input) GetForChoiceId() ChoiceId`

GetForChoiceId returns the ForChoiceId field if non-nil, zero value otherwise.

### GetForChoiceIdOk

`func (o *Input) GetForChoiceIdOk() (*ChoiceId, bool)`

GetForChoiceIdOk returns a tuple with the ForChoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForChoiceId

`func (o *Input) SetForChoiceId(v ChoiceId)`

SetForChoiceId sets ForChoiceId field to given value.


### GetInputThatChoosesNum

`func (o *Input) GetInputThatChoosesNum() int32`

GetInputThatChoosesNum returns the InputThatChoosesNum field if non-nil, zero value otherwise.

### GetInputThatChoosesNumOk

`func (o *Input) GetInputThatChoosesNumOk() (*int32, bool)`

GetInputThatChoosesNumOk returns a tuple with the InputThatChoosesNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputThatChoosesNum

`func (o *Input) SetInputThatChoosesNum(v int32)`

SetInputThatChoosesNum sets InputThatChoosesNum field to given value.


### GetInputFromParty

`func (o *Input) GetInputFromParty() Party`

GetInputFromParty returns the InputFromParty field if non-nil, zero value otherwise.

### GetInputFromPartyOk

`func (o *Input) GetInputFromPartyOk() (*Party, bool)`

GetInputFromPartyOk returns a tuple with the InputFromParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputFromParty

`func (o *Input) SetInputFromParty(v Party)`

SetInputFromParty sets InputFromParty field to given value.


### GetIntoAccount

`func (o *Input) GetIntoAccount() Party`

GetIntoAccount returns the IntoAccount field if non-nil, zero value otherwise.

### GetIntoAccountOk

`func (o *Input) GetIntoAccountOk() (*Party, bool)`

GetIntoAccountOk returns a tuple with the IntoAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntoAccount

`func (o *Input) SetIntoAccount(v Party)`

SetIntoAccount sets IntoAccount field to given value.


### GetOfToken

`func (o *Input) GetOfToken() Token`

GetOfToken returns the OfToken field if non-nil, zero value otherwise.

### GetOfTokenOk

`func (o *Input) GetOfTokenOk() (*Token, bool)`

GetOfTokenOk returns a tuple with the OfToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfToken

`func (o *Input) SetOfToken(v Token)`

SetOfToken sets OfToken field to given value.


### GetThatDeposits

`func (o *Input) GetThatDeposits() int32`

GetThatDeposits returns the ThatDeposits field if non-nil, zero value otherwise.

### GetThatDepositsOk

`func (o *Input) GetThatDepositsOk() (*int32, bool)`

GetThatDepositsOk returns a tuple with the ThatDeposits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThatDeposits

`func (o *Input) SetThatDeposits(v int32)`

SetThatDeposits sets ThatDeposits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



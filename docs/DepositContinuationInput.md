# DepositContinuationInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinuationHash** | **string** |  | 
**InputFromParty** | [**Party**](Party.md) |  | 
**IntoAccount** | [**Party**](Party.md) |  | 
**MerkleizedContinuation** | [**Contract**](Contract.md) |  | 
**OfToken** | [**Token**](Token.md) |  | 
**ThatDeposits** | **int32** |  | 

## Methods

### NewDepositContinuationInput

`func NewDepositContinuationInput(continuationHash string, inputFromParty Party, intoAccount Party, merkleizedContinuation Contract, ofToken Token, thatDeposits int32, ) *DepositContinuationInput`

NewDepositContinuationInput instantiates a new DepositContinuationInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepositContinuationInputWithDefaults

`func NewDepositContinuationInputWithDefaults() *DepositContinuationInput`

NewDepositContinuationInputWithDefaults instantiates a new DepositContinuationInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinuationHash

`func (o *DepositContinuationInput) GetContinuationHash() string`

GetContinuationHash returns the ContinuationHash field if non-nil, zero value otherwise.

### GetContinuationHashOk

`func (o *DepositContinuationInput) GetContinuationHashOk() (*string, bool)`

GetContinuationHashOk returns a tuple with the ContinuationHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuationHash

`func (o *DepositContinuationInput) SetContinuationHash(v string)`

SetContinuationHash sets ContinuationHash field to given value.


### GetInputFromParty

`func (o *DepositContinuationInput) GetInputFromParty() Party`

GetInputFromParty returns the InputFromParty field if non-nil, zero value otherwise.

### GetInputFromPartyOk

`func (o *DepositContinuationInput) GetInputFromPartyOk() (*Party, bool)`

GetInputFromPartyOk returns a tuple with the InputFromParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputFromParty

`func (o *DepositContinuationInput) SetInputFromParty(v Party)`

SetInputFromParty sets InputFromParty field to given value.


### GetIntoAccount

`func (o *DepositContinuationInput) GetIntoAccount() Party`

GetIntoAccount returns the IntoAccount field if non-nil, zero value otherwise.

### GetIntoAccountOk

`func (o *DepositContinuationInput) GetIntoAccountOk() (*Party, bool)`

GetIntoAccountOk returns a tuple with the IntoAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntoAccount

`func (o *DepositContinuationInput) SetIntoAccount(v Party)`

SetIntoAccount sets IntoAccount field to given value.


### GetMerkleizedContinuation

`func (o *DepositContinuationInput) GetMerkleizedContinuation() Contract`

GetMerkleizedContinuation returns the MerkleizedContinuation field if non-nil, zero value otherwise.

### GetMerkleizedContinuationOk

`func (o *DepositContinuationInput) GetMerkleizedContinuationOk() (*Contract, bool)`

GetMerkleizedContinuationOk returns a tuple with the MerkleizedContinuation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerkleizedContinuation

`func (o *DepositContinuationInput) SetMerkleizedContinuation(v Contract)`

SetMerkleizedContinuation sets MerkleizedContinuation field to given value.


### GetOfToken

`func (o *DepositContinuationInput) GetOfToken() Token`

GetOfToken returns the OfToken field if non-nil, zero value otherwise.

### GetOfTokenOk

`func (o *DepositContinuationInput) GetOfTokenOk() (*Token, bool)`

GetOfTokenOk returns a tuple with the OfToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfToken

`func (o *DepositContinuationInput) SetOfToken(v Token)`

SetOfToken sets OfToken field to given value.


### GetThatDeposits

`func (o *DepositContinuationInput) GetThatDeposits() int32`

GetThatDeposits returns the ThatDeposits field if non-nil, zero value otherwise.

### GetThatDepositsOk

`func (o *DepositContinuationInput) GetThatDepositsOk() (*int32, bool)`

GetThatDepositsOk returns a tuple with the ThatDeposits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThatDeposits

`func (o *DepositContinuationInput) SetThatDeposits(v int32)`

SetThatDeposits sets ThatDeposits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



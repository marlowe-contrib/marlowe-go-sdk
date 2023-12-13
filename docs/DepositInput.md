# DepositInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputFromParty** | [**Party**](Party.md) |  | 
**IntoAccount** | [**Party**](Party.md) |  | 
**OfToken** | [**Token**](Token.md) |  | 
**ThatDeposits** | **int32** |  | 

## Methods

### NewDepositInput

`func NewDepositInput(inputFromParty Party, intoAccount Party, ofToken Token, thatDeposits int32, ) *DepositInput`

NewDepositInput instantiates a new DepositInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepositInputWithDefaults

`func NewDepositInputWithDefaults() *DepositInput`

NewDepositInputWithDefaults instantiates a new DepositInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputFromParty

`func (o *DepositInput) GetInputFromParty() Party`

GetInputFromParty returns the InputFromParty field if non-nil, zero value otherwise.

### GetInputFromPartyOk

`func (o *DepositInput) GetInputFromPartyOk() (*Party, bool)`

GetInputFromPartyOk returns a tuple with the InputFromParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputFromParty

`func (o *DepositInput) SetInputFromParty(v Party)`

SetInputFromParty sets InputFromParty field to given value.


### GetIntoAccount

`func (o *DepositInput) GetIntoAccount() Party`

GetIntoAccount returns the IntoAccount field if non-nil, zero value otherwise.

### GetIntoAccountOk

`func (o *DepositInput) GetIntoAccountOk() (*Party, bool)`

GetIntoAccountOk returns a tuple with the IntoAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntoAccount

`func (o *DepositInput) SetIntoAccount(v Party)`

SetIntoAccount sets IntoAccount field to given value.


### GetOfToken

`func (o *DepositInput) GetOfToken() Token`

GetOfToken returns the OfToken field if non-nil, zero value otherwise.

### GetOfTokenOk

`func (o *DepositInput) GetOfTokenOk() (*Token, bool)`

GetOfTokenOk returns a tuple with the OfToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfToken

`func (o *DepositInput) SetOfToken(v Token)`

SetOfToken sets OfToken field to given value.


### GetThatDeposits

`func (o *DepositInput) GetThatDeposits() int32`

GetThatDeposits returns the ThatDeposits field if non-nil, zero value otherwise.

### GetThatDepositsOk

`func (o *DepositInput) GetThatDepositsOk() (*int32, bool)`

GetThatDepositsOk returns a tuple with the ThatDeposits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThatDeposits

`func (o *DepositInput) SetThatDeposits(v int32)`

SetThatDeposits sets ThatDeposits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



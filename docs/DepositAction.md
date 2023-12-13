# DepositAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deposits** | [**Value**](Value.md) |  | 
**IntoAccount** | [**Party**](Party.md) |  | 
**OfToken** | [**Token**](Token.md) |  | 
**Party** | [**Party**](Party.md) |  | 

## Methods

### NewDepositAction

`func NewDepositAction(deposits Value, intoAccount Party, ofToken Token, party Party, ) *DepositAction`

NewDepositAction instantiates a new DepositAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepositActionWithDefaults

`func NewDepositActionWithDefaults() *DepositAction`

NewDepositActionWithDefaults instantiates a new DepositAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeposits

`func (o *DepositAction) GetDeposits() Value`

GetDeposits returns the Deposits field if non-nil, zero value otherwise.

### GetDepositsOk

`func (o *DepositAction) GetDepositsOk() (*Value, bool)`

GetDepositsOk returns a tuple with the Deposits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeposits

`func (o *DepositAction) SetDeposits(v Value)`

SetDeposits sets Deposits field to given value.


### GetIntoAccount

`func (o *DepositAction) GetIntoAccount() Party`

GetIntoAccount returns the IntoAccount field if non-nil, zero value otherwise.

### GetIntoAccountOk

`func (o *DepositAction) GetIntoAccountOk() (*Party, bool)`

GetIntoAccountOk returns a tuple with the IntoAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntoAccount

`func (o *DepositAction) SetIntoAccount(v Party)`

SetIntoAccount sets IntoAccount field to given value.


### GetOfToken

`func (o *DepositAction) GetOfToken() Token`

GetOfToken returns the OfToken field if non-nil, zero value otherwise.

### GetOfTokenOk

`func (o *DepositAction) GetOfTokenOk() (*Token, bool)`

GetOfTokenOk returns a tuple with the OfToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfToken

`func (o *DepositAction) SetOfToken(v Token)`

SetOfToken sets OfToken field to given value.


### GetParty

`func (o *DepositAction) GetParty() Party`

GetParty returns the Party field if non-nil, zero value otherwise.

### GetPartyOk

`func (o *DepositAction) GetPartyOk() (*Party, bool)`

GetPartyOk returns a tuple with the Party field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParty

`func (o *DepositAction) SetParty(v Party)`

SetParty sets Party field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



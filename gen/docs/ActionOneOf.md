# ActionOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deposits** | [**Value**](Value.md) |  | 
**IntoAccount** | [**Party**](Party.md) |  | 
**OfToken** | [**Token**](Token.md) |  | 
**Party** | [**Party**](Party.md) |  | 

## Methods

### NewActionOneOf

`func NewActionOneOf(deposits Value, intoAccount Party, ofToken Token, party Party, ) *ActionOneOf`

NewActionOneOf instantiates a new ActionOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionOneOfWithDefaults

`func NewActionOneOfWithDefaults() *ActionOneOf`

NewActionOneOfWithDefaults instantiates a new ActionOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeposits

`func (o *ActionOneOf) GetDeposits() Value`

GetDeposits returns the Deposits field if non-nil, zero value otherwise.

### GetDepositsOk

`func (o *ActionOneOf) GetDepositsOk() (*Value, bool)`

GetDepositsOk returns a tuple with the Deposits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeposits

`func (o *ActionOneOf) SetDeposits(v Value)`

SetDeposits sets Deposits field to given value.


### GetIntoAccount

`func (o *ActionOneOf) GetIntoAccount() Party`

GetIntoAccount returns the IntoAccount field if non-nil, zero value otherwise.

### GetIntoAccountOk

`func (o *ActionOneOf) GetIntoAccountOk() (*Party, bool)`

GetIntoAccountOk returns a tuple with the IntoAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntoAccount

`func (o *ActionOneOf) SetIntoAccount(v Party)`

SetIntoAccount sets IntoAccount field to given value.


### GetOfToken

`func (o *ActionOneOf) GetOfToken() Token`

GetOfToken returns the OfToken field if non-nil, zero value otherwise.

### GetOfTokenOk

`func (o *ActionOneOf) GetOfTokenOk() (*Token, bool)`

GetOfTokenOk returns a tuple with the OfToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfToken

`func (o *ActionOneOf) SetOfToken(v Token)`

SetOfToken sets OfToken field to given value.


### GetParty

`func (o *ActionOneOf) GetParty() Party`

GetParty returns the Party field if non-nil, zero value otherwise.

### GetPartyOk

`func (o *ActionOneOf) GetPartyOk() (*Party, bool)`

GetPartyOk returns a tuple with the Party field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParty

`func (o *ActionOneOf) SetParty(v Party)`

SetParty sets Party field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



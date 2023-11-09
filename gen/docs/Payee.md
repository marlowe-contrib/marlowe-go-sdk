# Payee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | [**Party**](Party.md) |  | 
**Party** | [**Party**](Party.md) |  | 

## Methods

### NewPayee

`func NewPayee(account Party, party Party, ) *Payee`

NewPayee instantiates a new Payee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayeeWithDefaults

`func NewPayeeWithDefaults() *Payee`

NewPayeeWithDefaults instantiates a new Payee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *Payee) GetAccount() Party`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Payee) GetAccountOk() (*Party, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Payee) SetAccount(v Party)`

SetAccount sets Account field to given value.


### GetParty

`func (o *Payee) GetParty() Party`

GetParty returns the Party field if non-nil, zero value otherwise.

### GetPartyOk

`func (o *Payee) GetPartyOk() (*Party, bool)`

GetPartyOk returns a tuple with the Party field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParty

`func (o *Payee) SetParty(v Party)`

SetParty sets Party field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



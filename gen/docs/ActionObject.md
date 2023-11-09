# ActionObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deposits** | [**Value**](Value.md) |  | 
**IntoAccount** | [**Party**](Party.md) |  | 
**OfToken** | [**Token**](Token.md) |  | 
**Party** | [**Party**](Party.md) |  | 
**ChooseBetween** | [**[]Bound**](Bound.md) |  | 
**ForChoice** | [**ChoiceId**](ChoiceId.md) |  | 
**NotifyIf** | [**Observation**](Observation.md) |  | 
**Ref** | **string** | An arbitrary text identifier for an object in a Marlowe object bundle. | 

## Methods

### NewActionObject

`func NewActionObject(deposits Value, intoAccount Party, ofToken Token, party Party, chooseBetween []Bound, forChoice ChoiceId, notifyIf Observation, ref string, ) *ActionObject`

NewActionObject instantiates a new ActionObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionObjectWithDefaults

`func NewActionObjectWithDefaults() *ActionObject`

NewActionObjectWithDefaults instantiates a new ActionObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeposits

`func (o *ActionObject) GetDeposits() Value`

GetDeposits returns the Deposits field if non-nil, zero value otherwise.

### GetDepositsOk

`func (o *ActionObject) GetDepositsOk() (*Value, bool)`

GetDepositsOk returns a tuple with the Deposits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeposits

`func (o *ActionObject) SetDeposits(v Value)`

SetDeposits sets Deposits field to given value.


### GetIntoAccount

`func (o *ActionObject) GetIntoAccount() Party`

GetIntoAccount returns the IntoAccount field if non-nil, zero value otherwise.

### GetIntoAccountOk

`func (o *ActionObject) GetIntoAccountOk() (*Party, bool)`

GetIntoAccountOk returns a tuple with the IntoAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntoAccount

`func (o *ActionObject) SetIntoAccount(v Party)`

SetIntoAccount sets IntoAccount field to given value.


### GetOfToken

`func (o *ActionObject) GetOfToken() Token`

GetOfToken returns the OfToken field if non-nil, zero value otherwise.

### GetOfTokenOk

`func (o *ActionObject) GetOfTokenOk() (*Token, bool)`

GetOfTokenOk returns a tuple with the OfToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfToken

`func (o *ActionObject) SetOfToken(v Token)`

SetOfToken sets OfToken field to given value.


### GetParty

`func (o *ActionObject) GetParty() Party`

GetParty returns the Party field if non-nil, zero value otherwise.

### GetPartyOk

`func (o *ActionObject) GetPartyOk() (*Party, bool)`

GetPartyOk returns a tuple with the Party field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParty

`func (o *ActionObject) SetParty(v Party)`

SetParty sets Party field to given value.


### GetChooseBetween

`func (o *ActionObject) GetChooseBetween() []Bound`

GetChooseBetween returns the ChooseBetween field if non-nil, zero value otherwise.

### GetChooseBetweenOk

`func (o *ActionObject) GetChooseBetweenOk() (*[]Bound, bool)`

GetChooseBetweenOk returns a tuple with the ChooseBetween field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChooseBetween

`func (o *ActionObject) SetChooseBetween(v []Bound)`

SetChooseBetween sets ChooseBetween field to given value.


### GetForChoice

`func (o *ActionObject) GetForChoice() ChoiceId`

GetForChoice returns the ForChoice field if non-nil, zero value otherwise.

### GetForChoiceOk

`func (o *ActionObject) GetForChoiceOk() (*ChoiceId, bool)`

GetForChoiceOk returns a tuple with the ForChoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForChoice

`func (o *ActionObject) SetForChoice(v ChoiceId)`

SetForChoice sets ForChoice field to given value.


### GetNotifyIf

`func (o *ActionObject) GetNotifyIf() Observation`

GetNotifyIf returns the NotifyIf field if non-nil, zero value otherwise.

### GetNotifyIfOk

`func (o *ActionObject) GetNotifyIfOk() (*Observation, bool)`

GetNotifyIfOk returns a tuple with the NotifyIf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyIf

`func (o *ActionObject) SetNotifyIf(v Observation)`

SetNotifyIf sets NotifyIf field to given value.


### GetRef

`func (o *ActionObject) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *ActionObject) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *ActionObject) SetRef(v string)`

SetRef sets Ref field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



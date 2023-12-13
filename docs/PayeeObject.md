# PayeeObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | [**PartyObject**](PartyObject.md) |  | 
**Party** | [**PartyObject**](PartyObject.md) |  | 

## Methods

### NewPayeeObject

`func NewPayeeObject(account PartyObject, party PartyObject, ) *PayeeObject`

NewPayeeObject instantiates a new PayeeObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayeeObjectWithDefaults

`func NewPayeeObjectWithDefaults() *PayeeObject`

NewPayeeObjectWithDefaults instantiates a new PayeeObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *PayeeObject) GetAccount() PartyObject`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *PayeeObject) GetAccountOk() (*PartyObject, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *PayeeObject) SetAccount(v PartyObject)`

SetAccount sets Account field to given value.


### GetParty

`func (o *PayeeObject) GetParty() PartyObject`

GetParty returns the Party field if non-nil, zero value otherwise.

### GetPartyOk

`func (o *PayeeObject) GetPartyOk() (*PartyObject, bool)`

GetPartyOk returns a tuple with the Party field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParty

`func (o *PayeeObject) SetParty(v PartyObject)`

SetParty sets Party field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



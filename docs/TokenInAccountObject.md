# TokenInAccountObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountOfToken** | [**TokenObject**](TokenObject.md) |  | 
**InAccount** | [**PartyObject**](PartyObject.md) |  | 

## Methods

### NewTokenInAccountObject

`func NewTokenInAccountObject(amountOfToken TokenObject, inAccount PartyObject, ) *TokenInAccountObject`

NewTokenInAccountObject instantiates a new TokenInAccountObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenInAccountObjectWithDefaults

`func NewTokenInAccountObjectWithDefaults() *TokenInAccountObject`

NewTokenInAccountObjectWithDefaults instantiates a new TokenInAccountObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountOfToken

`func (o *TokenInAccountObject) GetAmountOfToken() TokenObject`

GetAmountOfToken returns the AmountOfToken field if non-nil, zero value otherwise.

### GetAmountOfTokenOk

`func (o *TokenInAccountObject) GetAmountOfTokenOk() (*TokenObject, bool)`

GetAmountOfTokenOk returns a tuple with the AmountOfToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountOfToken

`func (o *TokenInAccountObject) SetAmountOfToken(v TokenObject)`

SetAmountOfToken sets AmountOfToken field to given value.


### GetInAccount

`func (o *TokenInAccountObject) GetInAccount() PartyObject`

GetInAccount returns the InAccount field if non-nil, zero value otherwise.

### GetInAccountOk

`func (o *TokenInAccountObject) GetInAccountOk() (*PartyObject, bool)`

GetInAccountOk returns a tuple with the InAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInAccount

`func (o *TokenInAccountObject) SetInAccount(v PartyObject)`

SetInAccount sets InAccount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



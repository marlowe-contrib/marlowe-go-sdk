# ValueOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountOfToken** | [**Token**](Token.md) |  | 
**InAccount** | [**Party**](Party.md) |  | 

## Methods

### NewValueOneOf

`func NewValueOneOf(amountOfToken Token, inAccount Party, ) *ValueOneOf`

NewValueOneOf instantiates a new ValueOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValueOneOfWithDefaults

`func NewValueOneOfWithDefaults() *ValueOneOf`

NewValueOneOfWithDefaults instantiates a new ValueOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountOfToken

`func (o *ValueOneOf) GetAmountOfToken() Token`

GetAmountOfToken returns the AmountOfToken field if non-nil, zero value otherwise.

### GetAmountOfTokenOk

`func (o *ValueOneOf) GetAmountOfTokenOk() (*Token, bool)`

GetAmountOfTokenOk returns a tuple with the AmountOfToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountOfToken

`func (o *ValueOneOf) SetAmountOfToken(v Token)`

SetAmountOfToken sets AmountOfToken field to given value.


### GetInAccount

`func (o *ValueOneOf) GetInAccount() Party`

GetInAccount returns the InAccount field if non-nil, zero value otherwise.

### GetInAccountOk

`func (o *ValueOneOf) GetInAccountOk() (*Party, bool)`

GetInAccountOk returns a tuple with the InAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInAccount

`func (o *ValueOneOf) SetInAccount(v Party)`

SetInAccount sets InAccount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# Pay

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromAccount** | [**Party**](Party.md) |  | 
**Pay** | [**Value**](Value.md) |  | 
**Then** | [**Contract**](Contract.md) |  | 
**To** | [**Payee**](Payee.md) |  | 
**Token** | [**Token**](Token.md) |  | 

## Methods

### NewPay

`func NewPay(fromAccount Party, pay Value, then Contract, to Payee, token Token, ) *Pay`

NewPay instantiates a new Pay object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayWithDefaults

`func NewPayWithDefaults() *Pay`

NewPayWithDefaults instantiates a new Pay object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromAccount

`func (o *Pay) GetFromAccount() Party`

GetFromAccount returns the FromAccount field if non-nil, zero value otherwise.

### GetFromAccountOk

`func (o *Pay) GetFromAccountOk() (*Party, bool)`

GetFromAccountOk returns a tuple with the FromAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAccount

`func (o *Pay) SetFromAccount(v Party)`

SetFromAccount sets FromAccount field to given value.


### GetPay

`func (o *Pay) GetPay() Value`

GetPay returns the Pay field if non-nil, zero value otherwise.

### GetPayOk

`func (o *Pay) GetPayOk() (*Value, bool)`

GetPayOk returns a tuple with the Pay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPay

`func (o *Pay) SetPay(v Value)`

SetPay sets Pay field to given value.


### GetThen

`func (o *Pay) GetThen() Contract`

GetThen returns the Then field if non-nil, zero value otherwise.

### GetThenOk

`func (o *Pay) GetThenOk() (*Contract, bool)`

GetThenOk returns a tuple with the Then field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThen

`func (o *Pay) SetThen(v Contract)`

SetThen sets Then field to given value.


### GetTo

`func (o *Pay) GetTo() Payee`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *Pay) GetToOk() (*Payee, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *Pay) SetTo(v Payee)`

SetTo sets To field to given value.


### GetToken

`func (o *Pay) GetToken() Token`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *Pay) GetTokenOk() (*Token, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *Pay) SetToken(v Token)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


